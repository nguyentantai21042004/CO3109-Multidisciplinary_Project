package user

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name UseCase
type UseCase interface {
	GetOne(ctx context.Context, sc scope.Scope, ip GetOneInput) (models.User, error)
	Detail(ctx context.Context, sc scope.Scope, ID string) (UserOutput, error)
	DetailMe(ctx context.Context, sc scope.Scope) (UserOutput, error)
	Create(ctx context.Context, sc scope.Scope, ip CreateInput) (UserOutput, error)
	UpdateVerified(ctx context.Context, sc scope.Scope, ip UpdateVerifiedInput) (UserOutput, error)
	UpdateAvatar(ctx context.Context, sc scope.Scope, ip UpdateAvatarInput) error
	CheckIn(ctx context.Context, sc scope.Scope, ip CheckInInput) error
}
