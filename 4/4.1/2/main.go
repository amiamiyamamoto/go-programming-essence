package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%[1]T: %[1]s\n", e)
		}
	}()
	panic("my error")
}
