package service

import (
	"sync"
	"time"

	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
)

var (
	secLayerContext = &SecLayerContext{}
)

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
	RedisQueueName   string
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
	HandleUserGoroutineNum       int
	SecProductInfoMap            map[int]*SecProductInfoConf
	Read2HandleChanSize          int
	Handle2WriterChanSize        int
	MaxRequestWaitTimeout        int
	SendToHandleChanTimeout      int
	SendToWriterChanTimeout      int
	TokenPasswd                  string
}

type SecLayerContext struct {
	proxy2LayerRedisPool *redis.Pool
	layer2ProxyRedisPool *redis.Pool
	etcdClient           *etcd_client.Client
	RWSecProductLock     sync.RWMutex
	secLayerConf         *SecLayerConf
	waitGroup            sync.WaitGroup
	Read2HandleChan      chan *SecRequest
	Handle2WriterChan    chan *SecResponse
	HistoryMap           map[int]*UserBuyHistory
	HistoryMapLock       sync.Mutex
	productCountMgr      *ProductCountMgr
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Left      int

	// 限流控制
	secLimit *SecLimit
	// 每秒钟最多卖多少商品
	SoldLimit int

	// 对用户购买商品数量限制
	OnePersonBuyLimit int

	// 随机抽奖
	BuyRate float64
}

type SecRequest struct {
	ProductId       int
	Source          string
	AuthCode        string
	SecTime         string
	Nance           string
	UserId          int
	UserAuthSign    string
	AccessTime      time.Time
	ClientAddr      string
	ClientReference string
}

type SecResponse struct {
	ProductId int
	UserId    int
	Token     string
	TokenTime int64
	Code      int
}
