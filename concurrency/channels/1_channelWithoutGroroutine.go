package main

import "fmt"

//XXX: this program panics because no goroutines are using channel outside of main.

//error will be all goroutines are asleep

func main() {
	fmt.Println("hello world")
	c := make(chan int)
	c <- 10
	v := <-c
	fmt.Println("received", v)
}
