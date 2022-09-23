package main

import "fmt"

type Err string

func (e Err) Error() string {
	return string(e)
}

const ErrNotFound = Err("My man, Item Not found")

type Emp struct {
	name string
}

func main() {
	fmt.Println(ErrNotFound)
}
