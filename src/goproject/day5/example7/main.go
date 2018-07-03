package main

import (
	"regexp"
	"fmt"
)

func main()  {
	buf := "43.14 567 asdggg 1.23 7. 8.9 111111 6.66 7.8"

	// 1. +匹配前一个字符一次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("err")
	}
	// 提取关键信息
	//result := reg.FindAllString(buf, -1) // 返回一维切片
	result := reg.FindAllStringSubmatch(buf, -1)  // 返回二维切片
	fmt.Println(result)
}
