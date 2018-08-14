package main

import "fmt"

func main() {
	fmt.Println(1)

	// make a new channel of type string
	messages := make(chan string)

	fmt.Println(2)
	// this goroutine launches immediately
	go func() {
		fmt.Println(3)

		// This line blocks until the channel is read from
		messages <- "hello!"
		fmt.Println(4)
	}()

	fmt.Println(5)

	// this line blocks until someone writes to the channel
	msg := <-messages
	fmt.Println(6)

	fmt.Println(msg)
}