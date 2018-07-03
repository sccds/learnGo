package main

import "fmt"

func Print(n int)  {
	for i := 0; i < n + 1; i++  {
		for j := 0; j < i + 1; j++  {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

func main()  {
	Print(10)
}