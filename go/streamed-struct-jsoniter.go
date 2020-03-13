package main

import (
	"bufio"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type CityLots struct {
	Type     string `json:"type"`
	Features []Lot  `json:"data"`
}

type Lot struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type Properties struct {
	MapBlkLot string `json:"MAPBLKLOT"`
	BlkLot    string `json:"BLKLOT"`
	BlockNum  string `json:"BLOCK_NUM"`
	LotNum    string `json:"LOT_NUM"`
	FromSt    string `json:"FROM_ST"`
	ToSt      string `json:"TO_ST"`
	Street    string `json:"STREET"`
	StType    string `json:"ST_TYPE"`
	OddEvent  string `json:"ODD_EVEN"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates []interface{} `json:"coordinates"`
}

func main() {
	dec := json.NewDecoder(bufio.NewReader(os.Stdin))
	for dec.More() {
		var result CityLots
		err := dec.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s\n", result)
	}
}
