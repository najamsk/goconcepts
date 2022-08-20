package main

import "fmt"

type Age int
type Team struct {
	Name string
	Size int
}

type Teams []Team

func GetTeams() Teams {
	teams := Teams{
		Team{Name: "Paki", Size: 11},
		Team{Name: "Romani", Size: 4},
	}
	return teams
}

func main() {
	var a Age
	a = 22
	fmt.Println("hello", a)
	fmt.Printf("teams = %#v \n", GetTeams())
}
