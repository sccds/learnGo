package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person
	id int
	addr string
	name string
}


func main()  {
	// 声明结构体
	var s Student
	s.name = "mike"  // 就近原则，操作student里面的Name
	s.Person.name = "li si"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	fmt.Printf("s = %+v\n", s)
}
