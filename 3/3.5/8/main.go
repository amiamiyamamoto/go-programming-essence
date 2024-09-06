package main

import (
	"fmt"
	"math/rand"
	"time"
)

var zunDoko int8

const (
	DOKO int8 = iota
	ZUN
)

func main() {
	// var mu sync.Mutex

	znch := make(chan int8)
	dkch := make(chan int8)

	go makeZun(znch)

	go makeDoko(dkch)

L:
	for {
		select {
		case zn := <-znch:
			fmt.Println("ZUN")
			zunDoko = zunDoko << 1
			zunDoko = zunDoko | zn

		case <-dkch:
			fmt.Println("DOKO")
			zunDoko = zunDoko << 1
			if (zunDoko & 0b00001111) == 14 {
				fmt.Println("KIYOSHI!!")
				break L
			}

		}
	}

	fmt.Println("end")
}

func makeZun(znch chan int8) {
	for {
		if randomBinary() == 1 {
			znch <- ZUN
		}
	}
}

func makeDoko(dkch chan int8) {
	for {
		if randomBinary() == 1 {
			dkch <- DOKO
		}
	}
}

func randomBinary() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(2)
}
