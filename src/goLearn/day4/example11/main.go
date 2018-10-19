package main

import "fmt"

type mystr string // 自定义类型

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	*Person  // 指针类型
	id int
	addr string
}


func main()  {
	// 声明结构体
	s := Student{&Person{"mike", 'm', 18}, 666, "go"}
	fmt.Printf("s = %+v\n", s)
	fmt.Println(s.name, s.sex, s.age, s.id, s.addr)

	s.Person = &Person{"go", 'm', 22}

	// 先定义变量
	var s2 Student
	s2.Person = new(Person)
	s2.name = "go"
	s2.sex = 'm'
	s2.age = 18
	s2.id = 222
	s2.addr = "sz"
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.addr)
}
