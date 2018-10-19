package main

import (
	"fmt"
	"time"
)

func main()  {
	// 创建一个有缓冲区管道
	ch := make(chan int, 3)
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程: i = %d\n", i)
			ch <- i
			fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
		}
	}()
	// 延时
	time.Sleep(2*time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //读取管道内容，没有内容前阻塞
		fmt.Println("num =", num)
	}
}
