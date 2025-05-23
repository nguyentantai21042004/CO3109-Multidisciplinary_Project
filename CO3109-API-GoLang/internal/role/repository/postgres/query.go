package postgres

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gitlab.com/tantai-smap/authenticate-api/internal/role"
	"gitlab.com/tantai-smap/authenticate-api/pkg/postgres"
)

func (r implRepository) buildDetailQuery(ctx context.Context, ID string) ([]qm.QueryMod, error) {
	qr := postgres.BuildQueryWithSoftDelete()

	if err := postgres.IsUUID(ID); err != nil {
		r.l.Errorf(ctx, "role.repository.postgres.buildDetailQuery.InvalidID: %v", err)
		return nil, err
	}
	qr = append(qr, qm.Where("id = ?", ID))

	return qr, nil
}

func (r implRepository) buildGetQuery(ctx context.Context, opts role.GetOptions) ([]qm.QueryMod, error) {
	qr := postgres.BuildQueryWithSoftDelete()

	if opts.Filter.IDs != nil {
		for _, id := range opts.Filter.IDs {
			if err := postgres.IsUUID(id); err != nil {
				r.l.Errorf(ctx, "role.repository.postgres.buildGetQuery.InvalidID: %v", err)
				return nil, err
			}
			qr = append(qr, qm.Where("id = ?", id))
		}
	}

	if opts.Filter.Alias != nil {
		for _, alias := range opts.Filter.Alias {
			qr = append(qr, qm.Where("alias = ?", alias))
		}
	}

	if opts.Filter.Code != nil {
		for _, code := range opts.Filter.Code {
			qr = append(qr, qm.Where("code = ?", code))
		}
	}

	return qr, nil
}
