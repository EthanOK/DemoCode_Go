package main

import (
	"errors"
	"fmt"
)

/*
自定义错误
*/
func main() {
	err := test(0)
	if err != nil {
		fmt.Println(err)

		// 如果出现错误不想就继续执行
		panic("中断程序执行")
	}

	fmt.Println("Second")

}
func test(num int) error {
	if num == 0 {

		err := errors.New("除数不能为零啊")

		return err

	} else {
		n := 100 / num

		fmt.Println(n)

		return nil
	}

}
