package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func CreateWallet() string {

	// 生成钱包
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	return getWalletAddress(privateKey)

}

func GetWalletByPrivateKey(privateKey_ string) string {

	key := "ccec5314acec3d18eae81b6bd988b844fc4f7f7d3c828b351de6d0fede02d3f2"
	if privateKey_ != "" {
		key = privateKey_
	}

	// 导入已有的私钥
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}

	return getWalletAddress(privateKey)
}

func getWalletAddress(privateKey *ecdsa.PrivateKey) string {

	// 把私钥转成字节
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println("私钥:", hexutil.Encode(privateKeyBytes)[2:])

	// 使用私钥生成公钥
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("公钥:", hexutil.Encode(publicKeyBytes)[4:])

	// 把公钥转成钱包地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("地址", address)

	return address
}

func testWallet() {
	fmt.Println("测试生成钱包")
	CreateWallet()
	fmt.Println("测试使用默认密钥创建钱包")
	GetWalletByPrivateKey("")
}
