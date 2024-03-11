package main

import "fmt"

type Student struct {
	Name string
}

type Person struct {
	Name string
}

func main() {

	s := Student{Name: "zhangsan"}
	fmt.Println(s)
	p := Person{Name: "lisi"}

	s = Student(p) // 类型不同，但是结构体只有相同的字段才可以赋值
	fmt.Println(s)
}
