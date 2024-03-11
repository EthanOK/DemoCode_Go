package main

import "fmt"

func main() {

	// 声明方式 1
	var t1 Teacher

	fmt.Println(t1) // 结构体是值类型 { 0 }

	t1.Name = "张三"
	t1.Age = 18
	t1.School = "清华大学"
	fmt.Println(t1)

	// 声明方式 2
	var t2 Teacher = Teacher{}
	t2.Name = "李四"
	t2.Age = 20
	t2.School = "北京大学"
	fmt.Println(t2)

	// 声明方式 3
	// new 返回一个变量的指针
	var t3 *Teacher = new(Teacher)

	(*t3).Name = "王五"
	(*t3).Age = 22
	// 为符合程序员编程，go提供了简化的赋值方式
	t3.School = "浙江大学"

	fmt.Println(*t3)

	// 声明方式 4
	var t4 *Teacher = &Teacher{}
	t4.Name = "赵六"
	t4.Age = 24
	t4.School = "上海大学"
	fmt.Println(*t4)

}

type Teacher struct {
	Name   string
	Age    int
	School string
}
