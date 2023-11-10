package main

import "fmt"

func main() {
	// for
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)

}
