package main

import "fmt"

func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func main() {

	ch := make(chan string)
	go server(ch)

	// var s string
	// s = <-ch
	// fmt.Println(s)
	// s = <-ch
	// fmt.Println(s)
	// s = <-ch
	// fmt.Println(s)

	for s := range ch {
		fmt.Println(s)
	}

}
