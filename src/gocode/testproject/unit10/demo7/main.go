package main

import (
	"fmt"

	"gocode.ethan/democode-go/src/gocode/testproject/unit10/demo7/model"
)

func main() {

	person := model.NewPerson("zhangsan", 18)
	fmt.Println(person.GetAge())
	person.SetAge(160)

}
