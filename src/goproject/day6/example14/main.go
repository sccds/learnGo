package main

import (
	"time"
	"fmt"
)

func main()  {
	/*
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("子协程可以打印，定时器时间到")
	}()
	timer.Stop() // 在主协程停止定时器，导致子协程获取不到值，无法打印

	for {
	}
	*/

	timer := time.NewTimer(3 * time.Second)
	timer.Reset(1*time.Second)
	<-timer.C
	fmt.Println("时间到")
}
