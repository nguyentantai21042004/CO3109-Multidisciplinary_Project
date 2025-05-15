package role

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name UseCase
type UseCase interface {
	Detail(ctx context.Context, sc scope.Scope, ID string) (DetailOutput, error)
	GetOne(ctx context.Context, sc scope.Scope, ip GetOneInput) (GetOneOutput, error)
	Get(ctx context.Context, sc scope.Scope, ip GetInput) (GetOutput, error)
	List(ctx context.Context, sc scope.Scope, ip ListInput) (ListOutput, error)
}
