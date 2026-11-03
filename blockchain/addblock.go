package blockchain

import (
	"errors"
	"strings"
)

func AddBlock(bc *Blockchain, block Block) error {

	if !validateBatteryData(block.Data) {
		return errors.New("Invalid battery data")
	}

	if len(bc.Blocks) == 0 {
		return errors.New("Blockchain has not blocks.")
	}

	if block.Hash != CalculateHash(block) {
		return errors.New("Missmatched hashes")
	}

	if block.Difficulty < 1 || block.Difficulty > 6 {
		return errors.New("Impossible difficulty")
	}

	target := strings.Repeat("0", block.Difficulty)

	if !strings.HasPrefix(block.Hash, target) {
		return errors.New("Wrong difficulty")
	}

	if block.Difficulty != bc.Difficulty {
		return errors.New("Incorrect diffculty")
	}

	previous := bc.Blocks[len(bc.Blocks)-1]

	if previous.Hash != block.PreviousHash {
		return errors.New("The blockchain lost connection")
	}
	if block.Index != previous.Index+1 {
		return errors.New("Index miss-match")
	}

	for i := 0; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].Data == block.Data {
			return errors.New("Duplicate block")
		}
	}

	bc.Blocks = append(bc.Blocks, block)

	return nil
}

func (bc *Blockchain) FindBatteryHistory(BatteryID string) []Block {

	var history []Block

	for i := 0; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].Data.BatteryID == BatteryID {
			history = append(history, bc.Blocks[i])
		}
	}

	return history
}

func (bc *Blockchain) FindLatestBattery(BatteryID string) (Block, bool) {
	for i := len(bc.Blocks) - 1; i >= 0; i-- {
		if bc.Blocks[i].Data.BatteryID == BatteryID {
			return bc.Blocks[i], true
		}
	}
	return Block{}, false
}
