package main

import (
	"fmt"
	"time"
)

// XXX: main goroutine doing all the work: calling each function and their output
func main() {
	theMine := [5]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			oreChan <- item //send
		}
	}(theMine)

	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan //receive
			fmt.Println("Miner: Received " + foundOre + " from finder")
		}
	}()

	bufferedChan := make(chan string, 3)
	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
	}()
	<-time.After(time.Second * 1)
	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)
	}()

	<-time.After(time.Second * 5) // Again, ignore this for now

}

func finder1(mine []string) []string {
	output := []string{}
	for _, v := range mine {
		if v == "ore" {
			output = append(output, v)
			fmt.Println("finder1:found ore")
		}
	}
	return output
}
func finder2(mine []string) []string {
	output := []string{}
	for _, v := range mine {
		if v == "ore" {
			output = append(output, v)
			fmt.Println("finder2:found ore")
		}
	}
	return output
}

func miner(mine []string) []string {
	output := []string{}
	for _ = range mine {
		output = append(output, "minedOre")
		// output[k] = "minedOre"
	}
	return output
}

func smelter(mine []string) []string {
	output := []string{}
	for _ = range mine {
		output = append(output, "smeltedOre")
		// output[k] = "smeltedOre"
	}
	return output
}
