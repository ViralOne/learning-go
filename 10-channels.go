package main

import (
	"fmt"
	"time"
)

// channels are pipelines that connect concurrent gorutines

func main() {
	//channel()
	//buffering()

	//for synchronizing gorutines
	done := make(chan bool, 2)
	go sync(done)
	<-done
}

func channel() {
	// create a new channel
	message := make(chan string)

	// send a message to the channel
	go func() {
		message <- "Hello World"
	}()

	// receive a message from the channel
	msg := <-message
	fmt.Println(msg)
}

func buffering() {
	// create a new channel with a buffer of 3
	message := make(chan string, 3)

	// send a few messages to the channel
	message <- "Hello World"
	message <- "Bufferd message"

	// receive messages from the channel
	fmt.Println(<-message)
	fmt.Println(<-message)
}

// synchronizing gorutines
// using done channel to notify another gorutine
// that the first one is done
func sync(done chan bool) {
	fmt.Println("starting...\n")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
