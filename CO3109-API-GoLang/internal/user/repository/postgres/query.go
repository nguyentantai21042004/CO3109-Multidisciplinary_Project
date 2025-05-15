package postgres

import (
	"context"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/postgres"
)

func (r implRepository) buildListQuery(ctx context.Context, opts user.ListOptions) ([]qm.QueryMod, error) {
	qr := postgres.BuildQueryWithSoftDelete()

	if len(opts.IDs) > 0 {
		for _, id := range opts.IDs {
			if err := postgres.IsUUID(id); err != nil {
				r.l.Errorf(ctx, "user.repository.postgres.buildListQuery.InvalidID: %v", err)
				return nil, err
			}
		}
		// Create the correct number of placeholders for the IN clause
		placeholders := make([]string, len(opts.IDs))
		for i := range placeholders {
			placeholders[i] = "?"
		}
		qr = append(qr, qm.WhereIn("id IN ("+strings.Join(placeholders, ",")+")", convertToInterfaceSlice(opts.IDs)...))
	}

	return qr, nil
}

func convertToInterfaceSlice(slice []string) []interface{} {
	interfaces := make([]interface{}, len(slice))
	for i, v := range slice {
		interfaces[i] = v
	}
	return interfaces
}

func (r implRepository) buildDetailQuery(ctx context.Context, ID string) ([]qm.QueryMod, error) {
	qr := postgres.BuildQueryWithSoftDelete()

	if err := postgres.IsUUID(ID); err != nil {
		r.l.Errorf(ctx, "user.repository.postgres.buildDetailQuery.InvalidID: %v", err)
		return nil, err
	}
	qr = append(qr, qm.Where("id = ?", ID))

	return qr, nil
}

func (r implRepository) buildGetOneQuery(opts user.GetOneOptions) []qm.QueryMod {
	qr := postgres.BuildQueryWithSoftDelete()

	if opts.Email != "" {
		qr = append(qr, qm.Where("email = ?", opts.Email))
	}

	return qr
}
