package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"

	wallet "blockchain/crypto"
)

type Data struct {
	BatteryID      string    `json:"batteryID"`
	Voltage        float64   `json:"voltage"`
	Temperature    float64   `json:"temperature"`
	StateOfCharge  float64   `json:"stateOfCharge"`
	ProductionTime time.Time `json:"productionStamp"`
}

type Block struct {
	Index        int                  `json:"index"`
	Header       string               `json:"header"`
	TimeStamp    time.Time            `json:"timeStamp"`
	Data         Data                 `json:"data"`
	Transactions []wallet.Transaction `json:"transactions"`
	Hash         string               `json:"hash"`
	PreviousHash string               `json:"previousHash"`
	Nonce        int64                `json:"nonce"`
	Difficulty   int                  `json:"difficulty"`
}

func CreateGenesisBlock() Block {

	block := Block{
		Index:        0,
		Header:       "",
		TimeStamp:    time.Now(),
		Data:         Data{},
		Transactions: []wallet.Transaction{},
		Hash:         "hash",
		PreviousHash: "",
	}

	Data, _ := json.Marshal(block.Data)

	input := block.Header + string(Data) + block.PreviousHash + strconv.FormatInt(block.Nonce, 10) + strconv.Itoa(block.Index)

	hash := sha256.Sum256([]byte(input))

	block.Hash = hex.EncodeToString(hash[:])

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
	}

	Data, _ := json.Marshal(block.Data)

	input := block.Header + string(Data) + block.PreviousHash + strconv.FormatInt(block.Nonce, 10) + strconv.Itoa(block.Index)

	hash := sha256.Sum256([]byte(input))

	block.Hash = hex.EncodeToString(hash[:])

	return block
}
