package main

import (
	"fmt"
)

const (
	DOKO uint8 = iota
	ZUN
)

func main() {
	znch := make(chan uint8)
	dkch := make(chan uint8)

	go makeZun(znch)
	go makeDoko(dkch)

	var zunDoko uint8

L:
	for {
		select {
		case zn := <-znch:
			zunDoko = (zunDoko << 1) + zn
			format := fmt.Sprintf("%08b", zunDoko)
			fmt.Println("ZUN ", format)
		case <-dkch:
			zunDoko = zunDoko << 1
			format := fmt.Sprintf("%08b", zunDoko)
			fmt.Println("DOKO", format)
			if (zunDoko & 0b1111) == 14 {
				fmt.Println("KIYOSHI!!")
				break L
			}

		}
	}

	fmt.Println("end")
}

func makeZun(znch chan uint8) {
	for {
		znch <- ZUN
	}
}

func makeDoko(dkch chan uint8) {
	for {
		dkch <- DOKO
	}
}
