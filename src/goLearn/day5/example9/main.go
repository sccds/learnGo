package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	// 创建一个 map
	m := make(map[string]interface{}, 4)
	m["company"] = "51cto"
	m["subject"] = []string{"go", "c++", "python", "Test"}
	m["isok"] = true
	m["price"] = 666.666

	// 编码成 json
	result, err := json.MarshalIndent(m, "", "		")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}