package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"fmt"
	"log"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	user := ForContext(ctx)
	log.Printf("user read from ctx is : %#v \n", user)
	usr1 := &model.User{ID: "1", Name: "Najam Awan"}
	usr2 := &model.User{ID: "2", Name: "Naeem Hassan"}
	result := []*model.Todo{
		&model.Todo{
			ID:   "1",
			Text: "hello task 1",
			Done: false,
			User: usr1,
		},
		&model.Todo{
			ID:   "2",
			Text: "hello task 2",
			Done: true,
			User: usr2,
		},
	}
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
