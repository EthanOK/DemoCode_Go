package main

import "fmt"

/*
错误处理机制

defer + recover
*/
func main() {
	test(0)

	fmt.Println("Second")

}
func test(num int) {
	// defer + recover 捕获错误不中断程序，defer后加上匿名函数的调用
	defer func() {
		// 利用 recover 内置函数，捕获错误
		err := recover()
		// 如果没有捕获错误，返回值为 nil
		if err != nil {
			fmt.Println(err)
		}
	}()

	n := 5 / num

	fmt.Println(n)
}
