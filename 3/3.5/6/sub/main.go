package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, u string, ch chan []byte) {
	defer wg.Done()

	// HTTPサーバからのダウンロード
	resp, err := http.Get(u)
	if err != nil {
		log.Println("cannot download CSV: ", err)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("cannot read content: ", err)
		return
	}
	ch <- b
}

func insertRecords(records []string) {
	fmt.Println(records)
}

func main() {
	urls := []string{
		"http://www.mi.u-tokyo.ac.jp/consortium2/data/Lynx-new.csv",
		"http://www.mi.u-tokyo.ac.jp/consortium2/data/blsfood_new.csv",
	}

	ch := make(chan []byte)
	var wg sync.WaitGroup

	// 各URLのダウンロードを並行して行う
	for _, url := range urls {
		wg.Add(1)
		go downloadCSV(&wg, url, ch)
	}

	// ダウンロード完了を待ち、チャネルを閉じる
	go func() {
		wg.Wait()
		close(ch)
	}()

	// チャネルからデータを受信して処理
	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			insertRecords(records)
		}
	}
}
