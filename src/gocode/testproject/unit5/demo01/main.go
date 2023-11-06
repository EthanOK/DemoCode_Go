package main

import "fmt"

func main() {
	c := ADD(2, 3)
	fmt.Println(c)
}

// 大小写字母

// ADD public fuanction
func ADD(a int, b int) int {
	return a + b
}

// add private fuanction
func add(a int, b int) int {
	return a + b
}
