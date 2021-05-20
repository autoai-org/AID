package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/autoai-org/aid/ent/generated"
	generated1 "github.com/autoai-org/aid/internal/daemon/generated"
)

func (r *edgeResolver) Node(ctx context.Context, obj *generated.Edge) (generated.Noder, error) {
	panic("no support!")
}

func (r *edgeResolver) Cursor(ctx context.Context, obj *generated.Edge) (*generated.Cursor, error) {
	panic("no support!")
}

func (r *pageInfoResolver) HasPrevPage(ctx context.Context, obj *generated.PageInfo) (bool, error) {
	panic("no support!")
}

// Edge returns generated1.EdgeResolver implementation.
func (r *Resolver) Edge() generated1.EdgeResolver { return &edgeResolver{r} }

// PageInfo returns generated1.PageInfoResolver implementation.
func (r *Resolver) PageInfo() generated1.PageInfoResolver { return &pageInfoResolver{r} }

type edgeResolver struct{ *Resolver }
type pageInfoResolver struct{ *Resolver }
