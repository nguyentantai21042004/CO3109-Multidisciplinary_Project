package postgres

import (
	"context"

	"github.com/volatiletech/null/v8"
	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	"gitlab.com/tantai-smap/authenticate-api/internal/upload"
	"gitlab.com/tantai-smap/authenticate-api/pkg/postgres"
)

func (r implRepository) buildModel(ctx context.Context, opts upload.CreateOptions) models.Upload {
	upload := models.Upload{}

	if opts.Name != "" {
		upload.Name = opts.Name
	}

	if opts.Path != "" {
		upload.Path = opts.Path
	}

	if opts.Source != "" {
		upload.Source = opts.Source
	}

	if opts.FromLocation != "" {
		upload.FromLocation = opts.FromLocation
	}

	if opts.PublicID != "" {
		upload.PublicID = null.NewString(opts.PublicID, true)
	}

	if opts.CreatedUserID != "" {
		if err := postgres.IsUUID(opts.CreatedUserID); err != nil {
			r.l.Errorf(ctx, "internal.upload.repository.postgres.buildModel.IsUUID: %v", err)
			return models.Upload{}
		}
		upload.CreatedUserID = opts.CreatedUserID
	}

	upload.CreatedAt = null.NewTime(r.clock(), true)
	upload.UpdatedAt = null.NewTime(r.clock(), true)
	return upload
}
