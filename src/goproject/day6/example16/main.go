package main

import (
	"fmt"
)

// ch 只写， quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		// 监听 channel 数据流动
		select {
		case ch <- x:
			x, y = y, x + y
		case flag := <-quit:
			fmt.Println("flag =", flag)
			return
		}
	}
}

func main()  {
	ch := make(chan int) // 数据通道
	quit := make(chan bool) // 程序是否结束

	// 消费者，从chan里面读取内容
	go func() {
		for i := 0; i < 2; i++ {
			num := <- ch
			fmt.Println("num =", num)
		}
		quit <- true
	}()

	// 生产者，产生数据，写入channel
	fibonacci(ch, quit)

}
