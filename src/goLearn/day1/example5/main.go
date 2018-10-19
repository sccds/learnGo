package main

import "fmt"

func main()  {
	const LENGHT int = 10
	const WIDTH = 5
	var area int
	area = LENGHT * WIDTH

	fmt.Printf("area: %d\n", area)

	const a, b, c = 1, false, "str"
	fmt.Println(a, b, c)

	const (
		Unknown = iota
		Female
		Male
	)
	fmt.Println(Unknown, Female, Male)
}
