package main

import (
	"blockchain/blockchain"
	"blockchain/miners"
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {

	chain, err := blockchain.LoadBlockchain()
	if err != nil || chain == nil {
		chain = blockchain.CreateBlockchain()
	}

	fmt.Printf("\n========== | BATTERY BLOCKCHAIN | ==========\n")

menu:
	for {
		fmt.Printf("  1. Add battery record  \n  2. Display all blocks  \n  3. Search battery by ID\n  4. View Battery history\n  0. Exit				 \n  Choose an option: ")

		var answer int

		_, err = fmt.Scan(&answer)
		if err != nil {
			log.Fatal(errors.New("Incorrect input"))
		}

		switch answer {

		case 0:
			blockchain.SaveBlockchain(chain)
			break menu
		case 1:
			var batteryID, status string
			var voltage, temperature, stateOfCharge float64

			fmt.Print("\nEnter Battery ID: ")
			fmt.Scan(&batteryID)
			fmt.Print("\nEnter status: ")
			fmt.Scan(&status)
			fmt.Print("\nEnter voltage: ")
			fmt.Scan(&voltage)
			fmt.Print("\nEnter temperature: ")
			fmt.Scan(&temperature)
			fmt.Print("\nEnter state of charge: \n\n")
			fmt.Scan(&stateOfCharge)

			data := blockchain.Data{
				BatteryID:      batteryID,
				Status:         status,
				Voltage:        voltage,
				Temperature:    temperature,
				StateOfCharge:  stateOfCharge,
				ProductionTime: time.Now(),
			}

			miners.MineBlock(chain, data)

		case 2:
			for _, block := range chain.Blocks {
				fmt.Printf("\n")
				fmt.Printf("%+v\n\n", block)
			}
		case 3:
			var findBatteryID string
			fmt.Printf("Enter Battery ID: ")
			fmt.Scan(&findBatteryID)

			battery, exists := chain.FindBattery(findBatteryID)

			if exists {
				fmt.Printf("\n\nBattery ID: %s", battery.Data.BatteryID)
				fmt.Printf("\nStatus: %s", battery.Data.Status)
				fmt.Printf("\nVoltage: %.2f", battery.Data.Voltage)
				fmt.Printf("\nTemperature: %.2f", battery.Data.Temperature)
				fmt.Printf("\nState Of Health: %.2f\n\n", battery.Data.StateOfCharge)
			} else {
				fmt.Printf("\nNo battery found\n\n")
			}

		case 4:
			var findBatteryHistory string
			fmt.Printf("\nEnter Battery ID: ")
			fmt.Scan(&findBatteryHistory)

			battery := chain.FindBatteryHistory(findBatteryHistory)

			if len(battery) == 0 {
				fmt.Printf("\nNo battery history found\n\n")
			} else {
				for _, block := range battery {
					fmt.Printf("\n\nBattery ID: %s", block.Data.BatteryID)
					fmt.Printf("\nStatus: %s", block.Data.Status)
					fmt.Printf("\nProduction Time: %s\n\n", block.Data.ProductionTime.Format(time.RFC3339))
				}
			}

		default:
			fmt.Printf("\n\n___ _______ ___\n")
			fmt.Printf("Incorrect input")
			fmt.Printf("\n___ _______ ___\n\n\n")
		}
	}
}
