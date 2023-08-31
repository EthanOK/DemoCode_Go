package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// cannot use 130 (untyped int constant) as int8 value in variable declaration (overflows)
	var a int8 = 120
	fmt.Println(a)
	// 默认 int 类型 64位计算机8个字节
	var b = 259
	fmt.Printf("%T", b)
	fmt.Println("")
	fmt.Println(unsafe.Sizeof(b), "字节")

	/* 输出
	120
	int
	8 字节
	*/

}
