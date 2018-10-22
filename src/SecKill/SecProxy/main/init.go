package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	etcd "github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
)

var (
	redisPool  *redis.Pool
	etcdClient *etcd.Client
)

func initEtcd() (err error) {
	cli, err := etcd.New(
		etcd.Config{
			Endpoints:            []string{secKillConf.etcdConf.etcdAddr},
			DialKeepAliveTimeout: time.Duration(secKillConf.etcdConf.etcdTimeout) * time.Second,
		})
	if err != nil {
		logs.Error("connect etcd failed, err: %v", err)
		return
	}
	etcdClient = cli
	return
}

func initRedis() (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     secKillConf.redisConf.redisMaxIdle,
		MaxActive:   secKillConf.redisConf.redisMaxActive,
		IdleTimeout: time.Duration(secKillConf.redisConf.redisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.redisConf.redisAddr)
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
	config["filename"] = secKillConf.logPath
	config["level"] = convertLogLevel(secKillConf.logLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Marshal failed, err", err)
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

func loadSecConf() (err error) {
	resp, err := etcdClient.Get(context.Background(), secKillConf.etcdConf.etcdSecProductKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err: %v", secKillConf.etcdConf.etcdSecProductKey, err)
		return
	}
	var secProductInfo []SecProductInfoConf
	for _, v := range resp.Kvs {
		//logs.Debug("key[%s] : value[%s]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("unmarshal sec product info failed, err: %v", err)
			return
		}
		logs.Debug("sec info config is [%v]", secProductInfo)
	}
	secKillConf.secProductInfoConf = secProductInfo
	return
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

	// init redis
	err = initRedis()
	if err != nil {
		logs.Error("init redis failed, err: %v", err)
		return
	}

	err = loadSecConf()

	logs.Info("init sec succ")
	return
}
