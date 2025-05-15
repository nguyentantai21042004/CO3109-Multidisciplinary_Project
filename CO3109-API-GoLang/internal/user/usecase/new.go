package usecase

import (
	"time"

	"gitlab.com/tantai-smap/authenticate-api/internal/role"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
)

type implUsecase struct {
	l      log.Logger
	repo   user.Repository
	roleUC role.UseCase
	clock  func() time.Time
}

var _ user.UseCase = &implUsecase{}

func New(l log.Logger, repo user.Repository, roleUC role.UseCase) user.UseCase {
	return &implUsecase{
		l:      l,
		repo:   repo,
		roleUC: roleUC,
		clock:  time.Now,
	}
}
