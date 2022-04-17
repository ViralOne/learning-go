package main

import (
	"fmt"
	"time"
)

func main() {

	// we can use go when the app is running in a single thread and is waiting for a specific
	// event to happen, like a message from a server or a file to be read
	// go is used to run a function in a separate thread of execution

	go task("Milu", "devops")
	go task("Adi", "developer")
	go task("Mojo", "qa")
	go task("Ivan", "developer")

	task("main", "testing")
}

// create a function using multithread method
func task(name string, job string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, " - ", name, "is", job)
		time.Sleep(time.Second)
	}
}
