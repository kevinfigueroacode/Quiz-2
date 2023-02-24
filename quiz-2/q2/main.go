// communication between users
// channels
package main

import (
	"fmt"
	"time"
)

func sender(c chan<- string) {
	// sends back the value whcih is the message to the receiver
	// step1: put data in the channel
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Message %d", i+1)
		c <- msg //place the massage in the the channel
		time.Sleep(time.Second)
	}
	close(c)
}
func receiver(c <-chan string) {
	for msg := range c {
		fmt.Println("Received:", msg)
	}
}

func main() {
	// create a channel for the communication
	c := make(chan string)

	// step2: get data from the channel(read)
	go sender(c)
	go receiver(c)
	// waits for the goroutine to complete
	time.Sleep(5 * time.Second)

}
