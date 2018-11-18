package service

import (
	"time"

	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
)

func initEtcd(conf *SecLayerConf) (err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{conf.EtcdConfig.EtcdAddr},
		DialTimeout: time.Duration(conf.EtcdConfig.EtcdTimeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err: %v", err)
		return
	}
	secLayerContext.etcdClient = cli
	logs.Debug("init etcd succ")
	return
}

func InitSecLayer(config *SecLayerConf) (err error) {
	// 初始化 redis
	err = initRedis(config)
	if err != nil {
		logs.Error("init redis failed, err: %v", err)
		return
	}

	// 初始化 etcd
	err = initEtcd(config)
	if err != nil {
		logs.Error("init etcd failed, err: %v", err)
		return
	}
	logs.Debug("init etcd succ")

	// 从 etcd 加载产品信息
	err = loadProductFromEtcd(config)
	if err != nil {
		logs.Error("load product from etcd failed, err: %v", err)
		return
	}
	logs.Debug("load product from etcd succ")

	secLayerContext.secLayerConf = config
	secLayerContext.Read2HandleChan = make(chan *SecRequest, secLayerContext.secLayerConf.Read2HandleChanSize)
	secLayerContext.Handle2WriterChan = make(chan *SecResponse, secLayerContext.secLayerConf.Handle2WriterChanSize)
	secLayerContext.HistoryMap = make(map[int]*UserBuyHistory, 100000)
	secLayerContext.productCountMgr = NewProductCountMgr()

	logs.Debug("init all succ")
	return
}
