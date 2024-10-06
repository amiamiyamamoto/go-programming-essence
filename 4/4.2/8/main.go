package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))
}

// var s = "2022/12/25 07:42:38"
// d, err := time.Parse("2006/01/02 15:04:05", s)
