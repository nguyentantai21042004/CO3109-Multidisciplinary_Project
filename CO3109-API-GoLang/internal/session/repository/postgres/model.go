package postgres

import (
	"context"

	"github.com/volatiletech/null/v8"
	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/internal/session"
	"gitlab.com/tantai-smap/authenticate-api/pkg/postgres"
)

func (r implRepository) buildModel(ctx context.Context, opts session.CreateSessionOptions) (models.Session, error) {
	session := models.Session{}

	if opts.UserID != "" {
		if err := postgres.IsUUID(opts.UserID); err != nil {
			r.l.Errorf(ctx, "session.repository.postgres.buildModel.InvalidUserID: %v", err)
			return models.Session{}, err
		}
		session.UserID = opts.UserID
	}

	if opts.AccessToken != "" {
		session.AccessToken = opts.AccessToken
	}

	if opts.RefreshToken != "" {
		session.RefreshToken = opts.RefreshToken
	}

	if opts.UserAgent != "" {
		session.UserAgent = null.NewString(opts.UserAgent, true)
	}

	if opts.IPAddress != "" {
		session.IPAddress = null.NewString(opts.IPAddress, true)
	}

	if opts.DeviceName != "" {
		session.DeviceName = null.NewString(opts.DeviceName, true)
	}

	if !opts.ExpiresAt.IsZero() {
		session.ExpiresAt = opts.ExpiresAt
	}

	session.CreatedAt = null.NewTime(r.clock(), true)
	session.UpdatedAt = null.NewTime(r.clock(), true)
	return session, nil
}
