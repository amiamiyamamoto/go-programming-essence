package main

import (
	"fmt"
	"math/rand"
	"time"
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
			fmt.Println("ZUN ", fmt.Sprintf("%08b", zunDoko))
		case <-dkch:
			zunDoko = zunDoko << 1
			fmt.Println("DOKO", fmt.Sprintf("%08b", zunDoko))
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
		if rand01() == 1 {
			znch <- ZUN
		}
	}
}

func makeDoko(dkch chan uint8) {
	for {
		if rand01() == 1 {
			dkch <- DOKO
		}
	}
}

// 0 or 1の乱数を生成する
func rand01() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(2)
}
