// http 编程

package main

import (
	"net/http"
	"fmt"
)

func Hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle hello")
	fmt.Fprint(w, "hello")
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle login")
	fmt.Fprint(w, "login")
}

func history(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle history")
	fmt.Fprint(w, "history")
}

func main()  {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/user/history", history)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}

}
