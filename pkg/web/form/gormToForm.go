package form

import (
	"store/pkg/models"

	"gorm.io/gorm"
)

func GormToForm(entity any, db *gorm.DB) *Form {
	form := NewForm()

	var columnNames []string
	columns, err := db.Migrator().ColumnTypes(&models.Cart{})

	if err == nil {
		for _, column := range columns {
			columnNames = append(columnNames, column.Name()+" - "+column.DatabaseTypeName()+"-"+column.ScanType().String()+";")

			isRequired, _ := column.Nullable()
			switch column.ScanType().String() {
			case "string":
				form.AddField(&Field{Name: column.Name(), Type: "text", Required: isRequired})

			case "int64":
				form.AddField(&Field{Name: column.Name(), Type: "number", Required: isRequired})
			}
		}
	}
	// fmt.Println(columnNames)
	return form
}
