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
	dec := json.NewDecoder(f)
	for {
		err := dec.Decode(&data)
		if err != nil {
			break
		}
		fmt.Println(data, "\n\n")
	}
	// fmt.Printf("%T", data)
	fmt.Println(data, "aaaaaaaaaaaaaa")
}
