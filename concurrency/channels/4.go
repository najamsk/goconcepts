package main

import (
	"fmt"
	"time"
)

//XXX:Bufferd Channel

func main() {
	fmt.Println("main func invoked")
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			//XXX: there are cases when this message will not be printed
			fmt.Println(time.Now(), i, "sent")
		}
		//XXX: there are cases when this message will not be printed
		fmt.Println(time.Now(), "all completed")
	}()
	time.Sleep(2 * time.Second)

	fmt.Println(time.Now(), "waiting for messages")
	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "closing")
}
