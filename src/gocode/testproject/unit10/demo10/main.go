package main

import "fmt"

type Say interface {
	say()
}

type Chinese struct {
}

func (c Chinese) say() {
	fmt.Println("吃了吗？")

}

func (c Chinese) eat() {
	fmt.Println("吃馒头")
}

type American struct {
}

func (a American) say() {
	fmt.Println("Hi!")
}

func (a American) disco() {
	fmt.Println("Disco!")
}

func saysay(s Say) {
	s.say()
	// 断言
	// s.(Chinese): 看 s 是否为 Chinese
	if c, ok := s.(Chinese); ok {
		c.eat()
	}

}

func saysaySwitch(s Say) {
	s.say()
	// s.(type) 判断是哪种结构体类型
	switch s.(type) {
	case Chinese:
		s.(Chinese).eat()
	case American:
		s.(American).disco()
	}
}

func main() {
	c := Chinese{}
	a := American{}
	saysay(c)
	saysaySwitch(a)

}
