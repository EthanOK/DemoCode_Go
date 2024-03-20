package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁 场景：读多写少
var lock sync.RWMutex

var wg sync.WaitGroup

func read() {
	defer wg.Done()
	lock.RLock()

	fmt.Println("读数据")
	time.Sleep(time.Second)
	fmt.Println("读取成功")

	lock.RUnlock()

}

func write() {
	defer wg.Done()
	lock.Lock()

	fmt.Println("写数据")
	time.Sleep(time.Second)
	fmt.Println("写入成功")

	lock.Unlock()

}
func main() {
	wg.Add(6)
	// 读多写少
	for i := 0; i < 5; i++ {
		go read()
	}

	go write() // 在写的过程中，锁生效；但读可以并发读

	wg.Wait()
}
