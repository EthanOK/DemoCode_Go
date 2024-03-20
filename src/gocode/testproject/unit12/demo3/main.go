package main

import (
	"fmt"
	"sync"
)

/*
解决 主死从随 的问题
等协程执行完
*/
func main() {

	fmt.Println("main start")

	// 定义 WaitGroup
	var waitGroup sync.WaitGroup

	for i := 0; i < 5; i++ {

		waitGroup.Add(1) // 协程开始时 加 1

		go func(n int) {
			defer waitGroup.Done() // 协程结束时 减 1

			fmt.Println("Hello from goroutine:", n)
		}(i)
	}

	// 主线程阻塞，等 waitGroup 减为 0

	waitGroup.Wait()

	fmt.Println("main end")

}
