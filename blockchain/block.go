package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Data struct {
	BatteryID      string    `json:"batteryID"`
	Voltage        float64   `json:"voltage"`
	Temperature    float64   `json:"temperature"`
	StateOfCharge  float64   `json:"stateOfCharge"`
	ProductionTime time.Time `json:"productionStamp"`
	Status         string    `json:"status"`
}

type Block struct {
	Index          int           `json:"index"`
	Header         string        `json:"header"`
	TimeStamp      time.Time     `json:"timeStamp"`
	Data           Data          `json:"data"`
	Hash           string        `json:"hash"`
	PreviousHash   string        `json:"previousHash"`
	Nonce          int64         `json:"nonce"`
	Difficulty     int           `json:"difficulty"`
	MiningDuration time.Duration `json:"miningDuration"`
}

func CreateGenesisBlock() Block {

	block := Block{
		Index:        0,
		Header:       "",
		TimeStamp:    time.Now(),
		Data:         Data{},
		PreviousHash: "",
		Nonce:        0,
		Difficulty:   3,
	}

	target := strings.Repeat("0", block.Difficulty)

	for {
		block.Hash = CalculateHash(block)
		if strings.HasPrefix(block.Hash, target) {
			break
		}
		block.Nonce++
	}

	return block
}

func CreateBlock(previousBlock Block, data Data) Block {

	block := Block{
		Index:        previousBlock.Index + 1,
		Header:       "",
		TimeStamp:    time.Now(),
		Data:         data,
		PreviousHash: previousBlock.Hash,
		Nonce:        0,
		Difficulty:   3,
	}

	block.Hash = CalculateHash(block)

	return block
}

func CalculateHash(block Block) string {

	Data, _ := json.Marshal(block.Data)

	input := block.Header + string(Data) + block.PreviousHash + strconv.FormatInt(block.Nonce, 10) + strconv.Itoa(block.Index) + strconv.Itoa(block.Difficulty) + block.TimeStamp.Format(time.RFC3339Nano)

	hash := sha256.Sum256([]byte(input))

	return hex.EncodeToString(hash[:])
}
