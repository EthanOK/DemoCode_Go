package main

import "fmt"

func main() {

	// // 管道的定义
	// var intChan chan int
	// make 初始化
	intChan := make(chan int, 3)
	// 是引用类型
	fmt.Printf("%v\n", intChan)

	// 向管道存放数据,但不能存放大于容量的数据
	intChan <- 10
	intChan <- 23
	intChan <- 47

	fmt.Println(len(intChan), cap(intChan))

	// 读取管道的数据(队列，先进先出)
	num1 := <-intChan
	num2 := <-intChan
	fmt.Println(num1)
	fmt.Println(num2)

}
