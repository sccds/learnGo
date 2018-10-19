package main

import "fmt"

func main()  {
	// 声明一个 map
	var m1 map[int]string
	fmt.Println(m1, m1 == nil)

	// 对应map只有len, 没有cap
	fmt.Println("len =", len(m1))

	// 通过 make 创建 map
	m2 := make(map[int]string)
	fmt.Println("m2 =", m2, "len =", len(m2))

	// 通过 make 创建 map, 可以指定长度，如果没有数据，会显示0. 如果长度超过初始化长度，会自动扩容
	m3 := make(map[int]string, 2)
	m3[1] = "mike"
	m3[2] = "c++"
	m3[3] = "go"
	fmt.Println("m3 =", m3, "len =", len(m3))

	// 键值唯一
	m4 := map[int]string{1: "mike", 2: "c++", 3: "go"}
	fmt.Println("m4 =", m4, "len =", len(m4))
}
