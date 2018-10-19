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
}


func main()  {
	// 顺序初始化
	var s1 Student = Student{Person: Person{name: "mike", sex: 'm', age: 18}, id: 1, addr: "bj"}
	fmt.Println(s1)

	// 自动推导类型
	s2 := Student{Person: Person{name: "mike", sex: 'm', age: 18}, id: 1, addr: "bj"}
	fmt.Println(s2)
	fmt.Printf("s2 = %+v\n", s2)

	// 部分初始化，指定成员初始化，没有初始化的自动赋值为0
	s3 := Student{id: 1}
	fmt.Printf("s3 = %+v\n", s3)

	s4 := Student{Person: Person{name: "mike"}, id:1}
	fmt.Printf("s4 = %+v\n", s4)

	// 成员操作
	s2.name = "zhang san"
	s2.id = 100
	fmt.Printf("s2 = %+v\n", s2)

	s2.Person = Person{"go", 'm', 18}
	fmt.Printf("s2 = %+v\n", s2)
}
