package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/Ankr-network/coqchain/accounts/abi/bind"
	"github.com/Ankr-network/coqchain/common"
	"github.com/Ankr-network/coqchain/core/contracts/staking/staker"
	"github.com/Ankr-network/coqchain/core/types"
	"github.com/Ankr-network/coqchain/crypto"
	"github.com/Ankr-network/coqchain/ethclient"
)

var (
	contractAddress = "0x000000000000000000000000000000000000face"
	e               = flag.String("exe", "d", "exeute which case")
	taddr           = common.HexToAddress("0x4915f56a21F1f2e651f8130c5a9257Cd429c6136")
)

func main() {
	flag.Parse()

	switch *e {
	case "d":
		depost()
	case "b":
		balance()
	default:
		log.Println("nout supported: ", *e)
	}

}

func balance() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	addr := common.HexToAddress(contractAddress)
	instance, _ := staker.NewStaker(addr, client)

	rs, err := instance.Balances(&bind.CallOpts{}, taddr)
	if err != nil {
		log.Println("get epoch failed: ", err)
		return
	}

	fmt.Printf("addr: %v balance : %d \n", taddr, rs)
}

func depost() {
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

	auth.Value = big.NewInt(0).Mul(big.NewInt(1), big.NewInt(1e+18)) // in ether
	auth.GasLimit = uint64(3000000)                                  // in units
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
	instance, err := staker.NewStaker(addr, client)

	tx, err := instance.Stake(auth)
	if err != nil {
		log.Printf("deposit err: %v \n", err)
		return
	}
	log.Printf("tx hash: %s \n", tx.Hash())
}
