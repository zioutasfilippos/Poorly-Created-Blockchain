package tests

import (
	"os"
	"testing"
	"time"

	"blockchain/blockchain"
)

func TestSaveAndLoadBlockchain(t *testing.T) {
	bc := blockchain.CreateBlockchain()

	previous := bc.Blocks[len(bc.Blocks)-1]
	data := blockchain.Data{
		BatteryID:      "BAT-STORAGE-001",
		Voltage:        3.7,
		Temperature:    25.0,
		StateOfCharge:  80.0,
		ProductionTime: time.Now(),
	}

	block := blockchain.CreateBlock(previous, data)
	block.Difficulty = bc.Difficulty
	block = mineTestBlock(block)

	if err := blockchain.AddBlock(bc, block); err != nil {
		t.Fatalf("AddBlock failed: %v", err)
	}

	// We'll test saving and loading here next.
	_ = os.Remove
}
