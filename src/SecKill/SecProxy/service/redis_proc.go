package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

func WriteHandle() {
	// 往 chan secReqChan 里面写请求数据，
	// 业务逻辑层来处理，处理之后写入 redis, 另一个再来读取
	for {
		req := <-secKillConf.SecReqChan
		conn := secKillConf.proxy2LayerRedisPool.Get()
		data, err := json.Marshal(req)
		if err != nil {
			logs.Error("json marshal failed, err: %v, req: %v", err, req)
			conn.Close()
			continue
		}
		_, err = conn.Do("LPUSH", "sec_queue", string(data))
		if err != nil {
			logs.Error("LPUSH failed, err: %v, req: %v", err, req)
			conn.Close()
			continue
		}
		conn.Close()
	}
}

func ReadHandle() {
	for {
		conn := secKillConf.proxy2LayerRedisPool.Get()
		reply, err := conn.Do("RPOP", "recv_queue")
		data, err := redis.String(reply, err)
		if err == redis.ErrNil {
			time.Sleep(time.Second * 2)
			conn.Close()
			continue
		}
		if err != nil {
			logs.Error("RPOP failed, err: %v", err)
			conn.Close()
			continue
		}
		logs.Debug("rpop from redis succ, data: %s", string(data))

		var result SecResult
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			logs.Error("json.Unmarshal error: %v", err)
			conn.Close()
			continue
		}
		userKey := fmt.Sprintf("%d_%d", result.UserId, result.ProductId)

		secKillConf.UserConnMapLock.Lock()
		resultChan, ok := secKillConf.UserConnMap[userKey]
		secKillConf.UserConnMapLock.Unlock()
		if !ok {
			conn.Close()
			logs.Warn("user not found: %v", userKey)
			continue
		}
		resultChan <- &result
		conn.Close()
	}
}
