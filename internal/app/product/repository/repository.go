package repository

import (
	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/contract"
	"github.com/jmoiron/sqlx"
)

type RProduct struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) contract.IRProduct {
	return &RProduct{
		db: db,
	}
}
