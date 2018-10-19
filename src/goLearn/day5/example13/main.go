package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func WriteFile(path string)  {
	// 打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	// 使用延迟关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		// 格式化字符串
		buf = fmt.Sprintf("i = %d\n", i)
		fmt.Println("buf =", buf)
		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err =", err)
		}
		fmt.Println("n =", n)
	}
}

// 读文件
func ReadFile(path string)  {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	// 延时关闭
	defer f.Close()

	buf := make([]byte, 1024 * 2) // 每次读取2k大小
	// n 代表从文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1 =", err1)
		return
	}
	fmt.Println("buf =", string(buf[:n]))
}


// 按照行读取文件
func ReadFileLine(path string)  {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	// 延时关闭
	defer f.Close()

	// 把文件内容放在缓冲区
	r := bufio.NewReader(f)

	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { // 文件已经结束
				break
			}
			fmt.Println("err =", err)
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}
}


func main()  {
	path := "./demo.txt"
	//WriteFile(path)
	//ReadFile(path)
	ReadFileLine(path)
}
