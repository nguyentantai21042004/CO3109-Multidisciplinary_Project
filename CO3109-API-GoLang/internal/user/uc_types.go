package user

import (
	"mime/multipart"
	"strings"
	"time"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
)

type DetailOutput struct {
	User models.User
	Role models.Role
}

type GetOneInput struct {
	Email string
}

type CreateInput struct {
	Provider   string
	ProviderID string
	Email      string
	Password   string
	FullName   string
	AvatarURL  string
	IsVerified bool
}

type UserOutput struct {
	User models.User
	Role models.Role
}

type UpdateVerifiedInput struct {
	UserID     string
	Otp        string
	OtpExpired time.Time
	IsVerified bool
}

type UpdateAvatarInput struct {
	UserID    string
	AvatarURL string
}

var CommonImageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".svg", ".tiff", ".tif", ".ico", ".heic", ".heif", ".raw", ".cr2", ".nef", ".arw"}

func IsValidImageExtension(imageURL string) bool {
	for _, ext := range CommonImageExtensions {
		if strings.HasSuffix(imageURL, ext) {
			return true
		}
	}
	return false
}

type CheckInInput struct {
	ShopID string
	File   *multipart.FileHeader
}
