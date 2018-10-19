package main

import "fmt"

// 数组作为函数参数，值传递，是参数组每个元素给形参数组拷贝一份
// p 指向实现数组a, 是指向数组，是指针数组， *p 等价于 a
func modify(p *[6]int)  {
	(*p)[0] = 66666
	fmt.Println("modify a in func =", *p)
}

func main()  {
	a := [6]int{1,2,3,4,5,6}

	modify(&a)
	fmt.Println("modify a =", a)
}
