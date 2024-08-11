package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i) //forが先に回ってしまい、iが10になってしまう場合がある。けどなってない
		}()
	}
	wg.Wait()
}
