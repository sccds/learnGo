package main

import (
	"fmt"
	"regexp"
)

const text = `my email is 530979104@qq.com email is adbg@yahoo.com email is ccc@google.org.com`

func main() {
	//res := regexp.MustCompile("530979104@qq.com")
	//res := regexp.MustCompile("[A-Za-z0-9]+@[A-Za-z0-9]+\\.[A-Za-z0-9]+") // 如果用双引号，转义需要 \\
	//res := regexp.MustCompile(`[A-Za-z0-9]+@[A-Za-z0-9]+\.[A-Za-z0-9]+`) // 如果用反引号，转义\即可
	//match := res.FindString(text)
	//match := res.FindAllString(text, -1) // 找到所有

	res := regexp.MustCompile(`([A-Za-z0-9]+)@([A-Za-z0-9]+)\.([A-Za-z0-9.]+)`) // 提取值加括号
	match := res.FindAllStringSubmatch(text, -1)                                // 方法用 FindAllStringSubmatch 匹配成功进行截取，找到提取值

	fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}
