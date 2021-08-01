package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/autoai-org/aid/ent/generated"
	generated1 "github.com/autoai-org/aid/internal/daemon/generated"
)

func (r *imageEdgeResolver) Container(ctx context.Context, obj *generated.ImageEdge) (*generated.Container, error) {
	panic(fmt.Errorf("not implemented"))
}

// ImageEdge returns generated1.ImageEdgeResolver implementation.
func (r *Resolver) ImageEdge() generated1.ImageEdgeResolver { return &imageEdgeResolver{r} }

type imageEdgeResolver struct{ *Resolver }
