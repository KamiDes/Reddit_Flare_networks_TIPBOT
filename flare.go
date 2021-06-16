package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	zbot "./contracts/build"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var host = "https://costone.flare.network/ext/bc/C/rpc"

func currentBlock() string {
	client, err := ethclient.Dial(host)
	if err != nil {
		return ""
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return ""
	}

	return header.Number.String()
}

func loadContract() *zbot.Zbot {
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xCBB39D17B5a45198d57586e86EF05C715a3ce025") // This is address of contract.
	instance, err := zbot.NewZbot(address, client)
	if err != nil {
		log.Fatal(err)
	}
	v, _ := instance.Version(nil)
	fmt.Println("contract is loaded, version ", v)
	return instance
}
func Zregister(instance *zbot.Zbot, addr string, id string) error {
	client, err := ethclient.Dial(host)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("SECRET"))
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	addre := common.HexToAddress(addr)
	_, err = instance.Register(auth, addre, id)
	if err != nil {
		return err
	}
	return nil
}
func Ztip(instance *zbot.Zbot, from string, to string, amount *big.Int) error {
	client, err := ethclient.Dial(host)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("SECRET"))
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.Tip(auth, from, to, amount)
	if err != nil {
		return err
	}
	fmt.Println("Transaction hash of tip: ", tx.Hash().Hex())
	return nil
}
func Zwithdraw(instance *zbot.Zbot, id string) error {
	client, err := ethclient.Dial(host)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("SECRET"))
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, err = instance.Withdraw(auth, id)
	if err != nil {
		return err
	}
	return nil
}
