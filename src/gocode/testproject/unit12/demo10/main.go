package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int, 1)

	stringChan := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		intChan <- 111
	}()

	go func() {
		time.Sleep(time.Second * 1)
		stringChan <- "hello"
	}()

	// 哪条路径先执行完毕，就执行哪条路径
	select {
	case value := <-intChan:
		fmt.Println(value)

	case value := <-stringChan:
		fmt.Println(value)
	}

}
