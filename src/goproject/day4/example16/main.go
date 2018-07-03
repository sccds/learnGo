package main

import "fmt"

type Humaner interface {
	// 方法只有声明，没有实现，由别的类型实现
	sayHi()
}

type Student struct {
	name string
	id int
}

// Student 实现了此方法，必须用 *Student
func (tmp *Student) sayHi()  {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr string
	group string
}

func (tmp *Teacher) sayHi()  {
	fmt.Printf("Teacher[%s, %s] sayhi\n", tmp.addr, tmp.group)
}

type Mystr string

func (tmp *Mystr) sayHi()  {
	fmt.Printf("Mystr[%s] sayhi\n", *tmp)
}

// 定义一个普通的函数，函数的参数为接口类型。
// 只有一个函数，可以有不同的表现，多态
func WhoSayHi(i Humaner)  {
	i.sayHi()
}

func main()  {
	// 定义接口类型变量
	var i Humaner

	s := &Student{"mike", 666}
	i = s
	i.sayHi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayHi()

	var str Mystr = "hello world"
	i = &str
	i.sayHi()

	fmt.Println()

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	fmt.Println()

	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, i := range x {
		i.sayHi()
	}
}