package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  30,
	}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	fmt.Println(t) // 输出: main.Person

	fmt.Println(v) // 输出: {Alice 30}

	if pp, ok := v.Interface().(Person); ok {
		fmt.Println(pp.Name) // 输出: {Alice 30}
	}

}
