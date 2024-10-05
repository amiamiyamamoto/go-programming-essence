package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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
	f, err := os.Open("hato.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data Data
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T", data)
	fmt.Println(data)
}
