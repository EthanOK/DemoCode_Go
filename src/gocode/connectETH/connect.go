package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 使用 os 包中的 Getenv 函数读取环境变量
	ALCHEMY_PRC_HTTP := os.Getenv("ALCHEMY_PRC_HTTP")
	client, err := ethclient.Dial(ALCHEMY_PRC_HTTP)
	if err != nil {
		log.Fatal(err)
	}

	// Get the balance of an account
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account balance:", balance) // 25893180161173005034

	// Get the latest known block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockTime := block.Time()
	fmt.Println("Latest block:", block.Number().Uint64())
	fmt.Println("Latest block Time:", blockTime)
	// 打印 本地时间戳
	// 获取当前时间
	currentTime := time.Now()
	// 获取时间戳（秒级）
	timestamp := currentTime.Unix()
	fmt.Println("timestamp:", timestamp)
	// 间隔时间
	fmt.Println("Interval time:", uint64(timestamp)-uint64(blockTime))
}
