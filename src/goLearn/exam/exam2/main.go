package main

import "fmt"

func max(num1, num2 int) int  {
	if num1 >= num2 {
		return num1
	}
	return num2
}

func main()  {
	fmt.Println(max(1, 2))
}
