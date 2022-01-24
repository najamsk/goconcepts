package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	started := time.Now()
	menu := []string{"biryani", "qorma", "siri paye", "fish", "mutton"}
	var wg sync.WaitGroup
	wg.Add(len(menu))
	for _, f := range menu {
		go func(f string) {
			cook(f)
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Printf("done in %v\n", time.Since(started))

}

func cook(food string) {
	fmt.Printf("cooking %s... \n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s... \n", food)
	fmt.Println("")
}
