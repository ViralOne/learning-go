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
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	account types.Account
	client  *client.Client
}

var public_address string

func main() {
	//import_wallet("attract aim enlist stadium zone exhaust you close like decorate such spread wide click female worry woman junior phone speed same sell limb better")
	generate_wallet()
	RequestAirdrop(public_address, 1e8)
	time.Sleep(time.Second * 30)
	balance, _ := check_balance(public_address, rpc.DevnetRPCEndpoint)
	fmt.Println("Balance: ", balance/1e9, "SOL")
	time.Sleep(time.Second * 5)

	// To do: get the mnemonic from the wallet...
	// It will fail because of the lack of balance
	transfer_log, _ := Transfer("mnemonic-sender", "super-secret-phrase-for-my-wallet", "mneumonic-reciver")
	fmt.Println("Transaction URL:", "https://explorer.solana.com/tx/"+transfer_log+"?cluster=devnet")
}

func generate_wallet() {
	entropy, _ := bip39.NewEntropy(256)
	// Generate 24 word mnemonic
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed( // A 24 word seed phrase that is used to generate a private key.
		mnemonic, "super-secret-phrase-for-my-wallet")
	account, _ := types.AccountFromSeed(seed[:32])
	public_address = account.PublicKey.ToBase58()

	//masterKey, _ := bip32.NewMasterKey(seed)
	//publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	// fmt.Println("SOL adress: ", public_address)
	// fmt.Println("Master private key: ", masterKey)
	//fmt.Println("Master public key: ", publicKey)
}

func import_wallet(mnemonic string) {
	seed := bip39.NewSeed(mnemonic, "super-secret-phrase-for-my-wallet") // (mnemonic, password)
	account, _ := types.AccountFromSeed(seed[:32])
	public_address = account.PublicKey.ToBase58()
	fmt.Println(public_address)
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

func RequestAirdrop(wallet string, amount uint64) {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	sig, err := c.RequestAirdrop(
		context.TODO(),
		wallet,
		amount,
	)
	if err != nil {
		fmt.Println("failed to request airdrop, err: %v", err)
	}
	fmt.Println(sig)
}

func Transfer(mnemonic string, password string, to string) (string, error) {
	// create a RPC client
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// import a wallet with Devnet balance
	seed := bip39.NewSeed(mnemonic, password) // (mnemonic, password)
	feePayer, _ := types.AccountFromSeed(seed[:32])

	to_seed := bip39.NewSeed(to, password) // (mnemonic, password)
	to_account, _ := types.AccountFromSeed(to_seed[:32])
	to_address := to_account.PublicKey.ToBase58()

	// to fetch recent blockhash
	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	// make a transfer message with the latest block hash
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer, to_account},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				sysprog.Transfer(sysprog.TransferParam{
					From:   to_account.PublicKey,
					To:     common.PublicKeyFromString(to_address),
					Amount: 1e8, // 0.1 SOL
				}),
			},
		}),
	})

	if err != nil {
		log.Fatalf("failed to new a transaction, err: %v", err)
	}

	// send tx
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	log.Println("txhash:", txhash)
	return txhash, nil
}
