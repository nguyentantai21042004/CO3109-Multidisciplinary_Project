package usecase

import (
	"time"

	"gitlab.com/tantai-smap/authenticate-api/internal/core/cloudinary"
	"gitlab.com/tantai-smap/authenticate-api/internal/upload"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
)

type implUsecase struct {
	l          log.Logger
	repo       upload.Repository
	cloudinary cloudinary.Usecase
	userUC     user.UseCase
	clock      func() time.Time
}

var _ upload.UseCase = &implUsecase{}

func New(l log.Logger, repo upload.Repository, cloudinary cloudinary.Usecase, userUC user.UseCase) upload.UseCase {
	return &implUsecase{
		l:          l,
		repo:       repo,
		cloudinary: cloudinary,
		userUC:     userUC,
		clock:      time.Now,
	}
}
