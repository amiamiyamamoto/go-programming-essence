package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b [512]byte
	n, err := f.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))
	defertest()

	closure()
}
func defertest() {
	defer fmt.Println("6")
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

func doSomething(dir string) error {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	f, err := os.Create(filepath.Join(dir, "data2.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	//ファイルを使った処理
	return nil
	// Windowsではファイルハンドルが開かれた状態でディレクトリを削除できない

}

func closure() {
	var n = 1
	defer func() {
		fmt.Println(n)
	}()
	defer fmt.Println(n)
	n = 2
}
