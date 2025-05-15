package usecase

import (
	"time"

	"gitlab.com/tantai-smap/authenticate-api/internal/session"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
)

type implUsecase struct {
	l      log.Logger
	repo   session.Repository
	userUC user.UseCase
	clock  func() time.Time
}

var _ session.UseCase = &implUsecase{}

func New(l log.Logger, repo session.Repository, userUC user.UseCase) session.UseCase {
	return &implUsecase{
		l:      l,
		repo:   repo,
		userUC: userUC,
		clock:  time.Now,
	}
}

func (uc *implUsecase) SetUserUseCase(userUC user.UseCase) {
	uc.userUC = userUC
}
