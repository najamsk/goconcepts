package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	fmt.Println("hello world")
	ch <- 2
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
