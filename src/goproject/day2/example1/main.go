package main

import "fmt"

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int  {
	return a - b
}

func operator(op add_func, a int, b int) int {
	return op(a, b)
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func eval(a, b int, op string)(int, error)  {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		p, _ := div(a, b)
		return p, nil
	default:
		return 0, fmt.Errorf("unsupport operation " + op)

	}
}

func main()  {
	//c := add
	//c := sub
	sum := operator(func(a int, b int) int { return a * b }, 5, 6)
	fmt.Println(sum)

	p, q := div(9, 4)
	fmt.Println(p, q)

	fmt.Println(eval(4, 5, "+"))
}