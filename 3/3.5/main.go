package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	message := "hi"
	go sendMessage(message, &wg) // 引数はキャプチャされるため、hiが出力される
	go func() {
		sendMessage(message, &wg)
	}()
	message = "ho"

	wg.Wait()
}

func sendMessage(msg string, wg *sync.WaitGroup) {
	fmt.Println(msg)
	wg.Done()
}
