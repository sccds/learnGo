package main

import (
	"fmt"
	"time"
)

func Printer(str string)  {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1()  {
	Printer("hello")
}

func person2()  {
	Printer("world")
}

func main()  {
	go person1()
	go person2()
	for{}
}
