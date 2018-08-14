package main

import "sync"

func count(prefix string, count int, wg *sync.WaitGroup) {
	// Defer the waitgroup Done

	for i := 0; i < count; i++ {
		// Print out the prefix and the loop variable `i`
	}
}

func main() {
	// Define the wait group

	// Increment wait group

	// Call count with a goroutine, first argument "first", second argument 50,
	// third argument the wait group

	// Call count with a goroutine, first argument "second", second argument 50,
	// third argument the wait group

	// Wait...
}