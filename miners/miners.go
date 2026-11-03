package miners

import (
	"blockchain/blockchain"
	"errors"
	"strings"
	"time"
)

func MineBlock(bc *blockchain.Blockchain, data blockchain.Data) error {

	if len(bc.Blocks) == 0 {
		return errors.New("Blockchain has no blocks.")
	}

	if data.BatteryID == "" {
		return errors.New("Cant create block without battery ID field.")
	}

	startTime := time.Now()

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := blockchain.CreateBlock(lastBlock, data)

	newBlock.Difficulty = bc.Difficulty

	target := strings.Repeat("0", newBlock.Difficulty)

	for {
		newBlock.Hash = blockchain.CalculateHash(newBlock)

		if strings.HasPrefix(newBlock.Hash, target) {
			break
		}

		newBlock.Nonce++
	}

	finishedTime := time.Since(startTime)
	newBlock.MiningDuration = finishedTime

	if bc.ValidateChain() == false {
		return errors.New("Error validating the chain.")
	}

	if err := blockchain.AddBlock(bc, newBlock); err != nil {
		return err
	}

	bc.AdjustDifficulty()

	return nil
}
