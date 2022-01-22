package main

import (
	"fmt"
	"time"
)

//XXX: here main is reading from channel that puts main to block. anon fun goroutine
// will be invoked and its writing to channel so once chan writing statement done it will be blocked
// control will be back to main goroutine which reads value from chan and print it.
func main() {
	ch := make(chan string)
	go func() {
		fmt.Println(time.Now(), "taking a nap")
		time.Sleep(2 * time.Second)
		ch <- "hello"
	}()
	fmt.Println(time.Now(), "waiting for a message")
	v := <-ch
	fmt.Println(time.Now(), "value is ", v)

}
