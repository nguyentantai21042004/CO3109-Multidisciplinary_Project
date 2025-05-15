package usecase

import (
	"time"

	"gitlab.com/tantai-smap/authenticate-api/internal/role"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
)

type implUsecase struct {
	l     log.Logger
	repo  role.Repository
	clock func() time.Time
}

var _ role.UseCase = &implUsecase{}

func New(l log.Logger, repo role.Repository) role.UseCase {
	return &implUsecase{
		l:     l,
		repo:  repo,
		clock: time.Now,
	}
}
