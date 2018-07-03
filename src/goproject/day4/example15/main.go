package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}


func (p *Person) PrintInfo()  {
	fmt.Printf("name = %s, sex = %c, age = %d\n", p.name, p.sex, p.age)
}

// 学生，集成Person字段，成员和方法都会继承
type Student struct {
	Person  // 匿名
	id int
	addr string
}

func (p *Student) PrintInfo()  {  // 如果和匿名字段方法一样，叫方法重写
	fmt.Printf("addr = %s, sex = %c, age = %d\n", p.name, p.sex, p.age)
}

func main()  {
	s := Student{Person{"mike", 'm', 18}, 666, "bj"}
	s.PrintInfo()
	s.Person.PrintInfo()
}
