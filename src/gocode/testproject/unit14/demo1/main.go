package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {

	// 获取一个变量的数据类型
	t := reflect.TypeOf(i)
	fmt.Println(t)

	// 获取一个变量的值
	v := reflect.ValueOf(i)
	fmt.Println(v)

	//断言 断言是一个强制类型转换，可能会失败
	ii := v.Interface()
	if num, ok := ii.(integer); ok {
		fmt.Println("num = ", num)
	}

}

type integer int

func main() {

	var num integer = 100

	testReflect(num)

}
