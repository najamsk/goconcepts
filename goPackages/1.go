package main

import (
	"fmt"
	"pkgs/mystore/films"
	actors "pkgs/mystore/users"
)

func main() {
	fmt.Println(actors.GetActors())
	fmt.Println(films.GetFilms())
}
