package main

import "fmt"

func main() {

	// 指针

	a := 1
	b := 5
	convert(&a, &b)

	fmt.Println(a)
	fmt.Println(b)

}

func convert(aa *int, bb *int) {
	temp := *aa
	*aa = *bb
	*bb = temp

}
