package graph

//go:generate go run github.com/arsmn/fastgql

import "github.com/norbux/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
	canarios []*model.Canario
	//tipo []*model.Tipo
	//...
}
