package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Personer interface {
	Humaner
	sing(lrc string)
}

type Student struct {
	name string
	id int
}

// Student 中实现sayhi方法
func (tmp *Student) sayhi()  {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string)  {
	fmt.Printf("Student is singing: %s\n", lrc)
}

func main()  {
	// 定义一个接口类型的变量
	var i Personer
	s := &Student{"mike", 666}
	i = s
	i.sayhi()
	i.sing("lalala")

	var iPro Personer // 超级
	iPro = &Student{"mike", 666}
	var ii Humaner
	ii = iPro // 超级可以转换成子集，子集不能转换成超级
	ii.sayhi()
	// 子集中没有 sing 的方法
}
