package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string)  {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1()  {
	Printer("hello")
	ch <- 888 // 管道里写数据，
}

func person2()  {
	<-ch // 从管道里读取数据，如果没有值数据就会阻塞
	Printer("world")
}

func main()  {
	go person1()
	go person2()
	for{}
}
