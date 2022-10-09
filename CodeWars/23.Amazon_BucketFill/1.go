package main

import "fmt"

func traverseThroughRiver(mat [][]byte, char byte, row int, col int, visited [][]bool, currentRiverSize *int) {
	//if current point (x,y) already visited
	if visited[row][col] {
		return
	}
	//if current point (x,y) is 0 -> land
	if mat[row][col] != char {
		return
	}
	//current point is part of river
	visited[row][col] = true
	*currentRiverSize++
	//get current point neighbours
	neighbours := getNeighbours(mat, row, col)

	for _, neighbour := range neighbours {
		traverseThroughRiver(mat, char, neighbour[0], neighbour[1], visited, currentRiverSize)
	}
}

func getNeighbours(mat [][]byte, row int, col int) [][]int {
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

func makeMat(mat []string) int {
	sizes := []int{} //result array that will contain found river sizes
	rows := len(mat)
	cols := len(mat[0])
	can := make([][]byte, rows)

	visited := make([][]bool, rows)
	for k := range visited {
		visited[k] = make([]bool, cols)
	}
	//create mat
	for k, v := range mat {
		fmt.Println(v)
		can[k] = make([]byte, cols)
		for i, _ := range v {
			can[k][i] = byte(v[i])
		}
	}

	//traverse map
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			// fmt.Printf("%d:%v ", mat[i][j], visited[i][j], sizes)
			currentRiverSize := 0
			traverseThroughRiver(can, can[i][j], i, j, visited, &currentRiverSize)
			if currentRiverSize > 0 {
				sizes = append(sizes, currentRiverSize)
			}
		}
		// fmt.Println("")
	}
	return len(sizes)

}
func main() {
	// traverse([]string{"aabba", "aabba", "aaacb"})
	fmt.Println(makeMat([]string{"aabba", "aabba", "aaacb"}))
}
