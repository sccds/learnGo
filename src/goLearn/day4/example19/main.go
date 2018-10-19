package main

import "fmt"

type Student struct {
	name string
	id int
}

func main()  {
	i := make([]interface{}, 3)
	i[0] = 1 // int
	i[1] = "abc" // string
	i[2] = Student{"mike", 777}

	for index, data := range i {
		if value, ok := data.(int); ok {
			fmt.Printf("x[%d] type is int, content: %d\n", index, value)
		} else if value, ok := data.(string); ok {
			fmt.Printf("x[%d] type is string, content: %s\n", index, value)
		} else if value, ok := data.(Student); ok {
			fmt.Printf("x[%d] type is Student, content: %+v\n", index, value)
		}
	}

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] type is int, content: %d\n", index, value)
		case string:
			fmt.Printf("x[%d] type is string, content: %s\n", index, value)
		case Student:
			fmt.Printf("x[%d] type is Student, content: %+v\n", index, value)
		}
	}
}
