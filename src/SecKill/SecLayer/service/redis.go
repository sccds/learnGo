package service

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

func initRedisPool(redisConf RedisConf) (pool *redis.Pool, err error) {
	pool = &redis.Pool{
		MaxIdle:     redisConf.RedisMaxIdle,
		MaxActive:   redisConf.RedisMaxActive,
		IdleTimeout: time.Duration(redisConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisConf.RedisAddr)
		},
	}
	conn := pool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err: %v", err)
		return
	}
	return
}

func initRedis(conf *SecLayerConf) (err error) {
	secLayerContext.proxy2LayerRedisPool, err = initRedisPool(conf.Proxy2LayerRedis)
	if err != nil {
		logs.Error("init proxy2Layer redis pool failed, err: %v", err)
		return
	}
	secLayerContext.layer2ProxyRedisPool, err = initRedisPool(conf.Layer2ProxyRedis)
	if err != nil {
		logs.Error("init layer2Proxy redis pool failed, err: %v", err)
		return
	}
	return
}

func RunProcess() (err error) {
	// 读
	for i := 0; i < secLayerContext.secLayerConf.ReadLayer2ProxyGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleReader()
	}
	// 写
	for i := 0; i < secLayerContext.secLayerConf.WriteProxy2LayerGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleWriter()
	}
	// 处理
	for i := 0; i < secLayerContext.secLayerConf.HandleUserGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleUser()
	}
	logs.Debug("all process goroutine start")
	secLayerContext.waitGroup.Wait()
	logs.Debug("wait all goroutine exit")
	return
}

func HandleReader() {
	for {
		conn := secLayerContext.proxy2LayerRedisPool.Get()
		for {
			ret, err := conn.Do("blpop", secLayerContext.secLayerConf.Proxy2LayerRedis.RedisQueueName, 0)
			if err != nil {
				logs.Error("pop from queue failed, err: %v", err)
				break
			}
			tmp, ok := ret.([]interface{})
			if !ok || len(tmp) != 2 {
				logs.Error("pop from queue failed, err: %v", err)
				continue
			}
			data, ok := tmp[1].([]byte)
			logs.Debug("pop from queue, data: %s", string(data))

			var req SecRequest
			err = json.Unmarshal([]byte(data), &req)
			if err != nil {
				logs.Debug("unmarshal to secrequest failed, err: %v", err)
				continue
			}
			// 如果时间超过timeout时间，不做秒杀处理
			now := time.Now().Unix()
			if now-req.AccessTime.Unix() >= int64(secLayerContext.secLayerConf.MaxRequestWaitTimeout) {
				logs.Warn("req[%v] is expired", req)
				continue
			}

			timer := time.NewTicker(time.Microsecond * time.Duration(secLayerContext.secLayerConf.SendToHandleChanTimeout))
			select {
			case secLayerContext.Read2HandleChan <- &req:
			case <-timer.C:
				logs.Warn("send to handle chan timeout, req: %v", req)
				break
			}

		}
		conn.Close()
	}
}

func HandleWriter() {
	for res := range secLayerContext.Handle2WriterChan {
		err := sendToRedis(res)
		if err != nil {
			logs.Error("send to redis failed, err: %v, res: %v", err, res)
			continue
		}
	}
}

func sendToRedis(res *SecResponse) (err error) {
	data, err := json.Marshal(res)
	if err != nil {
		logs.Error("json marshal data err: %v", err)
		return
	}
	conn := secLayerContext.layer2ProxyRedisPool.Get()
	_, err = conn.Do("rpush", secLayerContext.secLayerConf.Layer2ProxyRedis.RedisQueueName, string(data))
	if err != nil {
		logs.Warn("rpush to redis failed, err: %v", err)
		return
	}
	return
}

func HandleUser() {
	for req := range secLayerContext.Read2HandleChan {
		logs.Debug("begin process request: %v", req)
		res, err := HandleSecKill(req)
		if err != nil {
			logs.Warn("process request %v failed, err: %v", req, err)
			res = &SecResponse{
				Code: ErrServiceBusy,
			}
		}
		logs.Debug("handel user res: %v", res)
		timer := time.NewTicker(time.Microsecond * time.Duration(secLayerContext.secLayerConf.SendToWriterChanTimeout))
		select {
		case secLayerContext.Handle2WriterChan <- res:
		case <-timer.C:
			logs.Warn("send to response chan timeout, res: %v", res)
			break
		}
	}
	return
}

func HandleSecKill(req *SecRequest) (res *SecResponse, err error) {
	secLayerContext.RWSecProductLock.RLock()
	defer secLayerContext.RWSecProductLock.RUnlock()
	res = &SecResponse{}
	res.UserId = req.UserId
	res.ProductId = req.ProductId
	product, ok := secLayerContext.secLayerConf.SecProductInfoMap[req.ProductId]
	if !ok {
		logs.Error("not found product: %v", req.ProductId)
		res.Code = ErrNotFoundPorduct
		return
	}
	// 售罄状态
	if product.Status == ProductStatusSoldout {
		res.Code = ErrSoldout
		return
	}
	logs.Debug("product check succ")

	// 限速
	now := time.Now().Unix()
	alreadySoldCount := product.secLimit.Check(now)
	logs.Debug("product soldlimit: %d", product.SoldLimit)
	if alreadySoldCount >= product.SoldLimit {
		res.Code = ErrRetry
		return
	}
	logs.Debug("speed limit succ")

	// 一个用户一个产品的购买量
	secLayerContext.HistoryMapLock.Lock()
	userHistory, ok := secLayerContext.HistoryMap[req.UserId]
	if !ok {
		userHistory = &UserBuyHistory{
			history: make(map[int]int, 16),
		}
		secLayerContext.HistoryMap[req.UserId] = userHistory
	}
	historyCount := userHistory.GetProductBuyCount(req.ProductId)
	secLayerContext.HistoryMapLock.Unlock()

	if historyCount >= product.OnePersonBuyLimit {
		res.Code = ErrAlreadyBuy
		return
	}
	logs.Debug("num one user per product check succ")

	// 一个产品总共有多少个
	curSoldCount := secLayerContext.productCountMgr.Count(req.ProductId)
	if curSoldCount >= product.Total {
		product.Status = ProductStatusSoldout
		return
	}

	// 概率随机抽奖, 小于某概率才能进行购买
	logs.Debug("before lottery")
	curRate := rand.Float64()
	logs.Debug("curRate:%v product:%v count:%v total:%v\n", curRate, product.BuyRate, curSoldCount, product.Total)
	if curRate > product.BuyRate {
		res.Code = ErrRetry
		return
	}
	userHistory.Add(req.ProductId, 1)
	secLayerContext.productCountMgr.Add(req.ProductId, 1)
	logs.Debug("lottery succ")

	// 购买成功，返回token，用户使用token可购买产品
	// 用户id, 产品id, 当前时间, 密钥 返回
	res.Code = ErrSecKillSucc
	tokenData := fmt.Sprintf("userId=%d&productId=%dtimestamp=%d&security=%s", req.UserId, req.ProductId, now, secLayerContext.secLayerConf.TokenPasswd)
	logs.Debug(tokenData)
	res.Token = fmt.Sprintf("%x", md5.Sum([]byte(tokenData)))
	logs.Debug(res.Token)
	res.TokenTime = now
	return
}
