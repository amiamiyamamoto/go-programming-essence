package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var content = `
{
	"species": "鳩",
	"description": "岩に止まるのが好き",
	"dimensions": {
		"height": 24,
		"width": 10
	}
}
`

type Dimensions struct {
	Width  int `json: "width"`
	Height int `json: "height"`
}

type Data struct {
	Species     string     `json: "species"`
	Description string     `json: "description"`
	Dimensions  Dimensions `json: "dimensions"`
}

func main() {
	var data Data
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
