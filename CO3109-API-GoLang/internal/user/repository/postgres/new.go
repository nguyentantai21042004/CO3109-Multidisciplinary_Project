package postgres

import (
	"database/sql"
	"time"

	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
	"gitlab.com/tantai-smap/authenticate-api/pkg/util"
)

type implRepository struct {
	l        log.Logger
	database *sql.DB
	clock    func() time.Time
}

var _ user.Repository = implRepository{}

func New(l log.Logger, database *sql.DB) implRepository {
	return implRepository{
		l:        l,
		database: database,
		clock:    util.Now,
	}
}
