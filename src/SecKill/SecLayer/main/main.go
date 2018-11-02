package main

import (
	"SecKill/SecLayer/service"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	// 加载配置文件
	err := initConfig("ini", "./conf/seclayer.conf")
	if err != nil {
		logs.Error("init config failed, err: %v", err)
		panic(fmt.Sprintf("init config failed"))
	}

	// 初始化日志
	err = initLogger()
	if err != nil {
		logs.Error("init logger failed, err: %v", err)
		panic(fmt.Sprintf("init logger failed"))
	}
	logs.Debug("init logger succ")

	// 初始化秒杀逻辑
	err = service.InitSecLayer(appConfig)
	if err != nil {
		logs.Error("init seclayer failed, err: %v", err)
		panic(fmt.Sprintf("init seclayer failed"))
	}
	logs.Debug("init sec layer succ")

	// 运行业务逻辑
	err = service.Run()
	if err != nil {
		logs.Error("init Run failed, err: %v", err)
		return
	}
	logs.Debug("service run existed")

}
