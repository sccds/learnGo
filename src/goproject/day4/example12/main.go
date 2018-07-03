package main

import "fmt"

// 面向过程
func Add01(a, b int) int {
	return a + b
}

// 面向对象 方法：给某个类型绑定一个函数
type long int

// temp 叫接收者，就是传递一个参数
func (temp long) Add02(a long) long  {
	return temp + a
}


func main()  {
	var result int
	result = Add01(1, 1)
	fmt.Println(result)

	// 使用，定一个变量
	var a long = 2

	// 调用方法
	r := a.Add02(2)
	fmt.Println(r)
}
