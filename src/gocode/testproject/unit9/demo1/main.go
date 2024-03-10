package main

import "fmt"

func main() {
	define_map()

	operate_map()
}

/*
map 特点：
1. map使用前一定要make，第二个参数可省略，默认分配一个内存
2. map的key-value是无序的
3. key不能重复，如遇重复，后值替换前值
*/
func define_map() {
	// map 声明 方式 1:
	var m map[int]string
	// 只声明 map， 没有分配内存空间
	// 必须通过make 初始化
	m = make(map[int]string, 5)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"

	for k, v := range m {
		fmt.Println(k, v)
	}

	// map 声明 方式 2:
	mm := make(map[int]string)
	mm[1] = "aa"
	mm[2] = "bb"
	mm[3] = "cc"
	fmt.Println(mm)

	// map 声明 方式 3:
	mmm := map[int]string{
		22: "aaa",
		33: "bbb",
		44: "ccc",
	}
	fmt.Println(mmm)

}

/*
增删改查
*/
func operate_map() {
	demo := make(map[int]string)
	// 增
	demo[1] = "a"
	demo[2] = "b"

	// 删
	delete(demo, 1)

	// 改
	demo[2] = "c"

	// 查
	value, ok := demo[2]
	if ok {
		fmt.Println(value)
	}

	// 清空 map
	demo = make(map[int]string)
	fmt.Println(demo)

}
