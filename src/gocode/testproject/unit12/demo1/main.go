package main

import (
	"fmt"
	"time"
)

// 主死从随
func main() { //主线程

	go printNumber() //创建一个goroutine去执行printNumber函数

	for i := 0; i < 5; i++ {
		fmt.Println("main:", i)

		time.Sleep(time.Second)

	}
}

func printNumber() {
	for i := 0; i < 10; i++ {
		fmt.Println("goroutine:", i)

		time.Sleep(time.Second)

	}
}
