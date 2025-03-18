package contract_operators

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go_demo/client"
	"go_demo/contracts/store"
	"log"
)

func LoadContract(contractAddr_ string) {
	client := client.ClientEthereum("")

	contractAddr := "0xD289907149F6F396FEdfE0552150926Dba8f6b1A"
	if contractAddr_ != "" {
		contractAddr = contractAddr_
	}

	// 需要把合约映射成 .go 文件
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	_ = storeContract

	fmt.Println("获取合约成功")
}

func TestLoadContract() {
	LoadContract("")
}
