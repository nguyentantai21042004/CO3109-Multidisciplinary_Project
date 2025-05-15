package user

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name Repository
type Repository interface {
	List(ctx context.Context, sc scope.Scope, opts ListOptions) ([]models.User, error)
	Detail(ctx context.Context, sc scope.Scope, ID string) (models.User, error)
	GetOne(ctx context.Context, sc scope.Scope, opts GetOneOptions) (models.User, error)
	Create(ctx context.Context, sc scope.Scope, opts CreateOptions) (models.User, error)
	UpdateVerified(ctx context.Context, sc scope.Scope, opts UpdateVerifiedOptions) (models.User, error)
	UpdateAvatar(ctx context.Context, sc scope.Scope, opts UpdateAvatarOptions) (models.User, error)
}
