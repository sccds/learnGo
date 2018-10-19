package main

import "fmt"

type Student struct {
	id		int
	name	string
	gender	byte // 字符类型
	age		int
	addr	string
}


func main()  {
	// 顺序初始化，每个成员都要初始化
	var s1 Student = Student{id: 1, name: "mike", gender: 'm', age: 18, addr: "bj"}
	fmt.Println(s1)

	// 指定成员初始化，没有初始化的成员，都是0
	var s2 Student = Student{name: "zhang san", addr: "bj"}
	fmt.Println(s2)

	// 指针变量初始化，别忘了加&
	var p1 *Student = &Student{id: 1, name: "mike", gender: 'm', age: 18, addr: "bj"}
	fmt.Println(*p1)

	p2 := &Student{name: "zhang san", addr: "bj"}
	fmt.Printf("p2 type is %T, content: %v", p2, *p2)
}
