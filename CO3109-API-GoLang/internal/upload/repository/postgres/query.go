package postgres

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gitlab.com/tantai-smap/authenticate-api/pkg/postgres"
)

func (r implRepository) buildDetailQuery(ctx context.Context, ID string) ([]qm.QueryMod, error) {
	qr := postgres.BuildQueryWithSoftDelete()

	if err := postgres.IsUUID(ID); err != nil {
		r.l.Errorf(ctx, "internal.upload.repository.postgres.buildDetailQuery.InvalidID: %v", err)
		return nil, err
	}
	qr = append(qr, qm.Where("id = ?", ID))

	return qr, nil
}
