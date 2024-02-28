package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type MappedEntity interface {
	// Primary key of the table.
	PrimaryKey() string
	// Name of the database table for the entity.
	TableName() string
}

// A repo that can save and find by uuid any mapped entity.
type GeneralRepo struct {
	gorm *gorm.DB
	sqlx *sqlx.DB
}

func NewGeneralRepo(sqlx *sqlx.DB, gorm *gorm.DB) *GeneralRepo {
	return &GeneralRepo{sqlx: sqlx, gorm: gorm}
}

// Save / Update entity via GORM (convenient).
func (repo *GeneralRepo) Save(entity any) error {
	tx := repo.gorm.Save(entity)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Retrieve entity via sqlx (GORM produces unexpected queries).
func (repo *GeneralRepo) GetByPrimaryKey(entity MappedEntity, search string) error {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s = $1;`, entity.TableName(), entity.PrimaryKey())
	err := repo.sqlx.Get(entity, query, search)
	if err != nil {
		return err
	}

	return nil
}
func (repo *GeneralRepo) FindAll(entity MappedEntity, results interface{}) error {
	query := fmt.Sprintf(`SELECT * FROM %s;`, entity.TableName())
	err := repo.sqlx.Select(results, query)
	if err != nil {
		return fmt.Errorf("sqlx: %v", err)
	}

	return nil
}
