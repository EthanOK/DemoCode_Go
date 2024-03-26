package main

import (
	"fmt"
	"reflect"
)

// 反射 获取变量 类别

type MyStruct struct {
	Name string
}

func main() {

	reType := reflect.TypeOf(MyStruct{})
	// 变量的类型
	fmt.Println(reType.Name()) // MyStruct
	// 变量的类别
	fmt.Println(reType.Kind()) // struct

	reValue := reflect.ValueOf(MyStruct{})

	fmt.Println(reValue.Kind()) // struct

}
