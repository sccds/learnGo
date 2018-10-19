package main

import (
	"time"
	"fmt"
)

func main()  {
	/*
	// 创建一个定时器，设置时间为2s，2s后往time通道写内容
	timer := time.NewTimer(2*time.Second)
	fmt.Println("当前时间", time.Now())

	// 2s 后往 timer.C 写数据，有数据后，可以读取
	t := <-timer.C  // 如果channel 里没有数据，阻塞
	fmt.Println("t =", t)
	*/


	/*
	// 验证，只能用一次
	timer := time.NewTimer(1*time.Second)
	for{
		<- timer.C
		fmt.Println("时间到")
	}
	*/

	// 延迟功能
	<- time.After(2*time.Second) // 定时2s，阻塞2s, 2s后产生一个事件，往chan写内容
	fmt.Println("时间到")
}
