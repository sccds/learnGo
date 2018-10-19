// 把数据渲染到html模板，html和业务绑定

package main

import (
	"html/template"
	"fmt"
	"net/http"
	"io"
)

var myTemplate *template.Template

type Result struct {
	output string
}

// 实现 Write 接口
func (p *Result) Write(b []byte) (n int, err error)  {
	fmt.Println("called by template")
	p.output += string(b)
	return len(b), nil
}

type Person struct {
	Name string
	Title string
	Age int
}

func userInfo(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle hello")
	var arr []Person
	p1 := Person{Name:"Mary001", Age:10, Title:"my web"}
	p2 := Person{Name:"Mary002", Age:10, Title:"my web"}
	p3 := Person{Name:"Mary003", Age:10, Title:"my web"}
	arr = append(arr, p1)
	arr = append(arr, p2)
	arr = append(arr, p3)

	resultWriter := &Result{}
	io.WriteString(resultWriter, "hello world")
	err := myTemplate.Execute(w, arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("template render data:", resultWriter.output)
}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main()  {
	initTemplate("/Users/xliu/Documents/51CTO_Golang/learnGo/src/goproject/day7/example13/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}