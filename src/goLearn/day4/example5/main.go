package main

import "fmt"

type Student struct {
	id int
	name string
	gender byte
	age int
	addr string
}


func main()  {
	var s Student

	// 操作成员
	s.id = 1
	s.name = "mike"
	s.gender = 'm'
	s.age = 18
	s.addr = "bj"
	fmt.Println("s =", s)

	// 指针有合法指向后，才能操作成员
	var s1 Student
	var p1 *Student
	p1 = &s1
	// 通过指针操作成员 p1.id 和 (*p1).id 完全等价
	p1.id = 1
	(*p1).name = "mike"
	p1.gender = 'm'
	p1.age = 18
	p1.addr = "bj"
	fmt.Println("p1 =", p1)

	// 通过 new 申请一个结构体
	p2 := new(Student)
	p2.id = 1
	p2.name = "mike"
	p2.gender = 'm'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2 =", p2)
}
