package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

/**
 * 连接到区块链
 * @param url 区块链的链接地址，默认使用测试地址
 * @param blockNumber_ 区块链的区块号，默认使用测试区块
 * @return *ethclient.Client 区块链的链接
 */
func ClientEthereum(url string) *ethclient.Client {

	// 没有传区块链的链接地址，默认使用测试地址
	if url == "" {
		url = "https://eth-sepolia.g.alchemy.com/v2/U-ZXtAtro-pIJtwdrhq-NNNq6sKGsK32"
	}

	// 链接区块链网络
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	//
	//// 默认使用一个测试的区块
	//var blockNumber *big.Int
	//if blockNumber_ == 0 {
	//	blockNumber = big.NewInt(5671744)
	//} else {
	//	blockNumber = big.NewInt(blockNumber_)
	//}
	//
	//// 获取区块的 header
	//header, err := client.HeaderByNumber(context.Background(), blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("------------------------------区块 header 信息：------------------------------")
	//fmt.Println(header.Number.Uint64())     // 5671744
	//fmt.Println(header.Time)                // 1712798400
	//fmt.Println(header.Difficulty.Uint64()) // 0
	//fmt.Println(header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	//
	//// 获取区块的 body
	//block, err := client.BlockByNumber(context.Background(), blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("------------------------------区块 body 信息：------------------------------")
	//fmt.Println(block.Number().Uint64())     // 5671744
	//fmt.Println(block.Time())                // 1712798400
	//fmt.Println(block.Difficulty().Uint64()) // 0
	//fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	//fmt.Println(len(block.Transactions()))   // 70
	//
	//// 获取交易数量
	//count, err := client.TransactionCount(context.Background(), block.Hash())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(count) // 70
	return client
}
