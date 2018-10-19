package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:"-"`
	Subject []string `json:"subject"`
	IsOk bool `json:",string"`
	Price float64 `json:",string"`
}

type IT2 struct {
	Subject []string `json:"subject"`
}



func main()  {
	jsonBuf := `
	{
		"company": "51cto",
		"isok": "true",
		"price": "666.666",
		"subject": [
				"go",
				"c++",
				"python",
				"Test"
		]
	}`

	var tmp IT // 定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Printf("tmp = %+v\n", tmp)

	var temp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &temp2)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", temp2)
}
