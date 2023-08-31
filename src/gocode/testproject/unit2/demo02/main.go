package main

import "fmt"

// 全局变量
var (
	name = "A"
	age  = 18
)

func main() {
	// 局部变量
	var num0 int = 19

	var num1 int

	var num2 = 20

	num3 := 40
	fmt.Println(num0, num1, num2, num3)
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

}
