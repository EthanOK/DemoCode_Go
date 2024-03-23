package main

import (
	"fmt"
	"time"
)

// 管道的阻塞

func main() {
	// 创建一个管道
	ch := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 5)
		ch <- 12

	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch <- 10

	}()

	// 阻塞等待管道中的数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
