package main

import "fmt"

/*
数组声明

var arr [10]int
or
balance := [5]int{1000, 2, 3, 17, 50}
or
balance := [...]int{1000, 2, 3, 17, 50}
*/
func main() {

	// 数组
	var scores [5]int

	for k, _ := range scores {

		fmt.Printf("请输入第 %d 个人的成绩：", k+1)
		fmt.Scanln(&scores[k])
	}
	// 切片
	scores_slice := scores[:]

	scores_sum := sum(scores_slice)

	scores_avg := scores_sum / len(scores)

	fmt.Println(scores_sum, scores_avg)

	var arr [10]int

	// arr 中存储的是地址值
	fmt.Printf("%p \n", &arr)
	// 第一个空间的地址
	fmt.Printf("%p \n", &arr[0])
	// 第二个空间的地址
	fmt.Printf("%p \n", &arr[1])
	// 最后一个空间的地址
	fmt.Printf("%p \n", &arr[9])

	test2D_Array()

}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func test2D_Array() {
	var array_2d [2][3]int
	array_2d[0] = [...]int{1, 2, 3}
	array_2d[1] = [...]int{4, 5, 6}

	// [[1 2 3] [4 5 6]]
	fmt.Println(array_2d)

	for k, v := range array_2d {
		for k2, v2 := range v {
			fmt.Printf("arr[%d][%d] = %d \t", k, k2, v2)
		}
		fmt.Println()

	}
}
