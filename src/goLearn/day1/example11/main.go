package main

import (
	"math/rand"
	"fmt"
)

func main()  {
	var n int
	n = rand.Intn(100)
	for {
		var input int
		fmt.Scanf("%d", &input)
		flag := false
		switch  {
		case input == n:
			fmt.Println("your are correct")
			flag = true
		case input > n:
			fmt.Println("bigger")
		case input < n:
			fmt.Println("small")
		}
		if flag {
			break
		}
	}
}
