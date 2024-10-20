package main

import "fmt"

type I int

func (i *I) Add(n int) I {
	return *i + I(n)
}

func main() {
	var n I = 0

	n = n.Add(1)
	n = n.Add(2)
	fmt.Println(n)

	add := n.Add
	fmt.Println(add(3))
	fmt.Printf("%T\n", (*I).Add)
}
