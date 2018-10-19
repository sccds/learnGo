package main

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"fmt"
)
func main()  {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://1.119.41.34:9200"))
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	ctx := context.Background()
	res, err := client.Delete().Index("test").Type("doc").Id("Ib9bg2QBOwO-n3_uszbK").Do(ctx)
	if err != nil {
		fmt.Println("delete err", err)
	}
	fmt.Println(res)
	client.Stop()
}
