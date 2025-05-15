package session

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name UseCase
type UseCase interface {
	Create(ctx context.Context, sc scope.Scope, ip CreateSessionInput) (CreateSessionOutput, error)
	SetUserUseCase(userUC user.UseCase)
}
