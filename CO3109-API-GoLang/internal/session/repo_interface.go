package session

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, sc scope.Scope, ip CreateSessionOptions) (models.Session, error)
}
