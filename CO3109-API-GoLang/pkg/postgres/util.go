package postgres

import (
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// IsUUID checks if the given string is a valid UUID
func IsUUID(u string) error {
	_, err := uuid.Parse(u)
	if err != nil {
		return err
	}
	return nil
}

func NewUUID() string {
	return uuid.New().String()
}

func BuildQueryWithSoftDelete() []qm.QueryMod {
	return []qm.QueryMod{
		qm.Where("deleted_at IS NULL"),
	}
}
