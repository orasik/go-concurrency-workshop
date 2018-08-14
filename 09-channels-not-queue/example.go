package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for m := range ch {
				fmt.Printf("routine %d received %d\n", i, m)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		ch <- i
	}
	time.Sleep(time.Second)
}
