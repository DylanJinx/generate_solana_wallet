package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gagliardetto/solana-go"
)

var generatedCount = 0 // 生成的钱包数量，在多少个钱包里，找到了指定的solana钱包
var numThreads = 16 // 线程数量
var startTime = time.Now() // 开始时间
var searchPreTherm = "9999" // 指定的solana钱包前缀
var shouldStopPreThreads = false // 是否停止线程
var searchPostTherm = "9999" // 指定的钱包后缀
var shouldStopPostThreads = false // 是否停止线程

func generateWallet_searchPreTherm() {
	for {
		if shouldStopPreThreads {
			return
		}

		newWallet := solana.NewWallet()

		// HasPrefix：检查newWallet.PublicKey()是否以指定的前缀searchPreTherm开头
		if strings.HasPrefix(newWallet.PublicKey().String(), searchPreTherm) && !shouldStopPreThreads {
			/*
			strings.Split(newWallet.PublicKey().String(), searchPreTherm)：这将把公钥字符串按照 searchPreTherm（例如 "1234") 进行分割。
			例如，假设 newWallet.PublicKey().String() 返回的是 "abcd1234efgh5678"，并且 searchPreTherm 是 "1234"，那么 strings.Split("abcd1234efgh5678", "1234") 会返回 ["abcd", "efgh5678"]。

			获取第二部分：
			[1] 获取第二部分（索引从0开始），即 efgh5678。
			取第二部分的第一个字符：
			[:1] 是字符串切片操作，它取的是第二部分的前1个字符。
			例如，从 efgh5678 中取出第一个字符是 "e"。
			*/
			firstCharAfterSearchTherm := strings.Split(newWallet.PublicKey().String(), searchPreTherm)[1][:1]
			if firstCharAfterSearchTherm == strings.ToUpper(firstCharAfterSearchTherm) {
				fmt.Printf("指定的钱包公钥: %s\n", newWallet.PublicKey())
				fmt.Printf("指定的钱包私钥: %v\n", newWallet.PrivateKey)
				fmt.Printf("尝试了： %d；花费了： %s\n", generatedCount+1, time.Since(startTime))
				shouldStopPreThreads = true
			}
		}
		generatedCount++
		if generatedCount%1000000 == 0 {
			fmt.Printf("尝试了： %d；花费了： %s\n", generatedCount, time.Since(startTime))
		}
	}
}

func generateWallet_searchPostTherm() {
	for {
		if shouldStopPostThreads {
			return
		}

		newWallet := solana.NewWallet()

		// HasSuffix：检查newWallet.PublicKey()是否以指定的后缀searchPostTherm结尾
		if strings.HasSuffix(newWallet.PublicKey().String(), searchPostTherm) && !shouldStopPostThreads {
			fmt.Printf("指定的钱包公钥: %s\n", newWallet.PublicKey())
			fmt.Printf("指定的钱包私钥: %v\n", newWallet.PrivateKey)
			fmt.Printf("尝试了： %d；花费了： %s\n", generatedCount+1, time.Since(startTime))
			shouldStopPostThreads = true
		}
		generatedCount++
		if generatedCount%1000000 == 0 {
			fmt.Printf("尝试了： %d；花费了： %s\n", generatedCount, time.Since(startTime))
		}
	}
}

func main() {
	fmt.Printf("指定的钱包前缀: %s\n", searchPreTherm)
	// 启动了 numThreads 个并发任务
	for i := 0; i < numThreads; i++ {
		go generateWallet_searchPreTherm()
	}

	fmt.Printf("指定的钱包后缀: %s\n", searchPostTherm)
	// 启动了 numThreads 个并发任务
	for i := 0; i < numThreads; i++ {
		go generateWallet_searchPostTherm()
	}

	fmt.Scanln()
}