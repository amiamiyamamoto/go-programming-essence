package main

import (
	"fmt"

	runewodth "github.com/mattn/go-runewidth"
)

func main() {
	fmt.Println(runewodth.StringWidth("こんにちは"))
}
