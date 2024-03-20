package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件写入操作
func main() {
	file, err := os.OpenFile("src/gocode/testproject/unit11/demo2/test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 写入字符串
	writer := bufio.NewWriter(file)

	writer.WriteString("xixixixi")

	// 刷新数据
	writer.Flush()

}
