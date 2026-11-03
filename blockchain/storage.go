package blockchain

import (
	"encoding/json"
	"errors"
	"os"
)

func SaveBlockchain(bc *Blockchain) error {

	Data, err := json.MarshalIndent(bc, "", " ")

	if err != nil {
		return errors.New("Couldn't save Blockchain in memory.")
	}

	err = os.WriteFile("blockchain.json", Data, 0644)
	if err != nil {
		return errors.New("Cant write on this file.")
	}

	return nil
}

func LoadBlockchain() (*Blockchain, error) {

	Data, err := os.ReadFile("blockchain.json")
	if err != nil {
		return nil, err
	}

	var bc Blockchain

	err = json.Unmarshal(Data, &bc)
	if err != nil {
		return nil, err
	}

	if bc.Difficulty < 1 || bc.Difficulty > 6 {
		return &bc, nil
	}

	return &bc, nil
}
