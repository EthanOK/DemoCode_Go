package main

import (
	"fmt"
	"time"
)

func main() {
	// 多协程 并发
	for i := 0; i < 5; i++ {
		// 使用匿名函数
		go func(n int) {
			fmt.Println("Hello from goroutine:", n)
		}(i)
	}

	time.Sleep(time.Second) // 等待goroutine执行完毕

}
