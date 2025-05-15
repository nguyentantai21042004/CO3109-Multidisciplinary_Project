package role

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/pkg/paginator"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name Repository
type Repository interface {
	GetOne(ctx context.Context, sc scope.Scope, opts GetOneOptions) (models.Role, error)
	Detail(ctx context.Context, sc scope.Scope, ID string) (models.Role, error)
	Get(ctx context.Context, sc scope.Scope, opts GetOptions) ([]models.Role, paginator.Paginator, error)
	List(ctx context.Context, sc scope.Scope, opts ListOptions) ([]models.Role, error)
}
