package postgres

import (
	"context"

	"github.com/volatiletech/null/v8"
	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/postgres"
)

func (r implRepository) buildModel(ctx context.Context, opts user.CreateOptions) models.User {
	user := models.User{
		Email: opts.Email,
	}

	if opts.RoleID != "" {
		if err := postgres.IsUUID(opts.RoleID); err != nil {
			r.l.Errorf(ctx, "user.repository.postgres.buildModel.IsUUID: %v", err)
			return models.User{}
		}
		user.RoleID = opts.RoleID
	}

	if opts.Password != "" {
		user.PasswordHash = null.NewString(opts.Password, true)
	}

	if opts.FullName != "" {
		user.FullName = null.NewString(opts.FullName, true)
	}

	if opts.IsVerified {
		user.IsVerified = null.NewBool(true, true)
	}

	if opts.AvatarURL != "" {
		user.AvatarURL = null.NewString(opts.AvatarURL, true)
	}

	if opts.Provider != "" {
		user.Provider = null.NewString(opts.Provider, true)
	}

	if opts.ProviderID != "" {
		user.ProviderID = null.NewString(opts.ProviderID, true)
	}

	if opts.OTP != "" {
		user.Otp = null.NewString(opts.OTP, true)
	}

	if !opts.OTPExpiredAt.IsZero() {
		user.OtpExpiredAt = null.NewTime(opts.OTPExpiredAt, true)
	}

	return user
}

func (r implRepository) buildUpdateVerifiedModel(ctx context.Context, opts user.UpdateVerifiedOptions) (models.User, []string, error) {
	user := models.User{
		IsVerified: null.NewBool(false, true),
		CreatedAt:  null.NewTime(r.clock(), true),
	}
	columns := make([]string, 0)
	columns = append(columns, models.UserColumns.CreatedAt)

	if err := postgres.IsUUID(opts.ID); err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.buildUpdateVerifiedModel.IsUUID: %v", err)
		return models.User{}, nil, err
	}
	user.ID = opts.ID

	if opts.Otp != "" {
		user.Otp = null.NewString(opts.Otp, true)
		columns = append(columns, models.UserColumns.Otp)
	}

	if !opts.OTPExpiredAt.IsZero() {
		user.OtpExpiredAt = null.NewTime(opts.OTPExpiredAt, true)
		columns = append(columns, models.UserColumns.OtpExpiredAt)
	}

	if opts.IsVerified {
		user.IsVerified = null.NewBool(true, true)
		columns = append(columns, models.UserColumns.IsVerified)
	}

	user.UpdatedAt = null.NewTime(r.clock(), true)
	columns = append(columns, models.UserColumns.UpdatedAt)

	return user, columns, nil
}

func (r implRepository) buildUpdateAvatarModel(ctx context.Context, opts user.UpdateAvatarOptions) (models.User, []string, error) {
	user := models.User{
		AvatarURL: null.NewString(opts.AvatarURL, true),
	}
	columns := make([]string, 0)
	columns = append(columns, models.UserColumns.AvatarURL)

	if err := postgres.IsUUID(opts.ID); err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.buildUpdateAvatarModel.IsUUID: %v", err)
		return models.User{}, nil, err
	}
	user.ID = opts.ID

	user.UpdatedAt = null.NewTime(r.clock(), true)
	columns = append(columns, models.UserColumns.UpdatedAt)

	return user, columns, nil
}
