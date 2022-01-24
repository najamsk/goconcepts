package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	foods := []string{"biryani", "qorma", "siri paye", "fish", "mutton"}
	for _, f := range foods {
		cook(f)
	}
	fmt.Printf("done in %v\n", time.Since(started))

}

func cook(food string) {
	fmt.Printf("cooking %s... \n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s... \n", food)
	fmt.Println("")
}
