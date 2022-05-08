package main

import (
		"context"
        "fmt"
		"github.com/portto/solana-go-sdk/types"
        "github.com/portto/solana-go-sdk/client"
        "github.com/portto/solana-go-sdk/rpc"
)

type Wallet struct {
	account types.Account
	client  *client.Client
}

func main() {
	wallet := create_wallet(rpc.DevnetRPCEndpoint)
	wallet_address := wallet.account.PublicKey
	wallet_private_key := wallet.account.PrivateKey
	fmt.Println(wallet_address)
	fmt.Println(wallet_private_key)

	balance, _ := wallet.get_balance()
	fmt.Println(balance)
}

func create_wallet(RPCEndpoint string) Wallet {
	return Wallet{
		types.NewAccount(),
		client.NewClient(RPCEndpoint),
	}
}

func import_wallet(privateKey []byte, RPCEndpoint string) (Wallet, error) {
	wallet, error := types.AccountFromBytes(privateKey)
	if error != nil {
		return Wallet{}, error
	}

	return Wallet{
		wallet,
		client.NewClient(RPCEndpoint),
	}, nil
}

// Get balance of the wallet
func (w Wallet) get_balance() (uint64, error) {
	balance, err := w.client.GetBalance(
			context.TODO(),
			w.account.PublicKey.ToBase58(), // wallet to fetch balance for
	)
	if err != nil {
			return 0, nil
	}

	return balance, nil
}

// Check balance of an public address
func check_balance(wallet string, RPCEndpoint string) (uint64, error) {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	balance, err := c.GetBalance(
			context.TODO(),
			wallet, // wallet to fetch balance for
	)
	if err != nil {
		panic(err)
	}
	return balance, nil
}

// How to use:
// balance, _ := check_balance("9B5XszUGdMaxCZ7uSQhPzdks5ZQSmWxrmzCSvtJ6Ns6g", rpc.DevnetRPCEndpoint)
// fmt.Println(balance)
