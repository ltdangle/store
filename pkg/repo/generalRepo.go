package repo

import (
	"fmt"
	"store/pkg/i"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

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
func (repo *GeneralRepo) GetByPrimaryKey(entity i.AdminEntity, search string) error {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s = $1;`, entity.TableName(), entity.PrimaryKey())
	err := repo.sqlx.Get(entity, query, search)
	if err != nil {
		return err
	}

	return nil
}
func (repo *GeneralRepo) FindAll(entity i.AdminEntity, results interface{}) error {
	query := fmt.Sprintf(`SELECT * FROM %s;`, entity.TableName())
	err := repo.sqlx.Select(results, query)
	if err != nil {
		return fmt.Errorf("sqlx: %v", err)
	}

	return nil
}

type QueryToMapResult struct {
	ColumnNames []string
	DataMap     []map[string]interface{}
}

func (repo *GeneralRepo) QueryToMap(query string) (*QueryToMapResult, error) {
	result := &QueryToMapResult{}
	// Execute the query
	rows, err := repo.sqlx.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result.ColumnNames, err = rows.Columns()
	if err != nil {
		return nil, err
	}

	// Iterate over each row
	for rows.Next() {
		rowMap := make(map[string]interface{})
		// MapScan will fill the map with column name-value pairs
		err := rows.MapScan(rowMap)
		if err != nil {
			return nil, err
		}
		// Append the map to the results slice
		result.DataMap = append(result.DataMap, rowMap)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
