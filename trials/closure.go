package main

import (
	"fmt"
	"sync"
)

//XXX: When iterating using i, v := range whatever, the i and v variables are defined only once for the scope, and then their values are overwritten. Since the goroutines reference the same variable all the time, but they are probably executed after the for loop is finished, they all access the same last value of it: 1.
func main() {
	fmt.Println("hello world")
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		i := i
		go func() {
			fmt.Println("hello world", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
