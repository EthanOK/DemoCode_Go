package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件中的环境变量
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 使用 os 包中的 Getenv 函数读取环境变量
	ALCHEMY_PRC_HTTP := os.Getenv("ALCHEMY_PRC_HTTP")

	// 打印读取的环境变量
	// log.Println("API Key:", ALCHEMY_PRC_HTTP)

	println(ALCHEMY_PRC_HTTP)

}
