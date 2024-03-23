package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	ch := make(chan int, 10)

	go writeChan(ch)
	go readChan(ch)

	wg.Wait()

}

func writeChan(ch chan int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i + 1 // 向通道发送数据
		fmt.Println("write", i+1)
		// time.Sleep(time.Second)
	}
	close(ch) // 关闭通道
}

func readChan(ch chan int) {
	defer wg.Done()

	for v := range ch {
		fmt.Println("read", v)
		time.Sleep(time.Second)
	}

}
