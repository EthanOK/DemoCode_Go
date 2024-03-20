package main

import (
	"fmt"
	"sort"
)

// 接口的定义
type SayHello interface {
	sayHello()
}

// 只要实现了接口的所有方法，就算实现了该接口
type Chinese struct {
	name string
}

func (c Chinese) sayHello() {
	fmt.Println("你好", c.name)
}

type English struct {
	name string
}

func (e English) sayHello() {
	fmt.Println("Hi", e.name)
}
func say(s SayHello) {
	s.sayHello()

}

type Person struct {
	Name string
	Age  int
}

// 定义了一个 ByAge 类型，它是一个包含了 Person 结构体的切片
type ByAge []Person

// ByAge 实现了 sort.Interface 接口
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	var c SayHello = Chinese{"张三"}
	var e SayHello = English{"John"}
	say(c)
	say(e)

	persons := []Person{{"张三", 28}, {"李四", 20}, {"王五", 22}}

	sort.Sort(ByAge(persons))

	fmt.Println(persons)

}
