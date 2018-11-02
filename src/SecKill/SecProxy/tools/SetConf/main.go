package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type SecInfoConf struct {
	ProductId int
	StartTime int
	EndTime   int
	Status    int
	Total     int
	Left      int
}

const (
	EtcdKey = "/zcz/secskill/product"
)

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect, err", err)
	}
	fmt.Println("connect succ")
	defer cli.Close()

	var secInfoConfArr []SecInfoConf

	secInfoConfArr = append(
		// normal product sec in progress
		secInfoConfArr, SecInfoConf{
			ProductId: 1032,
			StartTime: 1539936089,
			EndTime:   1640523153,
			Status:    0,
			Total:     10000,
			Left:      10000,
		},
	)

	secInfoConfArr = append(
		// product sec end
		secInfoConfArr, SecInfoConf{
			ProductId: 1042,
			StartTime: 1539936089,
			EndTime:   1540521968,
			Status:    0,
			Total:     10000,
			Left:      10000,
		},
	)
	secInfoConfArr = append(
		// product sec not start
		secInfoConfArr, SecInfoConf{
			ProductId: 1052,
			StartTime: 1640521968,
			EndTime:   1740108889,
			Status:    0,
			Total:     900,
			Left:      900,
		},
	)
	data, err := json.Marshal(secInfoConfArr)
	if err != nil {
		fmt.Println("json failed, err: ", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put data failed", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get etcd failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}
