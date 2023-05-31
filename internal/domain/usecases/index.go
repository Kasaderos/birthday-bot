package usecases

import (
	"birthday-bot/internal/adapters/db"
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/domain/core"
)

type St struct {
	lg logger.Lite
	db db.Transaction
	cr *core.St
}

func New(lg logger.Lite, db db.Transaction, cr *core.St) *St {
	return &St{
		lg: lg,
		db: db,
		cr: cr,
	}
}
