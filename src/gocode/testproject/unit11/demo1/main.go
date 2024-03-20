package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
文件读取操作
*/

func readFile1() {
	// open file
	file, err := os.Open("src/gocode/testproject/unit11/demo1/test.txt")
	if err != nil {
		fmt.Println("open file error!")
	}

	// close file
	defer file.Close()

	// read file
	// 创建一个流
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')

		fmt.Print(str)

		if err == io.EOF {

			break
		}

	}
	fmt.Println()
}

func readFile2() {
	// os.ReadFile 方法内包含 open close文件操作
	context, err := os.ReadFile("src/gocode/testproject/unit11/demo1/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(context))
}

func main() {

	readFile1()

	readFile2()

}
