package main

import "fmt"

func process(messages chan string, quit chan struct{}) {
	// Loop through your messages
	for m := range messages {
		// print the message with a for loop using range
	}
}

func main() {
	// declare the messages channel of type string and capacity of 5

	// declare a signal channel

	// launch process in a goroutine

	// declare 5 fruits in a []string
	for _, f := range fruits {
	}

	// loop through the fruits and send them to the messages channel

	// close the messages channel

	// wait for everything to finish

	fmt.Println("done")
}