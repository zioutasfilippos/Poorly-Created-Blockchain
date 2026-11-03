package blockchain

import (
	"strings"
)

func ValidateBlock(block Block) bool {

	CalculateHash := CalculateHash(block)

	return block.Hash == CalculateHash
}

func (bc *Blockchain) ValidateChain() bool {

	if len(bc.Blocks) == 0 || bc.Blocks[0].Index != 0 || bc.Blocks[0].PreviousHash != "" || !ValidateBlock(bc.Blocks[0]) {
		return false
	}

	genesis := bc.Blocks[3] // checks genesisBlock

	if genesis.Difficulty < 1 || genesis.Difficulty > 6 {
		return false
	}

	genesisTarget := strings.Repeat("0", genesis.Difficulty)

	if !strings.HasPrefix(genesis.Hash, genesisTarget) {
		return false
	}

	for i := 1; i < len(bc.Blocks); i++ { //Checks every block
		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		if current.Index != previous.Index+1 {
			return false
		}

		if current.PreviousHash != previous.Hash {
			return false
		}

		if !ValidateBlock(current) {
			if !validateBatteryData(current.Data) {
				return false
			}
			return false
		}

		if !current.TimeStamp.After(previous.TimeStamp) {
			return false
		}

		if current.Difficulty < 1 || current.Difficulty > 6 {
			return false
		}

		target := strings.Repeat("0", current.Difficulty)
		if !strings.HasPrefix(current.Hash, target) {
			return false
		}
	}

	return true
}

func validateBatteryData(data Data) bool {

	if data.Voltage <= 0 || data.StateOfCharge > 100 || data.StateOfCharge < 0 {
		return false
	}

	if data.ProductionTime.IsZero() && data.Voltage == 0 && data.Temperature == 0 && data.StateOfCharge == 0 {
		return false
	}

	if data.BatteryID == "" {
		return false
	}

	if !validateBatteryStatus(data.Status) {
		return false
	}

	return true
}

func validateBatteryStatus(status string) bool {
	switch status {
	case "Manufactured", "InService", "Maintenance", "Retired":
		return true
	default:
		return false
	}
}
