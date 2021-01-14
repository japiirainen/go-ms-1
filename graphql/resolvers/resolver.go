package resolvers

import "github.com/japiirainen/go-ms-1/graphql/generated"

//go:generate go run github.com/99designs/gqlgen

// Resolver is main resolver
type Resolver struct{}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
