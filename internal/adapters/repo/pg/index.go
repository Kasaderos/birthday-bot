package pg

import (
	"birthday-bot/internal/adapters/db"
	"birthday-bot/internal/adapters/logger"

	dopDb "github.com/rendau/dop/adapters/db"
)

type St struct {
	dopDb.RDBConnectionWithHelpers

	lg logger.WarnAndError
	db db.DB
}

func New(lg logger.WarnAndError, dopDb dopDb.RDBConnectionWithHelpers, db db.DB) *St {
	return &St{
		RDBConnectionWithHelpers: dopDb,
		lg:                       lg,
		db:                       db,
	}
}
