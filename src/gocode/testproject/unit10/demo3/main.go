package main

import "fmt"

// 结构体的方法

type Person struct {
	Name string
	Age  int
}

// 为 Person 结构体绑定方法
// 结构体类型是值传递，在方法调用中，是值拷贝
func (p Person) getName() {
	p.Name = "Bob" // 修改副本，不会影响原始结构体
	fmt.Println(p.Name)
}

// 为指针类型绑定方法
func (p *Person) setName(s string) {
	(*p).Name = s
	// p.Name = s // 效果同上
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	p.getName()

	(&p).setName("Bob2")
	// p.setName("Bob2") // 效果同上

	fmt.Println(p.Name) // Bob2
}
