package main

import (
	"fmt"
	"sync"
)

func main() {
	incrementRaceCondition()
	n := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(n)
}

func incrementRaceCondition() {
	n := 0

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			n++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			n++
		}
	}()

	wg.Wait()
	fmt.Println(n) //2000を期待しているが、データ競合（race condition）が発生しているため、期待通りの値が出力されない
}
