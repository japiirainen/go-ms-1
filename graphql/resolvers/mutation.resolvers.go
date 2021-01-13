package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/japiirainen/go-ms-1/graphql/generated"
	"github.com/japiirainen/go-ms-1/graphql/model"
)

func (r *mutationResolver) CreateOwner(ctx context.Context, input model.NewOwner) (*model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCat(ctx context.Context, input model.NewCat) (*model.Cat, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
