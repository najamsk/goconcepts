package main

import "fmt"

func ZerooInSum(val int) int {
	if val == 0 {
		return val
	}
	return ZerooInSum(val-1) + val
}

func main() {
	fmt.Println(ZerooInSum(3))
}
