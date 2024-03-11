package main

import "fmt"

type Integer int

func (i Integer) print() {
	fmt.Println(i)
}

func (i *Integer) change(i_ int) {
	*i = Integer(i_)

}

func main() {

	var i Integer = 43
	i.print()

	i.change(3)
	i.print()

}
