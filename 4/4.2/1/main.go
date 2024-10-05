package main

import "fmt"

type F struct {
	Name string
	Age  int
}

func main() {
	f := &F{
		Name: "John",
		Age:  20,
	}
	// fmt.Printf("%[1]v\n%[1]T", f)
	// fmt.Printf("%+v\n", f)
	fmt.Printf("%#v\n", f)
}
