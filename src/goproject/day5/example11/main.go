package main

import (
	"encoding/json"
	"fmt"
)

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

	// 创建 map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &m)
	if  err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("m = %+v\n", m)

	var str string
	for key, value := range m {
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的类型为string, value = %s\n", key, data)
		case bool:
			fmt.Printf("map[%s]的类型为bool, value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的类型为float64, value = %f\n", key, data)
		case []string:
			fmt.Printf("map[%s]的类型为[]string, value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的类型为[]interface, value = %v\n", key, data)
		}

	}
	fmt.Println(str)
}
