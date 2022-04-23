package main

import "fmt"

// channels are pipelines that connect concurrent gorutines

func main() {
	// create a new channel
	ping := make(chan string, 1)
	pong := make(chan string, 1)

	// send a message to the channel
	send(ping, "ping")

	// receive a message from the channel
	recive(ping, pong)
	fmt.Println(<-pong)
}

// function that only allows sending values
func send(ping chan<- string, msg string) {
	ping <- msg
}

//function that only allows receiving values
func recive(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}
