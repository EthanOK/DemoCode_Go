package main

import "fmt"

type Animal struct {
	Age    int
	Weight float32
}

func (a *Animal) ShowInfo() {
	fmt.Println("Age:", a.Age, "Weight:", a.Weight)
}

// Cat 继承了 Animal 结构体，匿名结构体，并且添加了自己的 Name 字段。
type Cat struct {
	Animal
	Name string
}

// 组合关系，不是继承
type AAA struct {
	AA   Animal
	Name string
}

func (c *Cat) ShowName() {
	fmt.Println("I'm a cat.", "My name is", c.Name)
}

func main() {
	c := Cat{
		Animal: Animal{
			Age:    2,
			Weight: 3.5,
		},
		Name: "Tom",
	}

	c.ShowInfo()

	c.ShowName()

}
