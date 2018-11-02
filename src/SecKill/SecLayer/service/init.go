package service

import (
	"github.com/astaxie/beego/logs"
)

func InitSecLayer(config *SecLayerConf) (err error) {
	err = initRedis(config)
	if err != nil {
		logs.Error("init redis failed, err: %v", err)
		return
	}

	return
}
