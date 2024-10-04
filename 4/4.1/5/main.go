package main

var name = "John"

// グローバル変数初期化のタイミングで呼ばれる
func init() {
	println("Hi! " + name)
}
func main() {
	println("Hello! " + name)
}
