package ethereum_operators

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	client2 "go_demo/client"
	"log"
	"math"
	"math/big"
)

/**
 * 通过账户地址查询余额
 */
func SelectBalanceByAddress(address_ string) *big.Int {

	client := client2.ClientEthereum("")

	address := "0x25836239F7b632635F815689389C537133248edb"
	if address_ != "" {
		address = address_
	}

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("账户 eth 余额:", ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("pendingBalance:", pendingBalance)
	return balance
}

/**
 * 查询通过区块高度，查询在这个区块里的余额
 */
func SelectBalanceByBlockNumber(address_ string, blockNumber_ int64) *big.Int {
	client := client2.ClientEthereum("")

	address := "0x25836239F7b632635F815689389C537133248edb"
	if address_ != "" {
		address = address_
	}

	account := common.HexToAddress(address)

	blockNumber := big.NewInt(5532993)
	if blockNumber_ != 0 {
		blockNumber = big.NewInt(blockNumber_)
	}

	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("账户 eth 余额:", ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("pendingBalance:", pendingBalance)

	return balance
}

func testSelectBalanceByAddress() {
	fmt.Println("测试通过账户地址查询余额")
	SelectBalanceByAddress("0xc1D863021f2dE7661A1666a6c4925993bDB8135C")
}
