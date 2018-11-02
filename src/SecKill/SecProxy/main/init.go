package main

import (
	"SecKill/SecProxy/service"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	etcd "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/garyburd/redigo/redis"
)

var (
	redisPool  *redis.Pool
	etcdClient *etcd.Client
)

func initEtcd() (err error) {
	cli, err := etcd.New(
		etcd.Config{
			Endpoints:            []string{secKillConf.EtcdConf.EtcdAddr},
			DialKeepAliveTimeout: time.Duration(secKillConf.EtcdConf.EtcdTimeout) * time.Second,
		})
	if err != nil {
		logs.Error("connect etcd failed, err: %v", err)
		return
	}
	etcdClient = cli
	return
}

/*
func initRedis() (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     secKillConf.RedisBlackConf.RedisMaxIdle,
		MaxActive:   secKillConf.RedisBlackConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisBlackConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.RedisBlackConf.RedisAddr)
		},
	}
	conn := redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err: %v", err)
		return
	}
	return
}
*/

// log level mapping
func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = secKillConf.LogPath
	config["level"] = convertLogLevel(secKillConf.LogLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Marshal failed, err", err)
	}
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

func loadSecConf() (err error) {
	resp, err := etcdClient.Get(context.Background(), secKillConf.EtcdConf.EtcdSecProductKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err: %v", secKillConf.EtcdConf.EtcdSecProductKey, err)
		return
	}
	var secProductInfo []service.SecProductInfoConf
	for k, v := range resp.Kvs {
		logs.Debug("key[%s] : value[%s]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("unmarshal sec product info failed, err: %v", err)
			return
		}
		logs.Debug("sec info config is [%v]", secProductInfo)
	}
	updateSecProductInfo(secProductInfo)
	return
}

// listen etcd changes
func initSecProductWatcher() {
	go watchSecProductKey(secKillConf.EtcdConf.EtcdSecProductKey)
}

func watchSecProductKey(key string) {
	cli, err := etcd.New(
		etcd.Config{
			Endpoints:            []string{secKillConf.EtcdConf.EtcdAddr},
			DialKeepAliveTimeout: time.Duration(secKillConf.EtcdConf.EtcdTimeout) * time.Second,
		})
	if err != nil {
		logs.Error("connect etcd failed, err: ", err)
		return
	}
	logs.Debug("begin watch key: %s", key)

	for {
		rch := cli.Watch(context.Background(), key)
		var secProductInfo []service.SecProductInfoConf
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] config deleted", key)
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &secProductInfo)
					if err != nil {
						logs.Error("key[%s] unmarshal[%s] error, err: ", ev.Kv.Key, ev.Kv.Value, err)
						getConfSucc = false
					}
				}
				logs.Debug("get config from etcd, %s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", secProductInfo)
				updateSecProductInfo(secProductInfo)
			}
		}
	}
}

func updateSecProductInfo(secProductInfo []service.SecProductInfoConf) {
	var tmp map[int]*service.SecProductInfoConf = make(map[int]*service.SecProductInfoConf, 1024)
	for _, v := range secProductInfo {
		productInfo := v
		tmp[v.ProductId] = &productInfo
	}

	secKillConf.RwSecProductLock.Lock()
	secKillConf.SecProductInfoConfMap = tmp
	secKillConf.RwSecProductLock.Unlock()
}

func initSec() (err error) {

	// init log
	err = initLogger()
	if err != nil {
		logs.Error("init logger failed, err: %v", err)
		return
	}

	// init etcd
	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed, err: %v", err)
		return
	}

	/*
		// init redis
		err = initRedis()
		if err != nil {
			logs.Error("init redis failed, err: %v", err)
			return
		}
	*/

	err = loadSecConf()
	if err != nil {
		logs.Error("init conf failed, err: %v", err)
	}

	service.InitService(secKillConf)

	initSecProductWatcher()

	logs.Info("init sec succ")
	return
}
