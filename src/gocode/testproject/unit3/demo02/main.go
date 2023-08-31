package main

import "fmt"

func main() {
	// 关系运算符 == != < > <= >=

	fmt.Println(4 == 5)
	fmt.Println(4 != 5)

	// 逻辑运算符 && || !

	fmt.Println(true && false)
	fmt.Println(true || false)

	// & * 指针

	var a int = 20
	fmt.Println("a的存储空间: ", &a)

	var p *int = &a
	fmt.Println("p指针指向的数值: ", *p)
}
