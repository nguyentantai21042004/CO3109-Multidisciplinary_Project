package usecase

import (
	"context"
	"database/sql"

	"gitlab.com/tantai-smap/authenticate-api/internal/upload"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (uc *implUsecase) Create(ctx context.Context, sc scope.Scope, ip upload.CreateInput) (upload.UploadOutput, error) {
	// validate input
	u, err := uc.userUC.Detail(ctx, sc, sc.UserID)
	if err != nil {
		if err == user.ErrUserNotFound {
			uc.l.Warnf(ctx, "internal.upload.usecase.Create.userUC.Detail: %v", err)
			return upload.UploadOutput{}, upload.ErrUnauthorized
		}
		uc.l.Errorf(ctx, "internal.upload.usecase.Create.userUC.Detail: %v", err)
		return upload.UploadOutput{}, err
	}

	f, err := uc.cloudinary.Upload(ctx, ip.FileHeader, ip.From)
	if err != nil {
		uc.l.Errorf(ctx, "internal.upload.usecase.Create.cloudinary.Upload: %v", err)
		return upload.UploadOutput{}, err
	}

	up, err := uc.repo.Create(ctx, sc, upload.CreateOptions{
		Name:          f.PublicID,
		Path:          f.URL,
		Source:        f.Type,
		FromLocation:  ip.From,
		PublicID:      f.PublicID,
		CreatedUserID: sc.UserID,
	})
	if err != nil {
		uc.l.Errorf(ctx, "internal.upload.usecase.Create.repo.Create: %v", err)
		return upload.UploadOutput{}, err
	}

	return upload.UploadOutput{
		User:   u.User,
		Upload: up,
	}, nil
}

func (uc *implUsecase) Detail(ctx context.Context, sc scope.Scope, ID string) (upload.UploadOutput, error) {
	u, err := uc.userUC.Detail(ctx, sc, sc.UserID)
	if err != nil {
		if err == user.ErrUserNotFound {
			uc.l.Warnf(ctx, "internal.upload.usecase.Detail.userUC.Detail: %v", err)
			return upload.UploadOutput{}, upload.ErrUnauthorized
		}
		uc.l.Errorf(ctx, "internal.upload.usecase.Detail.userUC.Detail: %v", err)
		return upload.UploadOutput{}, err
	}

	up, err := uc.repo.Detail(ctx, sc, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			uc.l.Warnf(ctx, "internal.upload.usecase.Detail.repo.Detail: %v", err)
			return upload.UploadOutput{}, upload.ErrUploadNotFound
		}
		uc.l.Errorf(ctx, "internal.upload.usecase.Detail.repo.Detail: %v", err)
		return upload.UploadOutput{}, err
	}

	return upload.UploadOutput{
		User:   u.User,
		Upload: up,
	}, nil
}
