package main

import (
	"logAgent/kafka"
	"logAgent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = kafka.SendToKafka(msg.Msg, msg.Topic)
		if err != nil {
			logs.Error("send to kafka failed, err: %v", err)
			time.Sleep(time.Second)
			continue
		}
	}
}
