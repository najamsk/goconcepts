package main

import "fmt"

func test(canvas []string) {
	for _, v := range canvas {
		fmt.Println(v)
	}
	rows := len(canvas)
	cols := len(canvas[0])

	sr := rows / 2
	sc := cols / 2

	fmt.Println(sr)
	fmt.Println(sc)
}

func main() {
	fmt.Println("hi")
	test([]string{"aabba", "aabba", "aaacb"})
}
