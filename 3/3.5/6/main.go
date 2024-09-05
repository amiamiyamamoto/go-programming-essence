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

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {

	defer wg.Done()
	defer close(ch) //(5)
	// HTTPサーバからのダウンロード
	for _, u := range urls {

		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV: ", err)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content: ", err)
			continue
		}
		resp.Body.Close()
		ch <- b //(3)
	}
}
func insertRedords(records []string) {
	fmt.Println(records)
}
func main() {
	urls := []string{
		"http://www.mi.u-tokyo.ac.jp/consortium2/data/blsfood_new.csv",
		"http://www.mi.u-tokyo.ac.jp/consortium2/data/Lynx-new.csv",
		// "http://my-server.com/data03.csv",
	}

	ch := make(chan []byte) //(1)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch) //(2)

	for b := range ch { //(4)
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			insertRedords(records)
		}
	}
	wg.Wait()
}
