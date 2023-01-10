package main

import (
	"fmt"
	"log"
	"time"
)

func greet(s string, ff func(string) string) string {
	return ff(s)
}

func closuer() func(string) string {
	t := "dr."
	return func(s string) string {
		return t + s
	}
}

func StartTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Since(t)
		log.Println(name, "took", d)
	}
}

func funcyFunction() {
	stop := StartTimer("funcyFunc")
	defer stop()
	time.Sleep(1 * time.Second)
}

func main() {
	dr := closuer()
	vv := greet("najam awan", dr)
	fmt.Println(vv)
	funcyFunction()
}
