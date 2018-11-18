package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func updateSecProductInfo(conf *SecLayerConf, secProductInfo []SecProductInfoConf) {
	var temp map[int]*SecProductInfoConf = make(map[int]*SecProductInfoConf, 1024)
	for _, v := range secProductInfo {
		productInfo := v
		productInfo.secLimit = &SecLimit{}
		temp[v.ProductId] = &productInfo
	}
	secLayerContext.RWSecProductLock.Lock()
	conf.SecProductInfoMap = temp
	secLayerContext.RWSecProductLock.Unlock()
}

func initSecProductWatcher(conf *SecLayerConf) {
	go watchSecProductKey(conf)
}

func watchSecProductKey(conf *SecLayerConf) {
	key := conf.EtcdConfig.EtcdSecProductKey
	logs.Debug("begin watch key: %s", key)
	var err error
	for {
		rch := secLayerContext.etcdClient.Watch(context.Background(), key)
		var secProductInfo []SecProductInfoConf
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] config deleted", key)
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &secProductInfo)
					if err != nil {
						logs.Error("key[%s] unmarshal error, err: %v", key, err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd, %s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", secProductInfo)
				updateSecProductInfo(conf, secProductInfo)
			}
		}
	}
}

func loadProductFromEtcd(conf *SecLayerConf) (err error) {
	logs.Debug("start get product from etcd")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := secLayerContext.etcdClient.Get(ctx, conf.EtcdConfig.EtcdSecProductKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err: %v", conf.EtcdConfig.EtcdSecProductKey, err)
		return
	}
	logs.Debug("get product from etcd succ, resp: %v", resp)

	var secProductInfo []SecProductInfoConf
	for k, v := range resp.Kvs {
		logs.Debug("key[%v] val[%v]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("unmarshal sec product info failed, err: %v", err)
			return
		}
		logs.Debug("sec info conf is [%v]", secProductInfo)
	}
	updateSecProductInfo(conf, secProductInfo)
	logs.Debug("update product info succ, data: %v", secProductInfo)

	initSecProductWatcher(conf)
	logs.Debug("init etcd watcher succ")
	return
}
