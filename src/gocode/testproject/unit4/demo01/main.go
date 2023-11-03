package main

import (
	"fmt"
)

func main() {
	// 条件语句
	if 1 < 2 {
		fmt.Println("1 < 2")

	}

	if count := 10; count < 30 {

		fmt.Println("count < 30")
	}

	var cc = 40
	if cc < 30 {
		fmt.Println(cc, " < 30")
	} else {
		fmt.Println(cc, " > 30")

	}

}
