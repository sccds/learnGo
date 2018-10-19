package main

import (
	"goproject/day4/test"
	"fmt"
)

func main()  {
	test.MyFunc()
	var s test.Stu
	s.Id = 666

	fmt.Println("s =", s)
}
