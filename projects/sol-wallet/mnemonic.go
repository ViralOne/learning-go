package main

import (
		"context"
        "fmt"
		"github.com/portto/solana-go-sdk/types"
        "github.com/portto/solana-go-sdk/client"
        "github.com/portto/solana-go-sdk/rpc"
		"github.com/tyler-smith/go-bip39"
		"github.com/tyler-smith/go-bip32"
)

type Wallet struct {
	account types.Account
	client  *client.Client
}

var public_address string

func main(){
	import_wallet("attract aim enlist stadium zone exhaust you close like decorate such spread wide click female worry woman junior phone speed same sell limb better")
	//generate_wallet()
	//RequestAirdrop(public_address, 1*1e9)

	balance, _ := check_balance(public_address, rpc.DevnetRPCEndpoint)
	fmt.Println("Balance: ", balance/1e9, "SOL")
}

func generate_wallet() {
  entropy, _ := bip39.NewEntropy(256)
  mnemonic, _ := bip39.NewMnemonic(entropy)

  // Generate a Bip32 HD wallet for the mnemonic and a user supplied password
  seed := bip39.NewSeed(mnemonic, "super-secret-phrase-for-my-wallet")
  account, _ := types.AccountFromSeed(seed[:32])
  public_address = account.PublicKey.ToBase58()

  masterKey, _ := bip32.NewMasterKey(seed)
  publicKey := masterKey.PublicKey()

  // Display mnemonic and keys
  fmt.Println("Mnemonic: ", mnemonic)
  fmt.Println("SOL adress: ", public_address)
  fmt.Println("Master private key: ", masterKey)
  fmt.Println("Master public key: ", publicKey)
}

func import_wallet(mnemonic string){
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

func RequestAirdrop(wallet string, amount uint64){
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
