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
	defer close(ch) //終わったら閉じる（５）

	//HTTPサーバからのダウンロード
	for _, u := range urls {
		fmt.Println("\nDownloading", u)
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV: ", err)
			continue
		}
		//b, err := ioutil.ReadAll(resp.Body)//ioutil.ReadAll is deprecated: As of Go 1.16, this function simply calls [io.ReadAll].
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content: ", err)
			continue
		}
		resp.Body.Close()
		ch <- b // main関数にコンテンツを送信（3）
	}
}

func insertRecords(records []string) {
	// レコードの登録
	fmt.Println(records)
}

func main() {
	urls := []string{
		"http://www.mi.u-tokyo.ac.jp/consortium2/data/blsfood_new.csv",
		"http://www.mi.u-tokyo.ac.jp/consortium2/data/Lynx-new.csv",
		// "http://my-server.com/data03.csv",
	}
	//バイト列を転送するためのchannelを作成（1）
	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch) //(2)

	//goroutineからコンテンツを受け取る（4）
	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			// レコードの登録
			insertRecords(records)
		}
	}
	wg.Wait()
}
