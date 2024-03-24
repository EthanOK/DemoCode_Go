package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建服务器端
	fmt.Println("服务器启动...")

	// 进行监听
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("服务器端启动失败:", err)
		return
	}
	fmt.Println("服务器监听中...")

	//  循环等待客户端连接

	for {
		conn, err := listener.Accept()

		if err != nil {

			fmt.Println("等待客户端连接失败:", err)
			return
		}

		// 启动一个goroutine处理客户端请求
		go handlerConn(conn)

	}

}

func handlerConn(conn net.Conn) {
	defer conn.Close()
	for {

		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			return
		}

		fmt.Println("收到客户端", conn.RemoteAddr(), "信息:", string(buffer[:n]))
	}

}
