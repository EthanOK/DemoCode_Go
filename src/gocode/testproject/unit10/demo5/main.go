package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// 如果一个类型实现了 String() 这个方法，
// fmt.Println() 会调用这个变量的的String()进行输出
func (s Student) String() string {
	str := fmt.Sprintf("Student: { Name:%s, Age:%d }", s.Name, s.Age)

	return str

}

func main() {

	strudent := Student{Name: "张三", Age: 18}

	// Student: { Name:张三, Age:18 }
	fmt.Println(strudent)

}
