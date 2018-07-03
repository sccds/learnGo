package main

import (
	"fmt"
	"time"
)


func main()  {
	// 匿名函数
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子线程 i =", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for  {
		i++
		fmt.Println("main i =", i)
		time.Sleep(time.Second)
		if i == 2{
			break
		}
	}
}
