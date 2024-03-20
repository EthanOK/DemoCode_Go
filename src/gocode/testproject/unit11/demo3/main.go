package main

import (
	"fmt"
	"os"
)

// 文件复制
func main() {
	// ...
	source_path := "src/gocode/testproject/unit11/demo2/test.txt"
	target_path := "src/gocode/testproject/unit11/demo3/test.txt"
	copy(source_path, target_path)

}

func copy(s string, t string) {

	content, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(t, content, 0777)
	if err != nil {
		fmt.Println(err)
	}

}
