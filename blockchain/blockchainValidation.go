package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func ValidateBlock(block Block) bool {

	transactionData, _ := json.Marshal(block.Transactions)
	Data, _ := json.Marshal(block.Data)

	input := block.Header + string(Data) + string(transactionData) + block.PreviousHash + strconv.FormatInt(block.Nonce, 10) + strconv.Itoa(block.Index)

	hash := sha256.Sum256([]byte(input))

	calculatedHash := hex.EncodeToString(hash[:])

	fmt.Println("Block number: ", block.Index)
	fmt.Println("Stored hash:    ", block.Hash)
	fmt.Println("Calculated hash:", calculatedHash)

	if block.Hash == calculatedHash {
		fmt.Println("Blockchain alterations: No")
		return true
	} else {
		fmt.Println("Blockchain alterations: Yes")
		return false
	}
}

func (bc *Blockchain) ValidateChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		if current.PreviousHash != previous.Hash {
			return false
		}

		if !ValidateBlock(current) {
			return false
		}

		if current.Difficulty < 1 ||
			current.Difficulty > len(current.Hash) {
			return false
		}

		target := strings.Repeat("0", current.Difficulty)

		if !strings.HasPrefix(current.Hash, target) {
			return false
		}

		fmt.Println("______________                       ______________")
		fmt.Println("Transaction validity is: ", ValidateBlockTransactions(bc.Blocks[i]))
		fmt.Println("Index validity is: ", bc.ValidateBlockIndex())
		fmt.Println("Difficulty validity is: ", ValidateBlockDifficulty(bc.Blocks[i]))
		fmt.Println("___________________________________________________")

	}

	fmt.Println("Genesis Block validity: ", bc.ValidateGenesis(bc.Blocks[0]))

	return true
}

func (bc *Blockchain) ValidateGenesis(block Block) bool {
	if len(bc.Blocks) == 0 || bc.Blocks[0].Index != 0 {
		return false
	}
	return true
}
