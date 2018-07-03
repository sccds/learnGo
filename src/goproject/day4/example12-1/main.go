package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

// 带有接收者的函数，叫方法
func (temp Person) PrintInfo()  {
	fmt.Println("temp =", temp)
}

// 通过一个函数，给成员赋值. 赋值的时候，接受者要是指针
func (p *Person) SetInfo(n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
}

func main()  {
	P := Person{"mike", 'm', 18}
	P.PrintInfo()

	var p2 Person
	p2.SetInfo("go", 'f', 22)
	p2.PrintInfo()
}
