package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/Ankr-network/coqchain/accounts/abi/bind"
	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/core/types"
	"github.com/Ankr-network/coqchain/crypto"
	"github.com/Ankr-network/coqchain/ethclient"
	"github.com/Ankr-network/coqchain/tests/contracts/store"
)

var (
	contractAddress = "0x3049CD94e70bB585bd6681f84B154725d84D26E9"
	testKey         = "hello"
	testVal         = "world"
)

func main() {
	deploy()
	// getItem()
	// setItem()
}

func deploy() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("17347d134755c7c27133fcfd3817e4302b05ba2253dcafaf21c2ebb3b5ddc47e")
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

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Printf("get chain id failed: %v \n", err)
		return
	}

	// create new singer
	signer := types.NewLondonSigner(chainid)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != fromAddress {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract address:", address.Hex())
	fmt.Println("tx hash: ", tx.Hash().Hex())

	_ = instance

}
func getItem() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	addr := common.HexToAddress(contractAddress)
	instance, err := store.NewStore(addr, client)
	// set key value
	rsp, err := instance.GetItem(&bind.CallOpts{}, testKey)
	if err != nil {
		fmt.Printf("get error: %s \n", err)
		return
	}
	fmt.Println("get item ", rsp)
}

func setItem() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("17347d134755c7c27133fcfd3817e4302b05ba2253dcafaf21c2ebb3b5ddc47e")
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

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Printf("get chain id failed: %v \n", err)
		return
	}

	// create new singer
	signer := types.NewLondonSigner(chainid)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != fromAddress {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}

	addr := common.HexToAddress(contractAddress)
	instance, err := store.NewStore(addr, client)
	// set key value
	tx, err := instance.SetItem(auth, testKey, testVal)
	if err != nil {
		fmt.Printf("get error: %s \n", err)
		return
	}
	fmt.Println("set item ", tx.Hash().Hex())
}
