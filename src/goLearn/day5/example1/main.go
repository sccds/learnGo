package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func MyDiv(a, b int) (result int, err error)  {
	err = nil
	if b == 0 {
		//err = errors.New("分母不能为0")
		err = fmt.Errorf("分母不能为0")
	} else {
		result = a / b
	}
	return
}

func main()  {
	err1 := fmt.Errorf("%s", "this is a normal err")
	fmt.Println(err1)

	err2 := errors.New("this is a normal err2")
	fmt.Println(err2)

	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println(result)
	}
}
