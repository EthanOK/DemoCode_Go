package main

import (
	"fmt"
	"sync"
)

var amount int

var wg sync.WaitGroup

// 互斥锁
var mu sync.Mutex

func add(count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		mu.Lock()
		amount++
		mu.Unlock()
	}

}
func sub(count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		mu.Lock()
		amount--
		mu.Unlock()
	}

}

func main() {
	wg.Add(2)

	go add(100000)
	go sub(100000)

	wg.Wait()

	// 理论 结果为 0，但不是
	// 多协程操作同一数据需要加锁
	fmt.Println(amount)

}
