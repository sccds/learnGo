// 文件传输接收方

package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func RecvFile(fileName string, conn net.Conn)  {
	// 新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err =", err)
		return
	}
	buf := make([]byte, 1024*4)
	// 接收多少，写多少
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err =", err)
			}
			return
		}
		f.Write(buf[:n])
	}


}

func main()  {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net listener err =", err)
		return
	}
	defer listener.Close()

	// 阻塞用户连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err1 =", err1)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err1 =", err)
		return
	}

	fileName := string(buf[:n])

	// 回复“ok”
	conn.Write([]byte("ok"))

	// 调用保存文件方法
	RecvFile(fileName, conn)
}
