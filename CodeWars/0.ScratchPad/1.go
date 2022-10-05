package main

import (
	"fmt"
)

type err string

func (e err) Error() string {
	return string(e)
}

const ErrNotFound = err("My man, Item Not found")
const ErrMyManNot = err("My man, wait up king kong you are not my man.")

type Emp struct {
	name string
}

func main() {
	s := "hello my friend"
	for _, v := range s {
		fmt.Println(string(v))
	}
	fmt.Println(ErrMyManNot)
}
