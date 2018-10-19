package main

import "fmt"

func main()  {
	var a int = 0
	switch a {
	case 0:
		fmt.Printf("a is 0\n")
		//fallthrough
	case 10:
		fmt.Printf("a is 10\n")
	default:
		fmt.Printf("a is default")
	}
	g := grad(100)
	fmt.Println(g)
}

func grad(score int) string  {
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprint("wrong score: %d", score))
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "a"
	}
	return g
}