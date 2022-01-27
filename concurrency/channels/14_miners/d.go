package main

import (
	"fmt"
)

// XXX: main goroutine doing all the work: calling each function and their output
func main() {
	myChan := make(chan string)
	go func() {
		myChan <- "Message!"
	}()

	for {
		select {
		case msg := <-myChan:
			fmt.Println(msg)
			return
		default:
			fmt.Println("No Msg")
		}
	}
	// <-done
}
