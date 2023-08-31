package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/* 浮点数 与操作系统无关
	float32
	float64

	*/
	var PI0 float32 = 3.14
	fmt.Println(PI0)
	fmt.Printf("%T \n", PI0)
	fmt.Println(unsafe.Sizeof(PI0), "byte")

	var PI1 float32 = -3.14
	fmt.Println(PI1)
	fmt.Printf("%T \n", PI1)
	fmt.Println(unsafe.Sizeof(PI1), "byte")

	var PI2 float32 = 314e-2
	fmt.Println("314e-2 :", PI2)
	fmt.Printf("%T \n", PI2)
	fmt.Println(unsafe.Sizeof(PI2), "byte")

	var PI3 float32 = 314e+2
	fmt.Println("314e+2 :", PI3)
	fmt.Printf("%T \n", PI3)
	fmt.Println(unsafe.Sizeof(PI3), "byte")

	// 浮点数可能会有精度丢失，推荐使用 float64
	var num1 float32 = 3.141592608097
	var num2 float64 = 3.141592608097
	fmt.Println(num1) // 3.1415925
	fmt.Println(num2) // 3.141592608097

	// 默认 float64
	var num = 3.14
	fmt.Printf("%T \n", num)

}
