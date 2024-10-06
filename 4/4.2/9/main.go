package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	d, err := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d) //3s

	d, err = time.ParseDuration("4m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	d, err = time.ParseDuration("5h")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
	fmt.Println(d * 3) //四則演算もできる

	// Duration型はtimeパッケージで以下のように定義されている
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
}
