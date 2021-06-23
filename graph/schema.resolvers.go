package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/norbux/gqlgen-todos/db"
	"github.com/norbux/gqlgen-todos/graph/generated"
	"github.com/norbux/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   strconv.Itoa(rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}

	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateCanario(ctx context.Context, input model.NewCanario) (*model.Canario, error) {
	canario := &model.Canario{
		ID: strconv.Itoa(rand.Int()),
		Name: input.Name,
		Email: input.Email,
	}

	// r.canarios = append(r.canarios, canario)
	// return canario, nil

	return db.CreateCanario(canario), nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Canarios(ctx context.Context) ([]*model.Canario, error) {
	//return r.canarios, nil
	return db.GetCanarios(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
