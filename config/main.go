package main

import (
	"blockchain/blockchain"
	wallet "blockchain/crypto"
	miners "blockchain/miner"
	"fmt"
	"time"
)

func main() {

	chain := blockchain.CreateBlockchain() // blockchain creation

	fmt.Println("Genesis hash in chain:", chain.Blocks[0].Hash)

	myWallet := wallet.CreateWallet() // wallet creation

	tx := wallet.CreateTransaction( // transaction creation
		myWallet,
		"Filippos",
		"NoOne",
		25000,
		"",
		myWallet.PrivateKey.PublicKey,
	)

	signature := wallet.SignTransaction(tx, myWallet.PrivateKey) // signature creation

	tx.Signature = signature

	chain.AddTransaction(tx, signature, &myWallet.PublicKey) // add transaction on chain

	valid := wallet.VerifyTransaction(
		tx,
		tx.Signature,
		&myWallet.PublicKey,
	)

	for i := 0; i < 3; i++ {
		fmt.Println("Mining block: ", i)
		record := blockchain.Data{
			BatteryID:      fmt.Sprintf("Bat.%03d", i+1),
			Voltage:        3.7 - float64(i)*0.01,
			Temperature:    25.0 + float64(i)*0.2,
			StateOfCharge:  80.0 - float64(i),
			ProductionTime: time.Now(),
		}

		previousBlock := chain.Blocks[len(chain.Blocks)-1]
		newBlock := blockchain.CreateBlock(previousBlock, record)

		miners.MineBlock(&chain, newBlock)
	}
	// chain.Blocks[1].Index = 99 //test for index tracking
	// fmt.Println("Chain valid after index tampering:", chain.ValidateChain())

	// chain.Blocks[1].Data = "Tampered data"								- check for tampered data
	// chain.Blocks[2].PreviousHash = "wrong previous hash"					- check for previous hash alteration

	fmt.Println(tx)
	fmt.Println("___________________________________________________")
	fmt.Println(chain)
	fmt.Println("___________________________________________________")
	fmt.Println(myWallet.PrivateKey)
	fmt.Println("___________________________________________________")
	fmt.Println(myWallet.PublicKey)
	fmt.Println("___________________________________________________")
	fmt.Println("Transaction: ", tx)
	fmt.Println("___________________________________________________")
	fmt.Println("Signature: ", tx.Signature)
	fmt.Println("___________________________________________________")
	fmt.Println("Valid: ", valid)
	fmt.Println("___________________________________________________")

	if wallet.VerifyTransaction(tx, tx.Signature, &myWallet.PublicKey) == true { // transaction verification
		fmt.Println("Validation: Successful")
	} else {
		fmt.Println("Validation: Unsuccessful")
	}
	fmt.Println("___________________________________________________")
	for i, tx := range chain.Mempool {
		fmt.Println("Transaction index:", i)
		fmt.Println("Sender:", tx.SenderID)
		fmt.Println("Receiver:", tx.ReceiverID)
		fmt.Println("Amount:", tx.Amount)
		fmt.Println("Signature:", tx.Signature)
	}
	fmt.Println("___________________________________________________")
	fmt.Println("Blocks: ", len(chain.Blocks))
	fmt.Println("Pending transactions: ", len(chain.Mempool))
	fmt.Println("Transactions in latest block: ", len(chain.Blocks[len(chain.Blocks)-1].Transactions))
	fmt.Println("___________________________________________________")
	fmt.Println(chain.ValidateChain())

}
