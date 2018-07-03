package main

import "fmt"

type mystr string // 自定义类型

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person  // 结构体匿名字段
	int		// 基础类型匿名字段
	mystr   // 自定义类型
}


func main()  {
	// 声明结构体
	s := Student{Person{"mike", 'm', 18}, 666, "go"}
	fmt.Printf("s = %+v\n", s)

	s.Person = Person{"go", 'm', 22}
	fmt.Println(s.name, s.age, s.int, s.mystr)
	fmt.Println(s.Person, s.int, s.mystr)
}
