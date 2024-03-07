package main

import "fmt"

/**
函数也是一种数据类型, 可作为参数

自定义数据类型(起别名，myint与int 不是同一种数据类型): type myint int

支持对函数返回值命名
*/

func main() {

	t := test
	println("%T", t)
	t(1)
	test2(2, t)
	test3(3, t)

	fmt.Println(sub_sum(10, 3))
	fmt.Println(sum_sub(10, 3))
}

func test(num int) {
	fmt.Println(num)
}

func test2(num int, tt func(int)) {
	tt(num)

}

type myFunc func(int)

func test3(num int, tt myFunc) {
	tt(num)

}

// 传统写法
func sub_sum(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

// 升级写法
func sum_sub(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}
