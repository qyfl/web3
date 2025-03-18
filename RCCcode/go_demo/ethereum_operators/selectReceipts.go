package ethereum_operators

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	client2 "go_demo/client"
	"log"
	"math/big"
)

/**
 * 通过区块哈希查询收据
 */
func SelectReceiptsByBlockHash(blockHash_ string) []*types.Receipt {
	client := client2.ClientEthereum("")

	// 没有传区块哈希使用默认值
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")

	if blockHash_ != "" {
		blockHash = common.HexToHash(blockHash_)
	}

	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)                // 1
		fmt.Println(receipt.Logs)                  // []
		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex)      // 0
		fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
		break
	}

	return receiptByHash
}

/**
 * 通过区块高度查询收据
 */
func SelectReceiptsByBlockNumber(blockNumber_ int64) []*types.Receipt {
	client := client2.ClientEthereum("")

	// 默认使用一个测试的区块
	blockNumber := big.NewInt(5671744)

	if blockNumber_ != 0 {
		blockNumber = big.NewInt(blockNumber_)
	}

	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)                // 1
		fmt.Println(receipt.Logs)                  // []
		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex)      // 0
		fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
		break
	}

	return receiptByHash
}

/**
 * 通过交易哈希查询对应的收据
 */

func SelectReceiptsByTxHash(txHash_ string) *types.Receipt {
	client := client2.ClientEthereum("")

	// 没有传区块哈希使用默认值
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")

	if txHash_ != "" {
		txHash = common.HexToHash(txHash_)
	}

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)                // 1
	fmt.Println(receipt.Logs)                  // []
	fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	fmt.Println(receipt.TransactionIndex)      // 0
	fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000

	return receipt
}

func testReceipts() {
	fmt.Println("测试通过区块高度查询收据")
	SelectReceiptsByBlockNumber(0)

	fmt.Println("测试通过区块哈希查询收据")
	SelectReceiptsByBlockHash("")

	fmt.Println("测试通过交易哈希查询收据")
	SelectReceiptsByTxHash("")
}
