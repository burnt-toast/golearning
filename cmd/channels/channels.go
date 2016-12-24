package main

import "fmt"

func main() {
	messages := make(chan string)
	bufmessages := make(chan string, 2)

	go func() { messages <- "ping" }()

	bufmessages <- "buffered"
	bufmessages <- "channel"

	msg := <-messages
	fmt.Println(msg)

	fmt.Println(<-bufmessages)
	fmt.Println(<-bufmessages)
}
