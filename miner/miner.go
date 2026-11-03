package miners

import (
	"blockchain/blockchain"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func MineBlock(bc *blockchain.Blockchain, block blockchain.Block) bool {

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	var nonce int64 = 0

	newBlock := blockchain.CreateBlock(lastBlock, block.Data)

	newBlock.Difficulty = bc.Difficulty

	Data, _ := json.Marshal(block.Data)

	startTime := time.Now()

	target := strings.Repeat("0", newBlock.Difficulty)

	for newBlock.Hash[:len(target)] != target {

		nonce++

		newBlock.Nonce = nonce

		input := newBlock.Header + string(Data) + newBlock.PreviousHash + strconv.FormatInt(nonce, 10) + strconv.Itoa(newBlock.Index)

		hash := sha256.Sum256([]byte(input))

		newBlock.Hash = hex.EncodeToString(hash[:])

	}

	calculatedTime := time.Since(startTime)

	if calculatedTime.Seconds() < 5 && bc.Difficulty < 6 {
		bc.Difficulty++
	} else if calculatedTime.Seconds() > 15 && bc.Difficulty > 1 {
		bc.Difficulty--
	}

	fmt.Println("Time to find nonce: ", calculatedTime)

	fmt.Println(newBlock.Hash, newBlock.Nonce)

	bc.Blocks = append(bc.Blocks, newBlock)

	bc.ValidateChain()

	return true
}
