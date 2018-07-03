package main

import (
	"net"
	"fmt"
	"strings"
)

func HandleConn(conn net.Conn)  {
	// 函数调动完毕，自动关闭conn
	defer conn.Close()

	// 获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect successful")

	buf := make([]byte, 2048)

	for {
		// 读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err =", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, "exit")
			return
		}

		// 数据转化成小写发送给对应的用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}


func main()  {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err =", err)
			return
		}

		// 处理用户发送的数据
		go HandleConn(conn)
	}
}
