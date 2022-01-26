package main

import (
	"fmt"
	"time"
)

// XXX: main goroutine doing all the work: calling each function and their output
func main() {
	myChan := make(chan string)
	go func() {
		myChan <- "Message!"
	}()

	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}
	<-time.After(time.Second * 2)
	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}

}
