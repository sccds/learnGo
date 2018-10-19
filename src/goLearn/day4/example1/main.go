package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int)  {
	// 设置种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++  {
		s[i] = rand.Intn(100) // 100以内随机数
	}
}

// 冒泡排序
func BubbleSort(s []int)  {
	n := len(s)
	for i := 0; i < n-1; i++  {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main()  {
	// 初始化切片
	n := 10

	// 创建切片， len 10
	s := make([]int, n)

	InitData(s) // 初始化数组

	fmt.Println("before sort:", s)

	BubbleSort(s)

	fmt.Println("after sort:", s)

}