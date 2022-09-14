package main

import "fmt"

func MaxSum2(values []int) int {
	c := values[0]
	m := values[0]

	for i := 1; i < len(values); i++ {
		t := c + values[i]
		c = max(t, values[i])
		if c > m {
			m = c
		}
	}
	return m
}

func MaxSum(values []int) int {
	maxSum := values[0]
	curSum := values[0]
	m := []int{}

	for i := 1; i < len(values); i++ {
		curSum = curSum + values[i]
		if curSum > maxSum {
			maxSum = curSum
			m = append(m, values[i])
		}
	}
	total := 0
	for _, v := range m {
		total = total + v
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("hello")
	fmt.Println(MaxSum([]int{-2, 2, 5, -11, 6}))
	fmt.Println(MaxSum2([]int{-2, 2, 5, -11, 6}))
}
