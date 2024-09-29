package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%[1]T: %[1]s\n", e)
		}
	}()
	var a [2]int
	n := 2
	println(a[n])
}
