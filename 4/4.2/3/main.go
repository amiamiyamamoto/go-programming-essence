//go:generate stringer -type Fruit main.go
package main

import "fmt"

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func main() {
	fmt.Printf("%v", Orange.String())
	fmt.Println(Apple.String())
}
