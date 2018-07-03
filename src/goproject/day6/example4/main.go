package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("ccccccccccc")
	//return // 终止此函数 ddd不会打印
	runtime.Goexit() // 终止所在协程， 不会打印 bbbb
	fmt.Println("dddddddddddd")
}

func main()  {
	go func() {
		fmt.Println("aaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbb")
	}()

	// 写个死循环，目的是不让程序终止
	for {}
}
