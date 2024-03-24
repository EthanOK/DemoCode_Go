package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 创建客户端
	fmt.Println("客户端启动...")
	// 1. 链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接服务端失败:", err)
		return
	}
	fmt.Println("连接服务端成功")

	defer conn.Close()

	// 1. os.Stdin 终端标准输入
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取用户输入失败:", err)
		return
	}
	// 2. 发送数据

	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("发送数据失败:", err)
		return
	}
	fmt.Printf("客户端发送了%d个字节的数据\n", n)
}
