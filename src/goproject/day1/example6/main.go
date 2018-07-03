package main

import "fmt"

func main()  {
	var a int = 50
	var b bool
	c := 'a'
	fmt.Printf("%v\n", a)
	fmt.Printf("%v, %#v\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("90%%\n")
	fmt.Printf("%b, %f, %q, %x\n", 100, 120.00, "test", 390022)
	fmt.Printf("%p\n", &a)

	str := fmt.Sprintf("a = %d", a)
	fmt.Printf("%q\n", str)
}
