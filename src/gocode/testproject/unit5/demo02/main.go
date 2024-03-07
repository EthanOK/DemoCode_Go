package main

import "fmt"

func main() {
	c := sum(2, 3)

	fmt.Println(c)

	d := sum(2, 3, 4, 5, 6)

	fmt.Println(d)

}

// 可变参数
func sum(a ...int) int {
	// 可变参数 作为 切片
	var ss int
	for _, v := range a {
		ss += v
	}

	return ss

}
