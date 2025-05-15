package upload

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, sc scope.Scope, opts CreateOptions) (models.Upload, error)
	Detail(ctx context.Context, sc scope.Scope, ID string) (models.Upload, error)
}
