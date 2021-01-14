package resolvers

import (
	"context"
	"fmt"

	"github.com/japiirainen/go-ms-1/graphql/model"
)

var owner *model.NewOwner

func (r *queryResolver) Owner(ctx context.Context, id string) (*model.Owner, error) {
	return &model.Owner{ID: 1, Name: "Jeppe Meikäläinen", Username: "greatUsername"}, nil
}

func (r *queryResolver) Owners(ctx context.Context) ([]model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateOwner(ctx context.Context, input model.NewOwner) (*model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}
