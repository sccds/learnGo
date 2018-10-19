package test

import "fmt"

type stu struct {
	id int
}

type Stu struct {
	Id int
}

func myFunc()  {
	fmt.Println("this is my function")
}

func MyFunc()  {
	fmt.Println("this is my MyFunc")
}