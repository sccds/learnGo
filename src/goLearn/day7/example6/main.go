// 实现文件传输 client / 文件发送端

package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

func SendFile(path string, conn net.Conn)  {
	// 以只读形式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.open err =", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	// 读取文件内容，读多少发送多少
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err =", err)
			}
			return
		}

		// 发送数据
		conn.Write(buf[:n]) // 给服务器发送内容
	}
}


func main()  {
	// 提示输入文件
	fmt.Println("请输入要传输的文件：")
	var path string
	fmt.Scan(&path)

	// 获取文件名 info.Name()
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.stat err =", err)
		return
	}

	// 主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err1 =", err1)
		return
	}
	defer conn.Close()

	// 给接收方先发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write err =", err)
		return
	}

	// 接收接收方的回复，如果回复OK，说明对方准备好，可以发送文件
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err =", err)
		return
	}
	if "ok" == string(buf[:n]) {
		// 开始发送文件内容
		SendFile(path, conn)
	}
}
