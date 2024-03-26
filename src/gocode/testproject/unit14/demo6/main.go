package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改变量
type Student struct {
	Name string
	Age  int
}

func testStructReflect(i interface{}) {

	value := reflect.ValueOf(i)
	n := value.Elem().NumField()
	fmt.Println(n)
	// 修改字段值
	value.Elem().Field(0).SetString("李四")

}

func main() {

	stu := Student{Name: "张三", Age: 18}

	testStructReflect(&stu)

	fmt.Println(stu)

}
