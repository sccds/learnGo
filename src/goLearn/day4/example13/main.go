package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

// 接收者为普通变量，非指针，值语意，一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
}

// 接收者为指针变量，引用传递
func (p *Person) SetInfoPointer(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
}


func main()  {
	s1 := Person{"go", 'm', 22}
	fmt.Println(s1)
	s1.SetInfoValue("mike", 'm', 18)
	fmt.Println(s1)
	s1.SetInfoPointer("mike", 'm', 18)
	fmt.Println(s1)
}
