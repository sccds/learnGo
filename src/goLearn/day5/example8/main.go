package main

import (
	"encoding/json"
	"fmt"
)

// 结构体成员变量首字母必须大写
type IT struct {
	Company string 	`json:"company"`   // 改名字
	Subject []string `json:"subject"`
	IsOk bool 		`json:"-"`   // 此字段不会输出到 json
	Price float64  `json:",string"`  // 给数据加引号
}


func main()  {
	// 定义一个结构体变量，同时初始化
	s := IT{"51cto", []string{"go", "c++", "bigdata"}, true, 666.666 }

	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", "	") // 格式化编码
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("buf: ", string(buf) )
}
