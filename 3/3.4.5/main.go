package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b [512]byte
	n, err := f.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))
}
