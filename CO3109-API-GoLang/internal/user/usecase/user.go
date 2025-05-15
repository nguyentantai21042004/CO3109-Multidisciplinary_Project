package usecase

import (
	"context"
	"database/sql"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/internal/role"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/otp"
	"gitlab.com/tantai-smap/authenticate-api/pkg/recognize"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

const (
	RoleUserCode = "GUEST"
)

func (uc implUsecase) GetOne(ctx context.Context, sc scope.Scope, ip user.GetOneInput) (models.User, error) {
	u, err := uc.repo.GetOne(ctx, sc, user.GetOneOptions(ip))
	if err != nil {
		if err == sql.ErrNoRows {
			uc.l.Warnf(ctx, "internal.user.usecase.GetOne.repo.GetOne: %v", err)
			return models.User{}, user.ErrUserNotFound
		}
		uc.l.Errorf(ctx, "internal.user.usecase.GetOne.repo.GetOne: %v", err)
		return models.User{}, err
	}

	return u, nil
}

func (uc implUsecase) Detail(ctx context.Context, sc scope.Scope, ID string) (user.UserOutput, error) {
	u, err := uc.repo.Detail(ctx, sc, ID)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.Detail.repo.Detail: %v", err)
		return user.UserOutput{}, err
	}

	r, err := uc.roleUC.Detail(ctx, sc, u.RoleID)
	if err != nil {
		if err == role.ErrRoleNotFound {
			uc.l.Warnf(ctx, "internal.user.usecase.Detail.roleUC.Detail: %v", err)
			return user.UserOutput{}, user.ErrRoleNotFound
		}
		uc.l.Errorf(ctx, "internal.user.usecase.Detail.roleUC.Detail: %v", err)
		return user.UserOutput{}, err
	}

	return user.UserOutput{
		User: u,
		Role: r.Role,
	}, nil
}

func (uc implUsecase) Create(ctx context.Context, sc scope.Scope, ip user.CreateInput) (user.UserOutput, error) {
	otp, otpExpiredAt := otp.GenerateOTP(uc.clock())

	r, err := uc.roleUC.GetOne(ctx, sc, role.GetOneInput{
		Filter: role.Filter{
			Code: []string{RoleUserCode},
		},
	})
	if err != nil {
		if err == role.ErrRoleNotFound {
			uc.l.Warnf(ctx, "internal.user.usecase.Create.roleUC.GetOne: %v", err)
			return user.UserOutput{}, user.ErrRoleNotFound
		}
		uc.l.Errorf(ctx, "internal.user.usecase.Create.roleUC.GetOne: %v", err)
		return user.UserOutput{}, err
	}

	u, err := uc.repo.Create(ctx, sc, user.CreateOptions{
		Email:        ip.Email,
		Password:     ip.Password,
		FullName:     ip.FullName,
		OTP:          otp,
		OTPExpiredAt: otpExpiredAt,
		IsVerified:   ip.IsVerified,
		RoleID:       r.Role.ID,
		Provider:     ip.Provider,
		ProviderID:   ip.ProviderID,
		AvatarURL:    ip.AvatarURL,
	})
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.Create.repo.Create: %v", err)
		return user.UserOutput{}, err
	}

	return user.UserOutput{
		User: u,
		Role: r.Role,
	}, nil
}

func (uc implUsecase) UpdateVerified(ctx context.Context, sc scope.Scope, ip user.UpdateVerifiedInput) (user.UserOutput, error) {
	u, err := uc.repo.Detail(ctx, sc, ip.UserID)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateVerified.repo.Detail: %v", err)
		return user.UserOutput{}, err
	}

	uo, err := uc.repo.UpdateVerified(ctx, sc, user.UpdateVerifiedOptions{
		ID:           u.ID,
		Otp:          ip.Otp,
		OTPExpiredAt: ip.OtpExpired,
		IsVerified:   ip.IsVerified,
	})
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateVerified.repo.UpdateVerified: %v", err)
		return user.UserOutput{}, err
	}

	return user.UserOutput{
		User: uo,
	}, nil
}

func (uc implUsecase) DetailMe(ctx context.Context, sc scope.Scope) (user.UserOutput, error) {
	u, err := uc.repo.Detail(ctx, sc, sc.UserID)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.DetailMe.repo.Detail: %v", err)
		return user.UserOutput{}, err
	}

	if u.ID != sc.UserID {
		uc.l.Warnf(ctx, "internal.user.usecase.DetailMe.repo.Detail: %v", "user id not match")
		return user.UserOutput{}, user.ErrPermissionDenied
	}

	r, err := uc.roleUC.Detail(ctx, sc, u.RoleID)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.DetailMe.roleUC.Detail: %v", err)
		return user.UserOutput{}, err
	}

	return user.UserOutput{
		User: u,
		Role: r.Role,
	}, nil
}

func (uc implUsecase) UpdateAvatar(ctx context.Context, sc scope.Scope, ip user.UpdateAvatarInput) error {
	_, err := uc.getAndCheckUserPermission(ctx, sc, ip.UserID)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.getAndCheckUserPermission: %v", err)
		return err
	}

	// Update the avatar
	_, err = uc.repo.UpdateAvatar(ctx, sc, user.UpdateAvatarOptions{
		ID:        ip.UserID,
		AvatarURL: ip.AvatarURL,
	})
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.repo.UpdateAvatar: %v", err)
		return err
	}

	err = uc.syncAvatarUser(ctx, ip.AvatarURL, ip.UserID)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.syncAvatarUser: %v", err)
		return err
	}

	return nil
}

func (uc implUsecase) CheckIn(ctx context.Context, sc scope.Scope, ip user.CheckInInput) error {
	req, err := recognize.CreateFindImagesRequest(ip.File, ip.ShopID)
	if err != nil {
		uc.l.Errorf(ctx, "internal.user.usecase.UpdateAvatar.recognize.CreateSaveImageRequest: %v", err)
		return err
	}

	_, err = recognize.SendFindImagesRequest(req)
	if err != nil {
		if err == recognize.ErrUserNotFound {
			uc.l.Warnf(ctx, "internal.user.usecase.CheckIn.recognize.SendFindImagesRequest: %v", err)
			return err
		}
		uc.l.Errorf(ctx, "internal.user.usecase.CheckIn.recognize.SendFindImagesRequest: %v", err)
		return err
	}

	// pub msg

	return nil
}
