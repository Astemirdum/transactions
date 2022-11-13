package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type BalanceDB struct {
	db       *sqlx.DB
	log      *zap.Logger
	initCash int
}

func NewBalanceDB(db *sqlx.DB, log *zap.Logger, initCash int) *BalanceDB {
	return &BalanceDB{
		db:       db,
		log:      log,
		initCash: initCash,
	}
}
