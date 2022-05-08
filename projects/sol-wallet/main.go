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
