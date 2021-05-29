package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	aid "github.com/autoai-org/aid/ent/generated"
	"github.com/autoai-org/aid/internal/daemon/generated"
)

func (r *queryResolver) Node(ctx context.Context, id int) (aid.Noder, error) {
	return r.Client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]aid.Noder, error) {
	return r.Client.Noders(ctx, ids)
}

func (r *queryResolver) Images(ctx context.Context) ([]*aid.Image, error) {
	return r.Client.Image.Query().All(ctx)
}

func (r *queryResolver) Repositories(ctx context.Context) ([]*aid.Repository, error) {
	return r.Client.Repository.Query().All(ctx)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
