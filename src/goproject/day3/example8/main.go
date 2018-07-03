package main

import "fmt"

func main()  {
	// 切片和数组的区别
	// 数组中长度固定，是常量；数组不能修改长度， len cap 固定
	a := [5]int{}
	fmt.Printf("s = %v, len = %d, cap = %d\n", a, len(a), cap(a))

	s := []int{}
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
	s = append(s, 13)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	// 切片创建
	// 1. 自动推导类型，同时初始化
	s1 := []int{1,2,3,4}
	fmt.Printf("s1 = %v, type = %T\n", s1, s1)

	// 2. 借助 make(切片类型，长度，容量)
	s2 := make([]int, 5, 10)
	fmt.Printf("s2 = %v, type = %T, len = %d, cap = %d\n", s2, s2, len(s2), cap(s2))

	// 3. 借助 make(切片类型，长度) 没有指定容量，容量与len一致
	s3 := make([]int, 5)
	fmt.Printf("s3 = %v, type = %T, len = %d, cap = %d\n", s3, s3, len(s3), cap(s3))

}
