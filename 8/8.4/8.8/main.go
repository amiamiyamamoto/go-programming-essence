package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func hevyFunc(wg *sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 1000000; i++ {
		s = append(s, "magicval pandas")
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		var wg sync.WaitGroup
		wg.Add(1)
		go hevyFunc(&wg)
		wg.Wait()
	}
}
