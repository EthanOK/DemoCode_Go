package main

import "fmt"

func main() {
	// switch ... case
	var score int = 98

	switch score / 10 {

	case 10, 9:
		fmt.Println("A")

	case 8:
		fmt.Println("B")

	case 7:
		fmt.Println("C")

	case 6:
		fmt.Println("D")

	default:
		fmt.Println("不及格")
	}

}
