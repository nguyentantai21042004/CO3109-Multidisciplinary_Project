package postgres

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (r implRepository) List(ctx context.Context, sc scope.Scope, opts user.ListOptions) ([]models.User, error) {
	qr, err := r.buildListQuery(ctx, opts)
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.List.buildListQuery: %v", err)
		return nil, err
	}

	var users []models.User
	err = models.Users(qr...).Bind(ctx, r.database, &users)
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.List.Bind: %v", err)
		return nil, err
	}

	return users, nil
}

func (r implRepository) Detail(ctx context.Context, sc scope.Scope, ID string) (models.User, error) {
	qr, err := r.buildDetailQuery(ctx, ID)
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.Detail.buildDetailQuery: %v", err)
		return models.User{}, err
	}

	user, err := models.Users(qr...).One(ctx, r.database)
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.Detail.One: %v", err)
		return models.User{}, err
	}

	return *user, nil
}

func (r implRepository) GetOne(ctx context.Context, sc scope.Scope, opts user.GetOneOptions) (models.User, error) {
	qr := r.buildGetOneQuery(opts)

	user, err := models.Users(qr...).One(ctx, r.database)
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.GetOne.One: %v", err)
		return models.User{}, err
	}

	return *user, nil
}

func (r implRepository) Create(ctx context.Context, sc scope.Scope, opts user.CreateOptions) (models.User, error) {
	user := r.buildModel(ctx, opts)

	err := user.Insert(ctx, r.database, boil.Infer())
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.Create.Insert: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (r implRepository) UpdateVerified(ctx context.Context, sc scope.Scope, opts user.UpdateVerifiedOptions) (models.User, error) {
	u, col, err := r.buildUpdateVerifiedModel(ctx, opts)
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.UpdateVerified.buildUpdateVerifiedModel: %v", err)
		return models.User{}, err
	}

	_, err = u.Update(ctx, r.database, boil.Whitelist(col...))
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.UpdateVerified.Update: %v", err)
		return models.User{}, err
	}

	return u, nil
}

func (r implRepository) UpdateAvatar(ctx context.Context, sc scope.Scope, opts user.UpdateAvatarOptions) (models.User, error) {
	u, col, err := r.buildUpdateAvatarModel(ctx, opts)
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.UpdateAvatar.buildUpdateAvatarModel: %v", err)
		return models.User{}, err
	}

	_, err = u.Update(ctx, r.database, boil.Whitelist(col...))
	if err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.UpdateAvatar.Update: %v", err)
		return models.User{}, err
	}

	return u, nil
}
