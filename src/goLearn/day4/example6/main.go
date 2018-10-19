package main

import "fmt"

type Student struct {
	id int
	name string
	gender byte
	age int
	addr string
}

func test1(s Student) {
	s.id = 666

}

func test2(s *Student)  {
	s.id = 666
}

func main()  {
	s1 := Student{id: 1, name: "mike", gender: 'm', age: 18, addr: "bj"}
	s2 := Student{id: 1, name: "mike", gender: 'm', age: 18, addr: "bj"}
	s3 := Student{id: 2, name: "mike", gender: 'm', age: 18, addr: "bj"}

	fmt.Println("s1 == s2 ?", s1 == s2)
	fmt.Println("s1 == s3 ?", s1 == s3)

	// 同类型的两个结构体变量可以相互赋值
	var tmp Student
	tmp = s3
	fmt.Println("before func:", tmp)

	///test1(tmp)
	test2(&tmp)
	fmt.Println("after func:", tmp)
}
