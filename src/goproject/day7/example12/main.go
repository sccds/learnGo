// 把数据渲染到html模板

package main

import (
	"html/template"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Title string
	age string
}

func main()  {
	t, err := template.ParseFiles("/Users/xliu/Documents/51CTO_Golang/learnGo/src/goproject/day7/example12/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", age: "31", Title: "my website"}
	// Execute实现 write 接口, 本例是输出到stdout终端
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}