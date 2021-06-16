package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	zbot "../contracts/build"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://coston.flare.network:9650/ext/bc/C/rpc")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(1 * int64(math.Pow(10, 17))) // CHANGE THIS TO AMOUNT OF WEI YOU WANT TO DEPOSIT
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xCBB39D17B5a45198d57586e86EF05C715a3ce025") // This is address of contract
	instance, err := zbot.NewZbot(address, client)
	if err != nil {
		log.Fatal(err)
	}
	id := "discord 401085830220742656" // CHANGE THIS TO YOUR ID
	tx, err := instance.Deposit(auth, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	//fmt.Println(string(tx.Data()))
	result, err := instance.Balance(nil, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance of " + id + ": " + result.String()) // "bar"
	//fmt.Println("Address of "+id+": ", result.Hex())
	/*
		txHash := common.HexToHash("0x96b652623aadfbe2055f435e0b9561e20d28a85d3af09cdd0c03f4e60d911f80")
		tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Transaction:", string(tx.Data()))
		fmt.Println(isPending)
	*/
}
