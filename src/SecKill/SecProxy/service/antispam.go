package service

import (
	"fmt"
	"sync"

	"github.com/astaxie/beego/logs"
)

type SecLimitMgr struct {
	UserLimitMap map[int]*Limit
	IpLimitMap   map[string]*Limit
	lock         sync.Mutex
}

func antiSpam(req *SecRequest) (err error) {
	_, ok := secKillConf.idBlackMap[req.UserId]
	if ok {
		err = fmt.Errorf("invalid request")
		logs.Error("user_id [%v] is blocked by black user_id", req.UserId)
		return
	}

	_, ok = secKillConf.ipBlackMap[req.ClientAddr]
	if ok {
		err = fmt.Errorf("invalid  request")
		logs.Error("user_id [%v] with ip [%v] is blocked by black ip", req.UserId, req.ClientAddr)
		return
	}

	secKillConf.secLimitMgr.lock.Lock()

	// user id 频率控制
	limit, ok := secKillConf.secLimitMgr.UserLimitMap[req.UserId]
	if !ok {
		limit = &Limit{
			secLimit: &SecLimit{},
			minLimit: &MinLimit{},
		}
		secKillConf.secLimitMgr.UserLimitMap[req.UserId] = limit
	}
	secIdCount := limit.secLimit.Count(req.AccessTime.Unix())
	minIdCount := limit.minLimit.Count(req.AccessTime.Unix())

	// ip 频率控制
	limit, ok = secKillConf.secLimitMgr.IpLimitMap[req.ClientAddr]
	if !ok {
		limit = &Limit{
			secLimit: &SecLimit{},
			minLimit: &MinLimit{},
		}
		secKillConf.secLimitMgr.IpLimitMap[req.ClientAddr] = limit
	}
	secIpCount := limit.secLimit.Count(req.AccessTime.Unix())
	minIpCount := limit.minLimit.Count(req.AccessTime.Unix())

	secKillConf.secLimitMgr.lock.Unlock()

	if secIdCount > secKillConf.AccessLimitConf.UserSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}
	if minIdCount > secKillConf.AccessLimitConf.UserMinAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}
	if secIpCount > secKillConf.AccessLimitConf.IpSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}
	if minIpCount > secKillConf.AccessLimitConf.IpMinAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	return
}
