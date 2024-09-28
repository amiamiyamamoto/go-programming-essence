package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	var a [2]int
	n := 2
	println(a[n])
}
