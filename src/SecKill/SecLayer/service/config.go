package service

import "github.com/garyburd/redigo/redis"

var (
	secLayerContext = SecLayerContext{}
)

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

type EtcdConf struct {
	EtcdAddr          string
	EtcdTimeout       int
	EtcdSecKeyPrefix  string
	EtcdSecProductKey string
}

type SecLayerConf struct {
	Proxy2LayerRedis             RedisConf
	Layer2ProxyRedis             RedisConf
	EtcdConfig                   EtcdConf
	LogPath                      string
	LogLevel                     string
	WriteProxy2LayerGoroutineNum int
	ReadLayer2ProxyGoroutineNum  int
}

type SecLayerContext struct {
	proxy2LayerRedisPool *redis.Pool
	layer2ProxyRedisPool *redis.Pool
}
