package main

import "fmt"

type person interface {
	greet(string)
}

type emp interface {
	say(string)
	person
}

func test(e emp) {
	if e == nil {
		fmt.Println("you sent nil")
	}
}

type employee struct {
}

func (e employee) greet(g string) {
	fmt.Println(g)
}
func (e employee) say(s string) {
	fmt.Println(s)
}

func greeter(e emp) {
	e.greet("greetings mr.")
	e.say("how are you?")
}

func main() {
	test(nil)
	e := employee{}
	greeter(e)
}
