package main

import "fmt"

// XXX: main goroutine doing all the work: calling each function and their output
func main() {
	mine := []string{"rock", "ore", "ore", "rock", "ore"}
	foundOre := finder(mine)
	fmt.Println("From Finder:", foundOre)
	minedOre := miner(foundOre)
	fmt.Println("From Miner:", minedOre)
	smeltedOre := smelter(minedOre)
	fmt.Println("From Smelter:", smeltedOre)
}
func finder(mine []string) []string {
	output := []string{}
	for _, v := range mine {
		if v == "ore" {
			output = append(output, v)
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
