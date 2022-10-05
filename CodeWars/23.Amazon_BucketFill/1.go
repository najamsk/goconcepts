package main

import "fmt"

func test(canvas []string) {
	r := len(canvas)
	c := len(canvas[0])
	visited := make([][]bool, r)
	mat := make([][]string, r)

	for k := range visited {
		visited[k] = make([]bool, c)
	}
	for k := range canvas {
		mat[k] = make([]string, c)
		for i, v := range canvas[k] {
			mat[k][i] = string(v)
		}
	}

	fmt.Println(r)
	fmt.Println(c)
	fmt.Printf("%v \n", mat[0])
	fmt.Printf("%v \n", mat[1])
	fmt.Printf("%v \n", mat[2])

}

func main() {
	fmt.Println("hi")
	test([]string{"aabba", "aabba", "aaacb"})
}
