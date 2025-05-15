package upload

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name UseCase
type UseCase interface {
	Create(ctx context.Context, sc scope.Scope, ip CreateInput) (UploadOutput, error)
	Detail(ctx context.Context, sc scope.Scope, ID string) (UploadOutput, error)
}
