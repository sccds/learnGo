package main

import (
	"runtime"
	"fmt"
)

func main()  {
	n := runtime.GOMAXPROCS(4) // 指定1个核工作
	fmt.Println("n =", n)

	for  {
		go fmt.Print("1")
		fmt.Print("0")
	}
}
