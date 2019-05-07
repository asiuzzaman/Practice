package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type FruitBasket struct {
	Name  string
	Fruit []string
	Id    int64 `json:"ref"`

	Created time.Time
}

func main() {

	jsonData := []byte(`
	{
	 "Name":"Standard",
	 "Fruit":[
		 "Apple",
		 "Banana",
		 "Orange"
	 ],
	 "ref":999,
	 "Create":"2019-05-07"
	}
	`)
	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket) // stores json data to a basket
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket.Name)
	fmt.Println(basket.Fruit)
	fmt.Println(basket.Id)
	fmt.Println(basket.Created)
}
