package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	aid "github.com/autoai-org/aid/ent/generated"
	generated1 "github.com/autoai-org/aid/internal/daemon/generated"
)

func (r *queryResolver) Node(ctx context.Context, id int) (aid.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]aid.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Images(ctx context.Context, after *aid.Cursor, first *int, before *aid.Cursor, last *int, orderBy *aid.ImageOrder) (*aid.ImageConnection, error) {
	return r.client.Image.Query().
		Paginate(ctx, after, first, before, last,
			aid.WithImageOrder(orderBy),
		)
}

func (r *queryResolver) Repositories(ctx context.Context, after *aid.Cursor, first *int, before *aid.Cursor, last *int, orderBy *aid.RepositoryOrder) (*aid.RepositoryConnection, error) {
	return r.client.Repository.Query().Paginate(ctx, after, first, before, last,
		aid.WithRepositoryOrder(orderBy),
	)
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
