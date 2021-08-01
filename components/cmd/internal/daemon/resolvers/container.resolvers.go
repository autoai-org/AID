package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/autoai-org/aid/ent/generated"
	generated1 "github.com/autoai-org/aid/internal/daemon/generated"
)

func (r *containerEdgeResolver) Port(ctx context.Context, obj *generated.Container) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Container returns generated1.ContainerResolver implementation.
func (r *Resolver) ContainerEdge() generated1.ResolverRoot { return &containerEdgeResolver{r} }

type containerEdgeResolver struct{ *Resolver }
