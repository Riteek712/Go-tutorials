package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil" is depricated so just use os.
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}
type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)
}
func loadJSON[T contactInfo | purchaseInfo](filepath string) []T {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	var loaded []T
	err = json.Unmarshal(data, &loaded)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	return loaded
}
