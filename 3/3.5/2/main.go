package main

import "time"

func sendMessage(msg string) {
	println(msg)
}

func main() {
	message := "hi"
	go func() {
		sendMessage(message)
	}()
	message = "ho"

	time.Sleep(time.Second)
	println(message)
	time.Sleep(time.Second)
	// データ競合（race condition）が起きていないかを-raceオプションで確認できる
	// go run -race main.go
	// go build -race main.go

}
