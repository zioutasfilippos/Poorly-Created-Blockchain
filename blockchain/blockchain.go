package blockchain

import (
	"time"
)

type Blockchain struct {
	Blocks     []Block
	Difficulty int
}

func CreateBlockchain() *Blockchain {

	genesis := CreateGenesisBlock()

	blockchain := Blockchain{
		Blocks:     []Block{genesis},
		Difficulty: 3,
	}

	return &blockchain
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

	if block.Index == 0 {
		return block.Difficulty == 0
	}
	if block.Difficulty > 6 || block.Difficulty < 1 {
		return false
	}

	return true
}

func (bc *Blockchain) AverageBlockTime() time.Duration {
	if len(bc.Blocks) < 4 {
		return 0
	}

	total := time.Duration(0)

	for i := len(bc.Blocks) - 3; i < len(bc.Blocks); i++ {
		timeDifference := bc.Blocks[i].TimeStamp.Sub(bc.Blocks[i-1].TimeStamp)
		total += timeDifference
	}

	return total / 3
}

func (bc *Blockchain) FindBattery(BatteryID string) (Block, bool) {
	for i := 0; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].Data.BatteryID == BatteryID {
			return bc.Blocks[i], true
		}
	}

	return Block{}, false
}

func (data *Data) GetBatteryStatus(batteryID string) string {

	if data.StateOfCharge > 50 {
		return "Normal"
	} else if data.StateOfCharge > 15 && data.StateOfCharge < 50 {
		return "Warnign"
	} else {
		return "critical"
	}

}

func (bc *Blockchain) GetLatestBatteries() []Block {
	latest := make(map[string]Block)

	for _, block := range bc.Blocks {
		batteryID := block.Data.BatteryID
		if batteryID != "" {
			latest[batteryID] = block
		}
	}

	result := make([]Block, 0, len(latest))
	for _, block := range latest {
		result = append(result, block)
	}

	return result
}

func (bc *Blockchain) AdjustDifficulty() {
	average := bc.AverageBlockTime()

	if average == 0 {
		return
	}

	if average < 5*time.Second && bc.Difficulty < 6 {
		bc.Difficulty++
	} else if average > 15*time.Second && bc.Difficulty > 1 {
		bc.Difficulty--
	}
}

func (bc *Blockchain) CountUniqueatteries() int {
	count := 0
	BatCount := make(map[string]bool)

	for i := 0; i < len(bc.Blocks); i++ {
		ID := bc.Blocks[i].Data.BatteryID

		if ID == "" {
			continue
		}

		_, exists := BatCount[ID]

		if !exists {
			count++
			BatCount[ID] = true
		}

	}
	return count
}

func (bc *Blockchain) FindBatteryStatus(status string) []Block {

	LastStatus := make(map[string]Block)

	for _, block := range bc.Blocks {
		batteryID := block.Data.BatteryID
		if batteryID != "" {
			LastStatus[batteryID] = block
		}
	}

	var result []Block

	for _, block := range LastStatus {
		if block.Data.Status == status {
			result = append(result, block)
		}
	}

	return bc.Blocks
}
