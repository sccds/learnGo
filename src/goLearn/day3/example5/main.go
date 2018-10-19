package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// 数组比较，比较的两个数组类型要一致，支持 ==, !=
	a := [5]int{1,2,3,4,5}
	b := [5]int{1,2,3,4,5}
	c := [5]int{1,2,3}
	fmt.Println("a == b:", a == b)
	fmt.Println("a == c:", a == c)

	// 同类型的数组可以赋值
	var d [5]int
	d = a
	fmt.Println("d =", d)

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano()) // 以当前系统时间作为种子，使得每次的随机数不同
	for i := 0; i < 5; i++  {
		// fmt.Println("rand =", rand.Int())
		fmt.Println("rand =", rand.Intn(100)) // 产生0-100的值
	}
}
