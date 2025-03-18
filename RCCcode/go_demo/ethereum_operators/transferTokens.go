package ethereum_operators

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	client2 "go_demo/client"
	"go_demo/env"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

/**
 * 发起交易
 * @param from 发送者私钥
 * @param to 接收者地址
 * @param amount 代币数量, 单位是 ETH !!!
 */
func TransferTokens(from string, to string, amount_ int) {
	client := client2.ClientEthereum("")
	ctx := context.Background()

	fromKey := env.PRIVATE_KEY
	if from != "" {
		fromKey = from
	}

	// 私钥
	privateKey, err := crypto.HexToECDSA(fromKey)
	if err != nil {
		log.Fatal(err)
	}

	// 通过私钥生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 通过公钥生成地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 代币的传输不需要 ETH
	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// solidity 交易函数的签名
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	fmt.Println("交易函数签名", hexutil.Encode(methodID)) // 0xa9059cbb

	// 接受代币的地址
	toAddress := common.HexToAddress(to)

	// 不足32字节，左填充0，到32字节。
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	// 代币数量
	amount := new(big.Int)
	amount.SetString("1000000000000000000", amount_) // 100 tokens

	// 不足32字节，左填充0，到32字节。
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 预计 gas limit
	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	// 代币合约地址
	tokenAddress := common.HexToAddress("0x500E478001eD7922Ab68E34c3284C85Cac3F10E6")
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}

func testTransferTokens() {
	fmt.Println("测试代币转账")
	TransferTokens(env.PRIVATE_KEY, "0x7368a94A5bDcf8F1d6616856B639123Ab36C3d2f", 30)
}
