package main

import "fmt"

func main() {
	fmt.Printf("%s %s", "hello")

	// go vet　はgoの標準静的解析ツール

	// $ go vet
	// # /go-programming-essence/3/3.8/1
	// # [/go-programming-essence/3/3.8/1]
	// ./main.go:6:2: fmt.Printf format %s reads arg #2, but call has 1 arg
}
