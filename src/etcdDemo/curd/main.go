package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	//SetLogConfToEtcd()
	EtcdExample()
}

type LogConf struct {
	Path  string
	Topic string
}

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()

	var logConfArr []LogConf
	logConfArr = append(
		logConfArr,
		LogConf{
			Path:  "/Users/xliu/Documents/51CTO_Golang/learnGo/src/logAgent/logs/logAgent.log",
			Topic: "nginx_log",
		},
	)
	logConfArr = append(
		logConfArr,
		LogConf{
			Path:  "/Users/xliu/Documents/51CTO_Golang/learnGo/src/logAgent/logs/err.log",
			Topic: "nginx_log_err",
		},
	)
}

func EtcdExample() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err", err)
		return
	}
	fmt.Println("conenct succ")
	defer cli.Close()
	// etcd 赋值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/conf", "sample_value")
	_, err = cli.Put(ctx, "/logagent/conf/aa", "sample_value2")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	// etcd 取值
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
