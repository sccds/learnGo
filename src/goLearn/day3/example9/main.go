package main

import "fmt"

func main()  {
	array := []int{0,1,2,3,4,5,6,7,8,9}

	s1 := array[:] // [0:len(array):len(array)]
	fmt.Println("s1 =", s1, len(s1), cap(s1))

	data := array[1]
	fmt.Println("data =", data)

	s2 := array[3:6:7]
	fmt.Println("s2 =", s2, len(s2), cap(s2))

	s3 := array[:6]
	fmt.Println("s3 =", s3, len(s3), cap(s3))

	s4 := array[3:]
	fmt.Println("s4 =", s4, len(s4), cap(s4))


	a := []int{0,1,2,3,4,5,6,7,8,9}
	s11 := a[2:5]
	s11[1] = 666
	fmt.Println("s11 =", s11, len(s11), cap(s11))
	fmt.Println("a =", a)

	s22 := s11[2:7]
	s22[2] = 777
	fmt.Println("a =", a)
}
