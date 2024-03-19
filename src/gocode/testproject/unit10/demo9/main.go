package main

// 接口的定义
type SayHello interface {
	sayHello()
}

// 只要实现了接口的所有方法，就算实现了该接口
type Chinese struct {
	name string
}

func (c Chinese) sayHello() {
	println("你好", c.name)
}

type English struct {
	name string
}

func (e English) sayHello() {
	println("Hi", e.name)
}
func say(s SayHello) {
	s.sayHello()

}

func main() {
	var c SayHello = Chinese{"张三"}
	var e SayHello = English{"John"}
	say(c)
	say(e)

}
