package main

import (
	"fmt"
	"logAgent/kafka"
	"logAgent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

const (
	EtcdKey = "/logagent/config/127.0.0.1"
)

func main() {
	// 加载配置文件
	filename := "/Users/xliu/Documents/51CTO_Golang/learnGo/src/logAgent/conf/loadConf.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err: %v\n", err)
		panic("load conf failed")
	}
	fmt.Println(appConfig.logLevel)

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err: %v\n", err)
		panic("load logger failed")
	}
	logs.Debug("load conf succ, conf: %+v", appConfig)

	// 初始化配置文件
	appConfig.collectConf, err = initEtcd(appConfig.etcdAddr, appConfig.etcdKey)
	if err != nil {
		logs.Error("init etcd failed, err: %v", err)
	}
	logs.Debug("initialize initEtcd succ")

	// 初始化日志信息
	err = tailf.InitTail(appConfig.collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err: %v", err)
	}
	logs.Debug("initialize all succ")

	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err: %v", err)
		return
	}

	go func() {
		var count int
		for {
			count++
			logs.Debug("test for logger %d", count)
			time.Sleep(time.Second * 1)
		}
	}()
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err: %v", err)
		return
	}
	logs.Info("program exited")
}
