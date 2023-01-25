package main

import "fmt"

type bird struct {
	name string
}

func (b *bird) speak() {
	fmt.Printf("%v says hello \n", b.name)
}

type parrot struct {
	bird
}

func main() {
	p := parrot{bird: bird{name: "totu"}}
	// p.name = "kameena"
	fmt.Println(p.name)
	p.speak()

}
