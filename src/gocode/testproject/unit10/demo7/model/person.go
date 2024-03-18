package model

import "fmt"

// 结构体封装 只能在包内直接访问
type person struct {
	name string
	age  int
}

// 定义工厂模式的函数，相当于构造器
func NewPerson(name string, age int) *person {
	return &person{name, age}
}

// 定义 set get 方法，对字段 进行封装

func (p *person) SetName(name string) {
	(*p).name = name
}

func (p *person) GetName() string {
	return (*p).name
}

func (p *person) SetAge(age int) {
	if age < 0 || age > 150 {
		fmt.Println("年龄不合法")
		return
	}

	(*p).age = age
}
func (p *person) GetAge() int {
	return (*p).age
}
