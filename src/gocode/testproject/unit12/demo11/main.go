package main

import (
	"fmt"
	"time"
)

func printData() {
	fmt.Println("Hello World")
}

// recover 用于捕获panic异常
func div0() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}

	}()
	mun := 0

	r := 2 / mun

	fmt.Println(r)

}

func main() {

	go printData()
	go div0()

	time.Sleep(2 * time.Second) // 等待goroutine执行完毕

}
