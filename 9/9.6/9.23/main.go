package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name string
	var max int

	flag.IntVar(&max, "max", 255, "max value")
	flag.StringVar(&name, "name", "something", "my name")
	flag.Parse()

	for _, arg := range flag.Args() {
		fmt.Println(arg)
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	println(max, name)
}
