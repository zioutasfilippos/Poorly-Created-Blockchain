package tests

import (
	"testing"
	"time"

	"blockchain/blockchain"
)

func TestAddBlockValidBlock(t *testing.T) {
	bc := blockchain.CreateBlockchain()
	previous := bc.Blocks[len(bc.Blocks)-1]

	data := blockchain.Data{
		BatteryID:      "BAT-TEST-001",
		Voltage:        3.7,
		Temperature:    25.0,
		StateOfCharge:  80.0,
		ProductionTime: time.Now(),
	}

	block := blockchain.CreateBlock(previous, data)
	block.Difficulty = bc.Difficulty
	block = mineTestBlock(block)

	err := blockchain.AddBlock(bc, block)
	if err != nil {
		t.Fatalf("Expected valid block to be accepted, got error: %v", err)
	}

	if len(bc.Blocks) != 2 {
		t.Fatalf("Expected 2 blocks, got %d", len(bc.Blocks))
	}
}

func TestAddBlockRejectsEmptyBatteryID(t *testing.T) {
	bc := blockchain.CreateBlockchain()
	previous := bc.Blocks[len(bc.Blocks)-1]

	data := blockchain.Data{
		BatteryID:      "",
		Voltage:        3.7,
		Temperature:    25.0,
		StateOfCharge:  80.0,
		ProductionTime: time.Now(),
	}

	block := blockchain.CreateBlock(previous, data)

	err := blockchain.AddBlock(bc, block)
	if err == nil {
		t.Fatal("Expected an error for an empty BatteryID, got nil")
	}

	if len(bc.Blocks) != 1 {
		t.Fatalf("Expected the blockchain to remain at 1 block, got %d", len(bc.Blocks))
	}
}
