package tests

import (
	"strings"
	"testing"
	"time"

	"blockchain/blockchain"
)

func TestAddBlockRejectsWrongPreviousHash(t *testing.T) {
	bc := blockchain.CreateBlockchain()
	previous := bc.Blocks[len(bc.Blocks)-1]

	data := blockchain.Data{
		BatteryID:      "BAT-WRONG-HASH",
		Voltage:        3.7,
		Temperature:    25.0,
		StateOfCharge:  80.0,
		ProductionTime: time.Now(),
	}

	block := blockchain.CreateBlock(previous, data)
	block.Difficulty = bc.Difficulty
	block.PreviousHash = "incorrect-previous-hash"

	err := blockchain.AddBlock(bc, block)
	if err == nil {
		t.Fatal("Expected an error for an incorrect previous hash, got nil")
	}

	if len(bc.Blocks) != 1 {
		t.Fatalf("Expected 1 block after rejection, got %d", len(bc.Blocks))
	}
}

func TestAddBlockRejectsWrongIndex(t *testing.T) {
	bc := blockchain.CreateBlockchain()
	previous := bc.Blocks[len(bc.Blocks)-1]

	data := blockchain.Data{
		BatteryID:      "BAT-WRONG-INDEX",
		Voltage:        3.7,
		Temperature:    25.0,
		StateOfCharge:  80.0,
		ProductionTime: time.Now(),
	}

	block := blockchain.CreateBlock(previous, data)
	block.Difficulty = bc.Difficulty
	block.Index = previous.Index + 2

	err := blockchain.AddBlock(bc, block)
	if err == nil {
		t.Fatal("Expected an error for an incorrect index, got nil")
	}

	if len(bc.Blocks) != 1 {
		t.Fatalf("Expected 1 block after rejection, got %d", len(bc.Blocks))
	}
}

func mineTestBlock(block blockchain.Block) blockchain.Block {
	target := strings.Repeat("0", block.Difficulty)

	for {
		block.Hash = blockchain.CalculateHash(block)

		if strings.HasPrefix(block.Hash, target) {
			return block
		}

		block.Nonce++
	}
}

func TestAddBlockRejectsDuplicateData(t *testing.T) {
	bc := blockchain.CreateBlockchain()
	previous := bc.Blocks[len(bc.Blocks)-1]

	data := blockchain.Data{
		BatteryID:      "BAT-DUPLICATE",
		Voltage:        3.7,
		Temperature:    25.0,
		StateOfCharge:  80.0,
		ProductionTime: time.Now(),
	}

	// Mine and add the first block.
	block := blockchain.CreateBlock(previous, data)
	block.Difficulty = bc.Difficulty
	block = mineTestBlock(block)

	if err := blockchain.AddBlock(bc, block); err != nil {
		t.Fatalf("Could not add first block: %v", err)
	}

	// Create another block containing exactly the same data.
	previous = bc.Blocks[len(bc.Blocks)-1]
	duplicate := blockchain.CreateBlock(previous, data)
	duplicate.Difficulty = bc.Difficulty
	duplicate = mineTestBlock(duplicate)

	err := blockchain.AddBlock(bc, duplicate)
	if err == nil {
		t.Fatal("Expected duplicate data to be rejected, got nil")
	}

	if len(bc.Blocks) != 2 {
		t.Fatalf("Expected 2 blocks after rejection, got %d", len(bc.Blocks))
	}
}
