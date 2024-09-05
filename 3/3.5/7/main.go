package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Hello"
	}()

	go func() {
		ch2 <- "World"
	}()

L:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println(v1)
			break L
		case v2 := <-ch2:
			fmt.Println(v2)
		}
	}
}
