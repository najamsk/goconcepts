package main

import (
	"context"
	"fmt"
	"log"

	graphql "github.com/hasura/go-graphql-client"
)

func main() {
	client := graphql.NewClient("https://rickandmortyapi.com/graphql", nil)
	var query struct {
		Characters struct {
			Info struct {
				Count int
			}
			Results []struct {
				Id   string
				Name string
			}
		}
	}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
		log.Println(err)
	}
	fmt.Println(query.Characters.Info.Count)
	fmt.Println(query.Characters.Results[0].Id)
	fmt.Println(query.Characters.Results[0].Name)
	fmt.Println("___________________________________")
	fmt.Println(query.Characters.Results[1].Id)
	fmt.Println(query.Characters.Results[1].Name)
}
