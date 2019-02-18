package main

import (
	"context"
	"encoding/json"
	"fmt"
	"logAgent/tailf"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const (
	EtcdKey = "/logagent/config/127.0.0.1"
)

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()

	var logConfArr []tailf.CollectConf
	logConfArr = append(
		logConfArr,
		tailf.CollectConf{
			LogPath: "/Users/xliu/Documents/51CTO_Golang/learnGo/src/logAgent/logs/logAgent.log",
			Topic:   "test",
		},
	)
	logConfArr = append(
		logConfArr,
		tailf.CollectConf{
			LogPath: "",
			Topic:   "nginx_log_err",
		},
	)
	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}
