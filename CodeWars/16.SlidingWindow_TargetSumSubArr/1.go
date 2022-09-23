package main

import (
	"fmt"
)

func maxSubarraySum(array []int, n int) int {
	//Write a check for the edge case if the array's length is smaller than the subarray's length. If it is then we return 0.
	if n > len(array) {
		return 0
	}
	maxSum := 0  //Assign a variable that will hold the highest sum
	tempSum := 0 //Assign a variable that will hold the sum of the current subarray we're looking at.

	//Create a loop that will Loop through 1 time to get the sum of the first set of numbers
	for i := 0; i < n; i++ {
		//Set maxSum to equal the sum of our first loop
		maxSum += array[i]
	}
	//Now that we have a max sum, assign temp sum to max sum
	tempSum = maxSum

	//Loop through the length of the whole array starting from the index of the next number after the first set
	for i := n; i < len(array); i++ {
		//assigning new tempsum, we get the current tempSum, move down the number of indexes == to n, subtract that number, then add the number of the current loops index
		fmt.Println(tempSum, array[i-n], array[i]) //Visual representation when ran to see the values so you can see whats going on.
		tempSum = tempSum - array[i-n] + array[i]
		//see which is larger the current maxSum or the tempSum we're holding, if tempSum is larger maxSum is now the value of TempSum
		maxSum = Max(maxSum, tempSum)
	}
	return maxSum
}

// helper function to get the higher of two integers
func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
func main() {
	fmt.Println("hello")
	items := []int{-1, -1, -2, 4, 2, 3, 5, 1}
	fmt.Println(maxSubarraySum(items, 4))
	// items := []int{1, 2, 4, 1, 1}
}
