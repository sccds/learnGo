package main

import (
	"fmt"
	"os"
)

func main()  {
	// os.Stdout.Close() // 关闭后无法输出，默认打开
	fmt.Println("are you ok")
	os.Stdout.WriteString("are you ok\n")

	var a int
	fmt.Println("请输入a: ")
	fmt.Scan(&a)  // 从标准输入设备中读取内容，把读取的内容放入a中
	fmt.Println("a =", a)
}
