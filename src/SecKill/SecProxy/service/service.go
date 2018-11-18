package service

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

func NewSecRequest() (secRequest *SecRequest) {
	secRequest = &SecRequest{
		ResultChan: make(chan *SecResult, 1),
	}
	return
}

func SecInfoList() (data []map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	for _, v := range secKillConf.SecProductInfoConfMap {
		item, _, err := SecInfoById(v.ProductId)
		if err != nil {
			logs.Error("get product_id[%d] failed, err: %v", v.ProductId, err)
			continue
		}
		logs.Debug("get product_id[%d], result[%v], all[%v], v[%v]", v.ProductId, item, secKillConf.SecProductInfoConfMap, v)
		data = append(data, item)
	}
	return
}

func SecInfo(productId int) (data []map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()
	item, code, err := SecInfoById(productId)
	if err != nil {
		return
	}
	data = append(data, item)
	return
}

func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	logs.Debug("get productInfoById: %v", productId)
	defer secKillConf.RwSecProductLock.RUnlock()
	v, ok := secKillConf.SecProductInfoConfMap[productId]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found product_id: %d", productId)
		return
	}

	start := false
	end := false
	status := "success"
	now := time.Now().Unix()

	if now-v.StartTime < 0 {
		// 秒杀还未开始
		start = false
		end = false
		status = "sec kill is not start"
		code = ErrActiveNotStart
	}

	if now-v.StartTime > 0 {
		start = true
	}

	if now-v.EndTime > 0 {
		start = false
		end = true
		status = "sec kill is already end"
		code = ErrActiveAlreadyEnd
	}

	if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusSaleOut {
		start = false
		end = true
		status = "product is sold out"
		code = ErrActiveSaleOut
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start_time"] = start
	data["end_time"] = end
	data["status"] = status

	return
}

// 用户验证
func userCheck(req *SecRequest) (err error) {
	found := false
	for _, refer := range secKillConf.ReferWhiteList {
		if refer == req.ClientReference {
			found = true
			break
		}
	}
	if !found {
		err = fmt.Errorf("invalid request")
		logs.Warn("user[%d] is rejected by refer, req[%v]", req.UserId, req)
		return
	}

	authData := fmt.Sprintf("%d:%s", req.UserId, secKillConf.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))
	if authSign != req.UserAuthSign {
		err = fmt.Errorf("invalid user cookie auth")
		return
	}
	return
}

func SecKill(req *SecRequest) (data map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	/*
		// 验证用户
		err = userCheck(req)
		if err != nil {
			code = ErrUserCheckAuthFailed
			logs.Warn("userid [%d] invalid, user check failed, req [%v]", req.UserId, req)
			return
		}
	*/

	// 防止刷屏
	err = antiSpam(req)
	if err != nil {
		code = ErrUserServiceBusy
		logs.Warn("userId [%d] invalid, spam check failed, req [%v]", req.UserId, req)
		return
	}
	data, code, err = SecInfoById(req.ProductId)
	if err != nil {
		logs.Warn("userid [%d] secInfoById failed, req [%v]", req.UserId, req)
		return
	}
	if code != 0 {
		logs.Warn("userid [%d] secInfoById failed, code [%d], req [%v]", req.UserId, code, req)
		return
	}

	userKey := fmt.Sprintf("%d_%d", req.UserId, req.ProductId)
	secKillConf.UserConnMap[userKey] = req.ResultChan

	// 符合条件的，写入 redis, 接入层 => 业务逻辑层
	secKillConf.SecReqChan <- req
	ticker := time.NewTicker(time.Second * 10) // 10s 没完成超时

	defer func() {
		ticker.Stop()
		secKillConf.UserConnMapLock.Lock()
		delete(secKillConf.UserConnMap, userKey)
		secKillConf.UserConnMapLock.Unlock()
	}()

	select {
	case <-ticker.C:
		code = ErrProcessTimeout
		err = fmt.Errorf("request timeout")
		return
	case <-req.CloseNotify:
		code = ErrClientClosed
		err = fmt.Errorf("client already closed")
		return
	case result := <-req.ResultChan:
		code = result.Code
		data["product_id"] = result.ProductId
		data["token"] = result.Token
		data["user_id"] = result.UserId
		return
	}
	return
}
