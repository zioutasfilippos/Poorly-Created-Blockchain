package blockchain

import (
	wallet "blockchain/crypto"
	"crypto/ecdsa"
	"fmt"
)

type Blockchain struct {
	Blocks     []Block
	Mempool    []wallet.Transaction
	Difficulty int
}

func CreateBlockchain() Blockchain {

	genesis := CreateGenesisBlock()

	blockchain := Blockchain{
		Blocks:     []Block{genesis},
		Difficulty: 3,
	}

	return blockchain
}

func (bc *Blockchain) AddTransaction(tx wallet.Transaction, Signature string, publicKey *ecdsa.PublicKey) bool {

	if wallet.VerifyTransaction(tx, Signature, publicKey) == true {
		for _, existingTx := range bc.Mempool {
			if tx.SenderID == existingTx.SenderID &&
				tx.ReceiverID == existingTx.ReceiverID &&
				tx.Amount == existingTx.Amount &&
				tx.Signature == existingTx.Signature {
				return false
			}
		}
		bc.Mempool = append(bc.Mempool, tx)
		fmt.Println("Transaction has been placed in the mempool")
		return true
	} else {
		err := "Cant be verified"
		fmt.Println(err)
		return false
	}
}

func (bc *Blockchain) MineBlock() {

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := CreateBlock(lastBlock, Data{})

	bc.Blocks = append(bc.Blocks, newBlock)

	bc.Mempool = []wallet.Transaction{}
}

func ValidateBlockTransactions(block Block) bool {

	for _, tx := range block.Transactions {

		if tx.Amount <= 0 || wallet.VerifyTransaction(tx, tx.Signature, &tx.SendersPublicKey) == false {
			return false
		}
	}

	return true
}

func (bc *Blockchain) ValidateBlockIndex() bool {
	for i := 1; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]

		previous := bc.Blocks[i-1]

		if current.Index != previous.Index+1 {
			return false
		}
	}
	return true
}

func ValidateBlockDifficulty(block Block) bool {

	if block.Difficulty > 6 || block.Difficulty < 1 {
		return false
	}

	return true
}
