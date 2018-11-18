package main

import (
	"SecKill/SecLayer/service"
	"fmt"
	"strings"

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

	// 读取 Redis 相关配置
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
	appConfig.Proxy2LayerRedis.RedisQueueName = conf.String("redis::redis_proxy2layer_queue_name")
	if len(appConfig.Proxy2LayerRedis.RedisQueueName) == 0 {
		err = fmt.Errorf("read redis::redis_proxy2layer_queue_name failed")
		logs.Error("read redis::redis_proxy2layer_queue_name failed")
		return
	}
	appConfig.HandleUserGoroutineNum, err = conf.Int("service::handle_user_goroutine_num")
	if err != nil {
		logs.Error("read service::handle_user_goroutine_numfailed")
		return
	}
	appConfig.Read2HandleChanSize, err = conf.Int("service::read2handle_chan_size")
	if err != nil {
		logs.Error("read service::read2handle_chan_size failed")
		return
	}
	appConfig.Handle2WriterChanSize, err = conf.Int("service::handle2writer_chan_size")
	if err != nil {
		logs.Error("read service::handle2writer_chan_size failed")
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
	appConfig.Layer2ProxyRedis.RedisQueueName = conf.String("redis::redis_layer2proxy_queue_name")
	if len(appConfig.Layer2ProxyRedis.RedisQueueName) == 0 {
		err = fmt.Errorf("read redis::redis_layer2proxy_queue_name failed")
		logs.Error("read redis::redis_layer2proxy_queue_name failed")
		return
	}

	// 读取 etcd 相关配置
	appConfig.EtcdConfig.EtcdAddr = conf.String("etcd::etcd_addr")
	if len(appConfig.EtcdConfig.EtcdAddr) == 0 {
		logs.Error("read etcd::etcd_addr failed")
		err = fmt.Errorf("read etcd::etcd_addr failed")
		return
	}
	etcdTimeout, err := conf.Int("etcd::etcd_timeout")
	if err != nil {
		err = fmt.Errorf("read etcd::etcd_timeout failed")
		return
	}
	appConfig.EtcdConfig.EtcdTimeout = etcdTimeout

	appConfig.EtcdConfig.EtcdSecKeyPrefix = conf.String("etcd::etcd_sec_key_prefix")
	if len(appConfig.EtcdConfig.EtcdSecKeyPrefix) == 0 {
		err = fmt.Errorf("read etcd::etcd_sec_key_prefix failed")
		return
	}

	productKey := conf.String("etcd::etcd_product_key")
	if len(productKey) == 0 {
		err = fmt.Errorf("read etcd::etcd_product_key failed")
		return
	}
	if strings.HasSuffix(appConfig.EtcdConfig.EtcdSecKeyPrefix, "/") == false {
		appConfig.EtcdConfig.EtcdSecKeyPrefix = appConfig.EtcdConfig.EtcdSecKeyPrefix + "/"
	}
	appConfig.EtcdConfig.EtcdSecProductKey = fmt.Sprintf("%s%s", appConfig.EtcdConfig.EtcdSecKeyPrefix, productKey)
	logs.Debug("appConfig.EtcdConfig.EtcdSecProductKey: %s", appConfig.EtcdConfig.EtcdSecProductKey)

	appConfig.MaxRequestWaitTimeout, err = conf.Int("service::max_request_wait_timeout")
	if err != nil {
		logs.Error("read service::max_request_wait_timeout failed")
		return
	}

	appConfig.SendToHandleChanTimeout, err = conf.Int("service::send_to_handle_chan_timeout")
	if err != nil {
		logs.Error("read service::send_to_handle_chan_timeout failed")
		return
	}
	appConfig.SendToWriterChanTimeout, err = conf.Int("service::send_to_writer_chan_timeout")
	if err != nil {
		logs.Error("read service::send_to_writer_chan_timeout failed")
		return
	}

	// 读取 token 密钥
	appConfig.TokenPasswd = conf.String("service::seckill_token_password")
	if len(appConfig.TokenPasswd) == 0 {
		err = fmt.Errorf("read service::seckill_token_password failed")
		return
	}

	return
}
