package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var sleep time.Duration
	flag.DurationVar(&sleep, "sleep", time.Second, "sleep time")
	flag.Parse()
	time.Sleep(sleep)
	fmt.Println("o-i")
}
