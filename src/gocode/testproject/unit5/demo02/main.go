package main

import "fmt"

func main() {
	c := sum(2, 3)

	fmt.Println(c)

}

func sum(a ...int) int {
	var ss int
	for _, v := range a {
		ss += v
	}

	return ss

}
