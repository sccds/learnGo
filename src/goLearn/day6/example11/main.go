package main

import "fmt"

func main()  {
	ch := make(chan int) // 创建一个无缓冲区channel

	// 新建goroutine
	go func() {
		for i:=0; i<5; i++ {
			ch <- i //往管道里写数据
		}
		close(ch)
	}()

	/*
	for {
		// 如果 ok, 说明管道没有关闭
		if num, ok := <-ch; ok {
			fmt.Println("num =", num)
		} else { // 可以监听到管道关闭
			break
		}
	}
	*/

	for num := range ch {
		fmt.Println("num =", num)
	}
}
