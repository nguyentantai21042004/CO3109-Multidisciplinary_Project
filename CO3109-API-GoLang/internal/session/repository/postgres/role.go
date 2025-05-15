package postgres

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/internal/session"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (r implRepository) Create(ctx context.Context, sc scope.Scope, ip session.CreateSessionOptions) (models.Session, error) {
	session, err := r.buildModel(ctx, ip)
	if err != nil {
		r.l.Errorf(ctx, "session.repository.postgres.Create.buildModel: %v", err)
		return models.Session{}, err
	}

	err = session.Insert(ctx, r.database, boil.Infer())
	if err != nil {
		r.l.Errorf(ctx, "session.repository.postgres.Create.Insert: %v", err)
		return models.Session{}, err
	}
	return session, nil
}
