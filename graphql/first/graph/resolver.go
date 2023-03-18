package graph

import "context"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type User struct {
	Name    string
	IsAdmin bool
}

func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(UserCtxKey).(*User)
	return raw
}
