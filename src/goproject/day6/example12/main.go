package main

import "fmt"

// 此通道只能写不能读
func producer(out chan<- int)  {
	for i:=0; i<10; i++ {
		out <- i * i
	}
	close(out)
}

// channel 只能读不能写
func consumer(in <-chan int)  {
	for num := range in {
		fmt.Println("num =", num)
	}
}


func main()  {
	// 创建一个双向管道
	ch := make(chan int)

	// 生产者，生成数字，写入channel
	go producer(ch)

	// 消费者，从channel读取内容打印
	consumer(ch)
}
