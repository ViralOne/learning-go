package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
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

	airdorp_txhas, _ := wallet.RequestAirdrop(1 * 1e9)
	fmt.Println("Transaction URL:", "https://explorer.solana.com/tx/"+airdorp_txhas+"?cluster=devnet")

	// Check balance of the wallet right after airdrop it will be 0 because I thnk it has no yet all the confirmations
	balance, _ := wallet.get_balance()
	fmt.Println("Balance: ", balance/1e9, "SOL")

	// Check balance of the wallet after 30 seconds because this is approximately when it will have all the confirmations
	time.Sleep(time.Second * 30)
	balance, _ = wallet.get_balance()
	fmt.Println("Balance after 30 sec: ", balance/1e9, "SOL")

	// Sending 0.9SOL to the AirDrop address
	transfer_log, _ := wallet.Transfer("9B5XszUGdMaxCZ7uSQhPzdks5ZQSmWxrmzCSvtJ6Ns6g", 1*1e9-1e8)
	fmt.Println("Transaction URL:", "https://explorer.solana.com/tx/"+transfer_log+"?cluster=devnet")
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

func (w Wallet) RequestAirdrop(amount uint64) (string, error) {
	txhash, err := w.client.RequestAirdrop(
		context.TODO(),
		w.account.PublicKey.ToBase58(), // wallet address requesting airdrop
		amount,                         // amount of SOL in lamport
	)
	if err != nil {
		return "", err
	}

	return txhash, nil
}

func (w Wallet) Transfer(receiver string, amount uint64) (string, error) {
	// fetch the most recent blockhash
	response, err := w.client.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	// make a transfer message with the latest block hash
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{w.account, w.account},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        w.account.PublicKey,
			RecentBlockhash: response.Blockhash,
			Instructions: []types.Instruction{
				sysprog.Transfer(sysprog.TransferParam{
					From:   w.account.PublicKey,
					To:     common.PublicKeyFromString(receiver),
					Amount: amount,
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to new a transaction, err: %v", err)
	}
	// send the transaction to the blockchain
	txhash, err := w.client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}

	return txhash, nil
}
