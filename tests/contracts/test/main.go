package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/Ankr-network/coqchain/accounts/abi/bind"
	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/core/types"
	"github.com/Ankr-network/coqchain/crypto"
	"github.com/Ankr-network/coqchain/ethclient"
	"github.com/Ankr-network/coqchain/tests/contracts/staker"
)

var (
	wsaddr          = "ws://127.0.0.1:8546"
	contractAddress = common.HexToAddress("0x3049CD94e70bB585bd6681f84B154725d84D26E9")
	// contractAddress = common.HexToAddress("0x924b48c3396B658CF29344e48f0ae1dEA089F167")
	depositAmt = big.NewInt(8e+18)
)

func main() {
	client, err := ethclient.Dial(wsaddr)
	if err != nil {
		log.Fatal("create client: ", err)
	}

	privateKey, err := crypto.HexToECDSA("17347d134755c7c27133fcfd3817e4302b05ba2253dcafaf21c2ebb3b5ddc47e")
	if err != nil {
		log.Fatal(err)
	}

	// deploy(client, privateKey)
	// deposit(client, privateKey)
	withdraw(client, privateKey)

}

func deploy(client *ethclient.Client, pk *ecdsa.PrivateKey) {
	// create new singer
	signer := types.NewLondonSigner(big.NewInt(1337))
	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("get nonce failed: ", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("get gas price: ", err)
	}

	auth := bind.NewKeyedTransactor(pk)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice
	auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != fromAddress {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), pk)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}

	addr, tx, _, err := staker.DeployStaker(auth, client)
	if err != nil {
		log.Fatal("deploy failed: ", err)
	}

	log.Printf("addr: %s tx: %s \n", addr, tx.Hash())

}

func deposit(client *ethclient.Client, pk *ecdsa.PrivateKey) {

	// create new singer
	signer := types.NewLondonSigner(big.NewInt(1337))
	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)
	auth := bind.NewKeyedTransactor(pk)
	auth.Value = depositAmt
	auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != fromAddress {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), pk)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}

	inst, err := staker.NewStaker(contractAddress, client)
	if err != nil {
		log.Fatal("new instance failed: ", err)
	}
	sink := make(chan *staker.StakerConsume)
	tx, err := inst.Consume(auth)
	if err != nil {
		log.Fatal("tx failed: ", err)
	}
	log.Printf("tx: %s \n", tx.Hash())

	sub, err := inst.WatchConsume(&bind.WatchOpts{}, sink)
	if err != nil {
		log.Printf("address: %s watch approve failed: %v \n", fromAddress, err)
		return
	}
	defer sub.Unsubscribe()
	for t := range sink {
		if t.Sender == fromAddress && t.Amt.Cmp(depositAmt) == 0 {
			break
		}
	}
	log.Printf("deposit amout: %d over\n", depositAmt)
}

func withdraw(client *ethclient.Client, pk *ecdsa.PrivateKey) {
	// create new singer
	signer := types.NewLondonSigner(big.NewInt(1337))
	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)
	auth := bind.NewKeyedTransactor(pk)
	auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != fromAddress {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), pk)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}

	inst, err := staker.NewStaker(contractAddress, client)
	if err != nil {
		log.Fatal("new instance failed: ", err)
	}
	sink := make(chan *staker.StakerWithdraw)
	tx, err := inst.Withdraw(auth)
	if err != nil {
		log.Fatal("tx failed: ", err)
	}
	log.Printf("tx: %s \n", tx.Hash())

	sub, err := inst.WatchWithdraw(&bind.WatchOpts{}, sink)
	if err != nil {
		log.Printf("address: %s watch approve failed: %v \n", fromAddress, err)
		return
	}
	defer sub.Unsubscribe()
	for t := range sink {
		if t.To == fromAddress {
			break
		}
	}
	log.Printf("withdraw  over\n")

}
