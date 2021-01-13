package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/japiirainen/go-ms-1/graphql/generated"
	"github.com/japiirainen/go-ms-1/graphql/model"
)

func (r *queryResolver) Cat(ctx context.Context, id string) (*model.Cat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cats(ctx context.Context) ([]*model.Cat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Owner(ctx context.Context, id string) (*model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Owners(ctx context.Context) ([]*model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
