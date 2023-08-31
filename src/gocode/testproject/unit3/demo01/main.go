package main

import "fmt"

func main() {
	// +-*/% ++ --
	// go语言 只有 i++ i-- （与Java不同）只表示自增
	var i = 10
	i++
	fmt.Println(i)
	// = += -=
	var a = 8
	var b = 10
	// 交换两个数的值
	a, b = b, a
	fmt.Println(a, b)
}
