package main

import "fmt"

func del(m map[int]string)  {
	delete(m, 1)
}

func main()  {
	m1 := map[int]string{1: "mike", 2:"yoyo"}
	fmt.Println(m1)

	// 赋值已经存在的key,会修改内容
	m1[1] = "c++"
	m1[3] = "go" // 追加，map底层自动扩容，和append类似
	fmt.Println(m1)

	// 遍历
	for key, value := range m1 {
		fmt.Printf("%d ===> %s\n", key, value)
	}

	// 判断一个 key 是否存在， 第一个返回值是key所对应value, 第二个返回值是key是否存在
	value, ok := m1[1]
	if ok {
		fmt.Println("m1[1] =", value)
	} else {
		fmt.Println("key not found")
	}

	// 删除
	//delete(m1, 1)
	//fmt.Println(m1)

	// map 作为函数参数，是引用传递
	del(m1)
	fmt.Println(m1)
}
