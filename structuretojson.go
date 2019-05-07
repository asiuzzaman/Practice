package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type FruitBasket struct {
	Name    string
	Fruit   []string
	Id      int64 `json:"ref"`
	private string
	Created time.Time
}

func main() {
	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"apple", "Banana", "Orange"},
		Id:      999,
		private: "second-rate",
		Created: time.Now(),
	}
	var jsonData []byte
	jsonData, err := json.MarshalIndent(basket, "", "   ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}
