package main

import (
	"strconv"
	"fmt"
)

func main()  {
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 16) // 第三个参数是 进制
	slice = strconv.AppendQuote(slice, "abcd")
	fmt.Println(string(slice))

	// 其他类型转换字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println(str)

	str = strconv.Itoa(666)
	fmt.Println(str)

	// 字符串转换成其他类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println(flag)
	} else {
		fmt.Println(err)
	}

	a, _ := strconv.Atoi("10")
	fmt.Println(a)
}
