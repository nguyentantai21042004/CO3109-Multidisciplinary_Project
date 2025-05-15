package session

import (
	"time"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
)

type CreateSessionInput struct {
	UserID       string
	AccessToken  string
	RefreshToken string
	UserAgent    string
	IPAddress    string
	DeviceName   string
	ExpiresAt    time.Time
}

type CreateSessionOutput struct {
	models.Session
}
