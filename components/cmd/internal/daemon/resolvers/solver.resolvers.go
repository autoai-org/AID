package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/autoai-org/aid/ent/generated"
	generated1 "github.com/autoai-org/aid/internal/daemon/generated"
)

func (r *solverEdgeResolver) Repository(ctx context.Context, obj *generated.SolverEdge) (*generated.Repository, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *solverEdgeResolver) Image(ctx context.Context, obj *generated.SolverEdge) (*generated.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// SolverEdge returns generated1.SolverEdgeResolver implementation.
func (r *Resolver) SolverEdge() generated1.SolverEdgeResolver { return &solverEdgeResolver{r} }

type solverEdgeResolver struct{ *Resolver }
