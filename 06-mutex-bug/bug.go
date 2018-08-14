package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	mu := &sync.Mutex{}

	var count int

	go func() {
		for i := 0; i < 5; i++ {
			mu.Lock()
			defer mu.Unlock()
			count = i
			fmt.Println(i)
		}
		wg.Done()
	}()

	wg.Wait()
}
