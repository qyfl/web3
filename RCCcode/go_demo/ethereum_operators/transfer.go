package ethereum_operators

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	client2 "go_demo/client"
	"go_demo/env"
	"log"
	"math"
	"math/big"
)

/**
 * 发起交易
 * @param from 发送者私钥
 * @param to 合约地址
 * @param amount 交易金额, 单位是 wei !!!
 */
func Transfer(from string, to string, amount int64) {
	if amount <= 0 {
		return
	}

	client := client2.ClientEthereum("")
	ctx := context.Background()

	fromKey := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	if from != "" {
		fromKey = from
	}

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

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 交易金额
	value := big.NewInt(amount) // in wei (1 eth)

	// 转账交易的 gas 上限
	gasLimit := uint64(21000) // in units

	// 设置 gas 费单位，默认是 wei
	//gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

	// 使用区块链建议的 gas 单位
	//gasPrice, err := client.SuggestGasPrice(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}

	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 默认收款地址
	toKey := "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d"
	if from != "" {
		toKey = to
	}
	toAddress := common.HexToAddress(toKey)

	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	// 把交易用私钥签名
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	chainID, err := client.NetworkID(ctx)
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("get header error: %v", err)
	}

	txData := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: header.BaseFee,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
	}
	// 发起交易和签名合二为一
	signedTx, err := types.SignNewTx(privateKey, types.LatestSignerForChainID(chainID), txData)
	if err != nil {
		log.Fatal(err)
	}

	// 把交易发布到区块链网络
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("转账成功，交易哈希: %s", signedTx.Hash().Hex())
}

func testTransfer() {
	fmt.Println("测试转账交易")
	amount := 0.001 * math.Pow(10, 18)
	Transfer(env.PRIVATE_KEY, "0x7FEAc2e17D0F3791835FDeAc5f73B973068f8183", int64(amount))
}
