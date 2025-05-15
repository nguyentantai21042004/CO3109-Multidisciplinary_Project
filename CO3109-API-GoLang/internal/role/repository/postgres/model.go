package postgres

import (
	"github.com/volatiletech/null/v8"
	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
)

func (r implRepository) buildModel(opts user.CreateOptions) models.User {
	user := models.User{
		Email: opts.Email,
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
