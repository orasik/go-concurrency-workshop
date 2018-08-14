package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var counter int

	set := func(i int) {
		mu.Lock()
		defer mu.Unlock()
		counter = i
	}

	get := func() int {
		mu.Lock()
		defer mu.Unlock()
		return counter
	}

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			i := get()
			fmt.Printf("counter: %d\n", i)
		}
	}()

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("setting ", i)
			set(i)
		}(i)
		time.Sleep(750 * time.Millisecond)
	}

	wg.Wait()
}m