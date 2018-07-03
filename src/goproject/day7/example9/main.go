// http client 请求百度

package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	res, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err:", err)
		return
	}

	fmt.Println(string(data))
}
