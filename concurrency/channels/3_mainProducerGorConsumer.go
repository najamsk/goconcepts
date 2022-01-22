package main

import (
	"fmt"
	"time"
)

// XXX: Main writing to ch and gor listening to ch
func main() {
	ch := make(chan string)
	fmt.Println(time.Now(), "main func invoked")
	go func() {
		//read form channcel
		time.Sleep(2 * time.Second)
		v := <-ch
		fmt.Println(time.Now(), "message recevied=", v)
	}()
	//write to channel
	ch <- "hello"
	fmt.Println("main ends")
}
