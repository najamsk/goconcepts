package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	menu := []string{"biryani", "qorma", "siri paye", "fish", "mutton"}
	results := make(chan bool)
	for _, f := range menu {
		go func(f string) {
			cook(f)
			results <- true
		}(f)
	}
	for i := 0; i < len(menu); i++ {
		<-results
	}
	fmt.Printf("done in %v\n", time.Since(started))

}

func cook(food string) {
	fmt.Printf("cooking %s... \n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s... \n", food)
	fmt.Println("")
}
