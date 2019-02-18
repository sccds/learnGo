package main

import (
	"errors"
	"fmt"
	"logAgent/tailf"

	"github.com/astaxie/beego/config"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel    string
	logPath     string
	chanSize    int
	kafkaAddr   string
	etcdAddr    string
	etcdKey     string
	collectConf []tailf.CollectConf
}

func loadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	appConfig = &Config{}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server::port failed, err:", err)
		return
	}
	fmt.Println("Port:", port)

	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}
	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "src/logAgent/logs/logAgent.log"
	}
	appConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	appConfig.etcdAddr = conf.String("etcd::addr")
	if len(appConfig.etcdAddr) == 0 {
		err = fmt.Errorf("invalid etcd addr")
		return
	}
	appConfig.etcdKey = conf.String("etcd::configKey")
	if len(appConfig.etcdKey) == 0 {
		err = fmt.Errorf("invalid etcd::configKey addr")
		return
	}
	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("loadcollect conf failed, err: %v\n", err)
		return
	}
	return
}

// 读取监听日志的文件信息和对应topic
func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}
	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}
	appConfig.collectConf = append(appConfig.collectConf, cc)
	return
}
