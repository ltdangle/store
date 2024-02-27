package repo

import "github.com/jmoiron/sqlx"

// A repo that can save and find by uuid any mapped entity.
type GeneralRepo struct {
	db *sqlx.DB
}

func NewGeneralRepo(db *sqlx.DB) *GeneralRepo {
	return &GeneralRepo{db: db}
}
