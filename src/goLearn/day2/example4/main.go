package main

import "fmt"

func swap(a *int, b *int)  {
	//*a, *b = *b, *a
	temp := *a
	*a = *b
	*b = temp
}

func main()  {
	a := 5
	b := 10
	fmt.Println(a, b)

	swap(&a, &b)
	fmt.Println(a, b)
}
