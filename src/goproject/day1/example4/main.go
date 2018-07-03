package main

import "fmt"

var (
	aa = 22
	ss = "kkk"
	bb = true
)

func variableZeroValue()  {
	var a int
	var b string
	var c bool
	a = 10
	b = "test"
	c = true
	fmt.Printf("%d %q %v\n", a, b, c)
}

func variableInitValue()  {
	var a, b = 2, 3
	var s string = "db"
	fmt.Println(a, b, s)
}

func variableShorter()  {
	a, b, s, c := 2, 3, "db", true
	fmt.Println(a, b, s, c)
}

func main()  {
	variableZeroValue()
	variableInitValue()
	variableShorter()
}
