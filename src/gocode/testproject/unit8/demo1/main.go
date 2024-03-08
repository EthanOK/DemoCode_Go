package main

import "fmt"

func main() {

	define_slice()

	append_slice()

	copy_slice()

}

func define_slice() {
	// 1. array
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := arr[2:5]
	slice1[0] = 10
	fmt.Println(slice1)
	//
	fmt.Println(arr)

	// 2. make
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)

	// 3. []int
	slice3 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice3)
}

func append_slice() {
	slice0 := []int{1, 2, 3, 4, 5}
	slice1 := append(slice0, 6)
	fmt.Println(slice0) //[1 2 3 4 5]
	fmt.Println(slice1) // [1 2 3 4 5 6]

	// append 底层创建了新数组
	slice0 = append(slice0, 7, 8)
	fmt.Println(slice0) // [1 2 3 4 5 7 8]
}

func copy_slice() {
	slice0 := []int{1, 2, 3, 4, 5}
	slice1 := make([]int, 6)

	// copy(dst, src) : copy 内置函数将源切片中的元素复制到目标切片中。
	copy(slice1, slice0)
	fmt.Println(slice1) // [1 2 3 4 5 0]
}
