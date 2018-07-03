package main

import "fmt"

func testa()  {
	fmt.Println("aaaaaaaa")
}

func testb(x int)  {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(recover())
		}
	}()

	// 显示调用panic，导致程序中断
	//panic("")

	// 数组越界导致panic
	var a [10]int
	a[x] = 111
}

func testc() {
	fmt.Println("ccccccccc")
}

func main()  {
	testa()
	testb(20)
	testc()
}
