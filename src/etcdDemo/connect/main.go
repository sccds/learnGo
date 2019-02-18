package main

import (
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	etcd_client "github.com/coreos/etcd/clientv3"
)

func main() {
	cli, err := etcd_client.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()
}
