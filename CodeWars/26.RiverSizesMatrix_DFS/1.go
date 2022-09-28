package main

import "fmt"

func traverse(mat [][]int) []int {
	sizes := []int{} //result array that will contain found river sizes
	r := len(mat)
	c := len(mat[0])

	visited := make([][]bool, r)
	for k := range visited {
		visited[k] = make([]bool, c)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			// fmt.Printf("%d:%v ", mat[i][j], visited[i][j], sizes)
			currentRiverSize := 0
			traverseThroughRiver(mat, i, j, visited, &currentRiverSize)
			if currentRiverSize > 0 {
				sizes = append(sizes, currentRiverSize)
			}
		}
		// fmt.Println("")
	}
	return sizes
}

func traverseThroughRiver(mat [][]int, row int, col int, visited [][]bool, currentRiverSize *int) {
	//if current point (x,y) already visited
	if visited[row][col] {
		return
	}
	//if current point (x,y) is 0 -> land
	if mat[row][col] == 0 {
		return
	}
	//current point is part of river
	visited[row][col] = true
	*currentRiverSize++
	//get current point neighbours
	neighbours := getNeighbours(mat, row, col)

	for _, neighbour := range neighbours {
		traverseThroughRiver(mat, neighbour[0], neighbour[1], visited, currentRiverSize)
	}
}

func getNeighbours(mat [][]int, row int, col int) [][]int {
	res := [][]int{}
	if row != 0 {
		res = append(res, []int{row - 1, col})
	}
	if col != 0 {
		res = append(res, []int{row, col - 1})
	}
	if row != len(mat)-1 {
		res = append(res, []int{row + 1, col})
	}
	if col != len(mat[row])-1 {
		res = append(res, []int{row, col + 1})
	}
	return res
}

func main() {
	mat := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}
	sizes := traverse(mat)
	fmt.Println(sizes)
}
