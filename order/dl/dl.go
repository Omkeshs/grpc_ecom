package dl

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ProductDL struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewDL(logger *zap.Logger, db *sqlx.DB) *ProductDL {
	return &ProductDL{
		logger: logger,
		db:     db,
	}
}
