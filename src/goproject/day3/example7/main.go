package main

import "fmt"

func main()  {
	a := []int{1, 2, 3, 4, 5}
	s := a[0:3:5]
	fmt.Println("s =", s)
	fmt.Println("len(s) =", len(s)) // length = high - low
	fmt.Println("cap(s) =", cap(s)) // capacity = max - low

	s = a[1:4:5]
	fmt.Println("s =", s)
	fmt.Println("len(s) =", len(s)) // length = high - low
	fmt.Println("cap(s) =", cap(s)) // capacity = max - low
}