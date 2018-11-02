package main

import (
	"SecKill/SecLayer/service"
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var (
	appConfig *service.SecLayerConf
)

func initConfig(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		logs.Debug("new config failed, err: %v", err)
		return
	}
	// 读取日志配置信息
	appConfig = &service.SecLayerConf{}
	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}
	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogPath = "./logs"
	}

	// 读取相关配置
	appConfig.Proxy2LayerRedis.RedisAddr = conf.String("redis::redis_proxy2layer_addr")
	if len(appConfig.Proxy2LayerRedis.RedisAddr) == 0 {
		err = fmt.Errorf("read redis::redis_proxy2layer_addr failed")
		logs.Error("read redis::redis_proxy2layer_addr failed")
		return
	}
	appConfig.Proxy2LayerRedis.RedisMaxIdle, err = conf.Int("redis::redis_proxy2layer_max_idle")
	if err != nil {
		logs.Error("read redis::redis_proxy2layer_max_idle failed")
		return
	}
	appConfig.Proxy2LayerRedis.RedisIdleTimeout, err = conf.Int("redis::redis_proxy2layer_idle_timeout")
	if err != nil {
		logs.Error("read redis::redis_proxy2layer_idle_timeout failed")
		return
	}
	appConfig.Proxy2LayerRedis.RedisMaxActive, err = conf.Int("redis::redis_proxy2layer_max_active")
	if err != nil {
		logs.Error("read redis::redis_proxy2layer_max_active failed")
		return
	}

	appConfig.WriteProxy2LayerGoroutineNum, err = conf.Int("redis::write_proxy2layer_goroutine_num")
	if err != nil {
		logs.Error("read redis::write_proxy2layer_goroutine_num failed")
		return
	}
	appConfig.ReadLayer2ProxyGoroutineNum, err = conf.Int("redis::read_layer2proxy_goroutine_num")
	if err != nil {
		logs.Error("read redis::read_layer2proxy_goroutine_num failed")
		return
	}

	appConfig.Layer2ProxyRedis.RedisAddr = conf.String("redis::redis_layer2proxy_addr")
	if len(appConfig.Layer2ProxyRedis.RedisAddr) == 0 {
		err = fmt.Errorf("read redis::redis_layer2proxy_addr failed")
		logs.Error("read redis::redis_layer2proxy_addr failed")
		return
	}
	appConfig.Layer2ProxyRedis.RedisMaxIdle, err = conf.Int("redis::redis_layer2proxy_max_idle")
	if err != nil {
		logs.Error("read redis::redis_layer2proxy_max_idle failed")
		return
	}
	appConfig.Layer2ProxyRedis.RedisIdleTimeout, err = conf.Int("redis::redis_layer2proxy_idle_timeout")
	if err != nil {
		logs.Error("read redis::redis_layer2proxy_idle_timeout failed")
		return
	}
	appConfig.Layer2ProxyRedis.RedisMaxActive, err = conf.Int("redis::redis_layer2proxy_max_active")
	if err != nil {
		logs.Error("read redis::redis_layer2proxy_max_active failed")
		return
	}

	return
}
