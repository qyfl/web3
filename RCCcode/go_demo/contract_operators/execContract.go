package contract_operators

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go_demo/client"
	"go_demo/contracts/store"
	"go_demo/env"
	"log"
	"math/big"
)

func ExecContract(contractAddr_ string, privateKey_ string) {
	client := client.ClientEthereum("")

	contractAddr := "0xD289907149F6F396FEdfE0552150926Dba8f6b1A"
	if contractAddr_ != "" {
		contractAddr = contractAddr_
	}

	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	pk := env.PRIVATE_KEY

	if privateKey_ != "" {
		pk = privateKey_
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
}
func TestExecContract() {
	fmt.Println("开始测试执行合约的操作")
	ExecContract("", "")
}
