package service

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	ProductStatusNormal       = 0
	ProductStatusSaleOut      = 1
	ProductStatusForceSaleOut = 2
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

type AccessLimitConf struct {
	UserSecAccessLimit int
	UserMinAccessLimit int
	IpSecAccessLimit   int
	IpMinAccessLimit   int
}

type SecKillConf struct {
	RedisBlackConf               RedisConf
	RedisProxy2LayerConf         RedisConf
	RedisLayer2ProxyConf         RedisConf
	EtcdConf                     EtcdConf
	LogPath                      string
	LogLevel                     string
	SecProductInfoConfMap        map[int]*SecProductInfoConf
	CookieSecretKey              string
	UserSecAccessLimit           int
	IpSecAccessLimit             int
	ReferWhiteList               []string
	ipBlackMap                   map[string]bool
	idBlackMap                   map[int]bool
	blackRedisPool               *redis.Pool
	proxy2LayerRedisPool         *redis.Pool
	layer2ProxyRedisPool         *redis.Pool
	secLimitMgr                  *SecLimitMgr
	AccessLimitConf              AccessLimitConf
	WriteProxy2LayerGoroutineNum int
	ReadLayer2ProxyGoroutineNum  int
	RwSecProductLock             sync.RWMutex
	RwBlackLock                  sync.RWMutex
	secReqChan                   chan *SecRequest
	secReqChanSize               int
	UserConnMap                  map[string]chan *SecResult
	UserConnMapLock              sync.Mutex
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Left      int
}

type SecResult struct {
	ProductId int
	UserId    int
	Code      int
	Token     string
}

type SecRequest struct {
	ProductId       int
	Source          string
	Authcode        string
	SecTime         string
	Nance           string
	UserId          int
	AccessTime      time.Time
	UserAuthSign    string
	ClientAddr      string
	ClientReference string
	CloseNotify     <-chan bool     `json:"-"`
	ResultChan      chan *SecResult `json:"-"`
}
