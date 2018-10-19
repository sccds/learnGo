package main

import "fmt"

func main()  {
	data := []int{0,1,2,3,4,5,6,7,8,9}
	s1 := data[8:]
	s2 := data[:5]
	fmt.Println(s1, data)
	fmt.Println(s2, data)
	copy(s2, s1)
	fmt.Println(data)
}
