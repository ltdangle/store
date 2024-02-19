package form

import (
	"fmt"
	"reflect"
	"store/pkg/models"

	"gorm.io/gorm"
)

func GormToForm(entity any, db *gorm.DB) *Form {
	form := NewForm()
	var anyStruct struct{}
	schema := db.Model(&models.Cart{}).First(&anyStruct).Statement.Schema
	fmt.Println(schema)
	for strucFieldName, field := range schema.FieldsByBindName {
		fmt.Println(strucFieldName + ": " + field.DBName)
	}
	columns, err := db.Migrator().ColumnTypes(entity)

	if err == nil {
		for _, column := range columns {

			isRequired, _ := column.Nullable()
			switch column.ScanType().String() {
			case "string":
				form.AddField(&Field{Name: column.Name(), Type: "text", Required: isRequired, Value: GetFieldValueByName(entity, column.Name())})

			case "int64":
				form.AddField(&Field{Name: column.Name(), Type: "number", Required: isRequired})
			}
		}
	}
	return form
}

func GetFieldValueByName(data interface{}, name string) string {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		fmt.Println("provided interface is not a struct")
	}

	fieldValue := value.FieldByName(name)
	if !fieldValue.IsValid() {
		fmt.Printf("No field with name %s found\n", name)
		return ""
	}

	var str string
	if str, ok := fieldValue.Interface().(string); ok {
		// The assertion was successful, str is of type string now
		fmt.Println(str)
	} else {
		// The assertion has failed, myInterface is not of type string
		fmt.Println("Value is not a string")
	}
	return str
}
