package resolvers

import (
	"context"
	"fmt"

	"github.com/japiirainen/go-ms-1/graphql/model"
)

func (r *queryResolver) Cat(ctx context.Context, id string) (*model.Cat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cats(ctx context.Context) ([]model.Cat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCat(ctx context.Context, input model.NewCat) (*model.Cat, error) {
	panic(fmt.Errorf("not implemented"))
}
