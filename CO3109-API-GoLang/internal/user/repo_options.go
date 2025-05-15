package user

import "time"

type GetOneOptions struct {
	Email string
}

type CreateOptions struct {
	Email        string
	Password     string
	FullName     string
	IsVerified   bool
	AvatarURL    string
	Provider     string
	ProviderID   string
	OTP          string
	OTPExpiredAt time.Time
	RoleID       string
}

type UpdateVerifiedOptions struct {
	ID           string
	Otp          string
	OTPExpiredAt time.Time
	IsVerified   bool
}

type ListOptions struct {
	IDs []string
}

type UpdateAvatarOptions struct {
	ID        string
	AvatarURL string
}
