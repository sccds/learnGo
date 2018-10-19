package main

import "fmt"

func main()  {
	s1 := []int{}
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))

	// 在原切片末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))

	s2 := []int{1,2,3}
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	s := make([]int, 0, 3)
	oldCap := cap(s)
	for i:=0; i<20; i++  {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("current length: %d, cap: %d ====> %d\n", len(s), oldCap, newCap)
			oldCap = newCap
		}
	}

	// copy
	srcSlice := []int{1,2,3,4,5,6,7}
	dstSlice := []int{6,6,6,6}
	copy(dstSlice, srcSlice)
	fmt.Println(dstSlice)

}
