package main

import (
	"SecKill/SecWeb/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/jmoiron/sqlx"
)

var (
	Db         *sqlx.DB
	EtcdClient *etcd_client.Client
)

func initDb() (err error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", AppConf.mysqlConf.UserName, AppConf.mysqlConf.Passwd, AppConf.mysqlConf.Host, AppConf.mysqlConf.Port, AppConf.mysqlConf.Database)
	database, err := sqlx.Open("mysql", dns)
	if err != nil {
		logs.Error("open mysql failed, dns: %s, err: %v", dns, err)
		return
	}
	Db = database
	logs.Debug("connect to mysql succ")
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
	config["filename"] = AppConf.LogPath
	config["level"] = convertLogLevel(AppConf.LogLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Marshal failed, err", err)
	}
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

func initEtcd() (err error) {
	cli, err := etcd_client.New(
		etcd_client.Config{
			Endpoints:   []string{AppConf.etcdConf.Addr},
			DialTimeout: time.Duration(AppConf.etcdConf.Timeout) * time.Second,
		})
	if err != nil {
		logs.Error("connect etcd failed, err: %v", err)
		return
	}

	EtcdClient = cli
	logs.Debug("init etcd succ")
	return
}

func initAll() (err error) {
	err = initConfig()
	if err != nil {
		logs.Warn("init config failed, err: %v", err)
		return
	}

	err = initLogger()
	if err != nil {
		logs.Warn("init logger failed, err: %v", err)
		return
	}

	err = initDb()
	if err != nil {
		logs.Warn("init db failed, err: %v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err: %v", err)
		return
	}

	err = model.Init(Db, EtcdClient, AppConf.etcdConf.EtcdKeyPrefix, AppConf.etcdConf.ProductKey)
	if err != nil {
		logs.Warn("init model failed, err: %v", err)
		return
	}
	return
}
