package ethereum_operators

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	client2 "go_demo/client"
	token "go_demo/contracts/erc20"
	"log"
	"math"
	"math/big"
)

func SelectTokenBalanceByAddress(contractAddress_ string, account_address_ string) {
	client := client2.ClientEthereum("")
	contractAddress := "0x500E478001eD7922Ab68E34c3284C85Cac3F10E6"
	if contractAddress_ != "" {
		contractAddress = contractAddress_
	}

	// 通过合约地址，拿到合约实例。
	// Golem (GNT) Address
	tokenAddress := common.HexToAddress(contractAddress)
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	account_address := "0xc1D863021f2dE7661A1666a6c4925993bDB8135C"
	if account_address_ != "" {
		account_address = account_address_
	}

	address := common.HexToAddress(account_address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}

func testSelectTokenBalance() {
	fmt.Println("测试查询代币余额")
	SelectTokenBalanceByAddress("", "")
}
