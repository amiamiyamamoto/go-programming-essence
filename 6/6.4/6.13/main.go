package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadJSON(u string) {
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		// i := i//今は不要
		go func() {
			defer wg.Done()
			u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
			downloadJSON(u)
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Since(before))
}
