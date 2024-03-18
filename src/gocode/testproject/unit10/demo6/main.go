package main

import "fmt"

// 函数和方法的区别

type Student struct {
	Name string
	Age  int
}

func (s *Student) SayHello() {
	fmt.Printf("Hello, my name is %s\n", s.Name)
}

func SayHello(s *Student) {
	fmt.Printf("Hello, my name is %s\n", s.Name)
}

func main() {
	s := Student{Name: "Alice", Age: 18}

	// 对于方法， 接受者为值类型，可传入指针类型；接受者为指针类型，可传入值类型
	s.SayHello() // 调用方法
	(&s).SayHello()

	// 对于函数来说，参数类型对应是什么就要传入什么
	SayHello(&s) // 调用函数
	// SayHello(s) error
}
