package main

import (
	"regexp"
	"fmt"
)

func main()  {
	buf := "abc azc a7c 888 a9c tac axxc acvv"
	// 解析规则
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("regexp error")
		return
	}
	// 如果解析成功，根据规则提取信息
	result := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println(result)
}
