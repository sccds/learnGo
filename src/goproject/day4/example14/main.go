package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

// 接收者为普通变量，非指针，值语意，一份拷贝
func (p Person) SetInfoValue()  {
	fmt.Println("SetInfoValue")
}

// 接收者为指针变量，引用传递
func (p *Person) SetInfoPointer()  {
	fmt.Println("SetInfoPointer")
}


func main()  {
	// 指针变量，能调用哪些方法，这些方法称为方法集
	p := &Person{"mike", 'm', 18}
	p.SetInfoPointer()

	// 内部做了转换，先把指针 p 转换成 *p 然后再调用
	p.SetInfoValue()

	p1 := Person{"mike", 'm', 18}
	p1.SetInfoValue()

	// 内部做转换，先把 p 转换成 &p, 然后再调用
	p1.SetInfoPointer()
}
