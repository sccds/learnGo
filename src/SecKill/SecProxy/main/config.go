package main

import (
	"SecKill/SecProxy/service"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	secKillConf = &service.SecKillConf{
		SecProductInfoConfMap: make(map[int]*service.SecProductInfoConf, 1024),
	}
)

func initConfig() (err error) {
	etcdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("read config succ, etcd addr: %v", etcdAddr)
	secKillConf.EtcdConf.EtcdAddr = etcdAddr

	redisBlackAddr := beego.AppConfig.String("redis_black_addr")
	logs.Debug("read config succ, redis black addr: %v", redisBlackAddr)
	secKillConf.RedisBlackConf.RedisAddr = redisBlackAddr

	if len(redisBlackAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redisBlack[%s] or etcd[%s] config is null", redisBlackAddr, etcdAddr)
		return
	}

	// redis black parameters config
	redisBlackMaxIdle, err := beego.AppConfig.Int("redis_black_max_idle")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_max_idle error: %v", err)
		return
	}
	redisBlackMaxActive, err := beego.AppConfig.Int("redis_black_max_active")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_max_active error: %v", err)
		return
	}
	redisBlackIdleTimeout, err := beego.AppConfig.Int("redis_black_idle_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_idle_timeout error: %v", err)
		return
	}
	secKillConf.RedisBlackConf.RedisMaxIdle = redisBlackMaxIdle
	secKillConf.RedisBlackConf.RedisMaxActive = redisBlackMaxActive
	secKillConf.RedisBlackConf.RedisIdleTimeout = redisBlackIdleTimeout

	// etcd parameters config
	etcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed, read etcd_timeout error: %v", err)
		return
	}
	secKillConf.EtcdConf.EtcdTimeout = etcdTimeout

	secKillConf.EtcdConf.EtcdSecKeyPrefix = beego.AppConfig.String("etcd_sec_key_prefix")
	if len(secKillConf.EtcdConf.EtcdSecKeyPrefix) == 0 {
		err = fmt.Errorf("init config failed, read etcd_sec_key error: %v", err)
		return
	}
	productKey := beego.AppConfig.String("etcd_product_key")
	if len(productKey) == 0 {
		err = fmt.Errorf("init config failed, read etcd_product_key error: %v", err)
		return
	}
	if strings.HasSuffix(secKillConf.EtcdConf.EtcdSecKeyPrefix, "/") == false {
		secKillConf.EtcdConf.EtcdSecKeyPrefix = secKillConf.EtcdConf.EtcdSecKeyPrefix + "/"
	}
	secKillConf.EtcdConf.EtcdSecProductKey = fmt.Sprintf("%s%s", secKillConf.EtcdConf.EtcdSecKeyPrefix, productKey)

	// log config
	secKillConf.LogPath = beego.AppConfig.String("log_path")
	secKillConf.LogLevel = beego.AppConfig.String("log_level")

	// secret key config
	secKillConf.CookieSecretKey = beego.AppConfig.String("cookie_secretkey")

	referList := beego.AppConfig.String("refer_whitelist")
	if len(referList) > 0 {
		secKillConf.ReferWhiteList = strings.Split(referList, ",")
	}

	// access limit config
	secKillConf.AccessLimitConf.UserSecAccessLimit, err = beego.AppConfig.Int("user_sec_access_limit")
	if err != nil {
		err = fmt.Errorf("init config failed, read user_sec_access_limit error: %v", err)
		return
	}
	secKillConf.AccessLimitConf.UserMinAccessLimit, err = beego.AppConfig.Int("user_min_access_limit")
	if err != nil {
		err = fmt.Errorf("init config failed, read user_min_access_limit error: %v", err)
		return
	}
	secKillConf.AccessLimitConf.IpSecAccessLimit, err = beego.AppConfig.Int("ip_sec_access_limit")
	if err != nil {
		err = fmt.Errorf("init config failed, read ip_sec_access_limit error: %v", err)
		return
	}
	secKillConf.AccessLimitConf.IpMinAccessLimit, err = beego.AppConfig.Int("ip_min_access_limit")
	if err != nil {
		err = fmt.Errorf("init config failed, read ip_min_access_limit error: %v", err)
		return
	}

	// proxy => layer redis config
	redisProxy2LayerAddr := beego.AppConfig.String("redis_proxy2layer_addr")
	logs.Debug("read config succ, redis addr: %v", redisProxy2LayerAddr)
	secKillConf.RedisProxy2LayerConf.RedisAddr = redisProxy2LayerAddr
	if len(redisProxy2LayerAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] config is null", redisProxy2LayerAddr)
		return
	}

	redisProxy2LayerMaxIdle, err := beego.AppConfig.Int("redis_proxy2layer_max_idle")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_max_idle error: %v", err)
		return
	}
	redisProxy2LayerMaxActive, err := beego.AppConfig.Int("redis_proxy2layer_max_active")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_max_active error: %v", err)
		return
	}
	redisProxy2LayerIdleTimeout, err := beego.AppConfig.Int("redis_proxy2layer_idle_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_idle_timeout error: %v", err)
		return
	}
	secKillConf.RedisProxy2LayerConf.RedisMaxIdle = redisProxy2LayerMaxIdle
	secKillConf.RedisProxy2LayerConf.RedisMaxActive = redisProxy2LayerMaxActive
	secKillConf.RedisProxy2LayerConf.RedisIdleTimeout = redisProxy2LayerIdleTimeout

	// write, read chan number
	writeGoNums, err := beego.AppConfig.Int("write_proxy2layer_goroutine_num")
	if err != nil {
		err = fmt.Errorf("init config failed, read write_proxy2layer_goroutine_num error: %v", err)
		return
	}
	secKillConf.WriteProxy2LayerGoroutineNum = writeGoNums

	readGoNums, err := beego.AppConfig.Int("read_layer2proxy_goroutine_num")
	if err != nil {
		err = fmt.Errorf("init config failed, read read_layer2proxy_goroutine_num error: %v", err)
		return
	}
	secKillConf.ReadLayer2ProxyGoroutineNum = readGoNums

	// layer => proxy redis config
	redisLayer2ProxyAddr := beego.AppConfig.String("redis_layer2proxy_addr")
	logs.Debug("read config succ, redis addr: %v", redisLayer2ProxyAddr)
	secKillConf.RedisLayer2ProxyConf.RedisAddr = redisLayer2ProxyAddr
	if len(redisLayer2ProxyAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] config is null", redisLayer2ProxyAddr)
		return
	}

	redisLayer2ProxyMaxIdle, err := beego.AppConfig.Int("redis_layer2proxy_max_idle")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_max_idle error: %v", err)
		return
	}
	redisLayer2ProxyMaxActive, err := beego.AppConfig.Int("redis_layer2proxy_max_active")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_max_active error: %v", err)
		return
	}
	redisLayer2ProxyIdleTimeout, err := beego.AppConfig.Int("redis_layer2proxy_idle_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed, read redis_idle_timeout error: %v", err)
		return
	}
	secKillConf.RedisLayer2ProxyConf.RedisMaxIdle = redisLayer2ProxyMaxIdle
	secKillConf.RedisLayer2ProxyConf.RedisMaxActive = redisLayer2ProxyMaxActive
	secKillConf.RedisLayer2ProxyConf.RedisIdleTimeout = redisLayer2ProxyIdleTimeout

	return
}
