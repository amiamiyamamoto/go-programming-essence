package main

import "fmt"

func main() {
	hiho()
	hihoClosure()

}

func hiho() {
	message := "hi"
	go sendMessage(message) // 引数はキャプチャされるため、hiが出力される
	message = "ho"
}

func hihoClosure() {
	message := "hi"
	go func() {
		sendMessage(message)
	}()
	message = "ho"
}
func sendMessage(msg string) {
	fmt.Println(msg)
}
