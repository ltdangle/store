package form

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func GormToForm(entity any, db *gorm.DB) *Form {
	form := NewForm()
	var anyStruct struct{}
	schema := db.Model(entity).First(&anyStruct).Statement.Schema

	for strucFieldName, field := range schema.FieldsByBindName {
		fmt.Println(strucFieldName + ": " + field.DBName)
		switch field.DataType {
		case "string":
			form.AddField(&Field{Name: field.Name, Type: "text", Required: field.NotNull, Value: GetFieldValueByName(entity, field.Name)})
		case "uint":
			form.AddField(&Field{Name: field.Name, Type: "number", Required: field.NotNull, Value: GetFieldValueByName(entity, field.Name)})
		case "int":
			form.AddField(&Field{Name: field.Name, Type: "number", Required: field.NotNull, Value: GetFieldValueByName(entity, field.Name)})
		case "time":
			form.AddField(&Field{Name: field.Name, Type: "time", Required: field.NotNull, Value: GetFieldValueByName(entity, field.Name)})
		}

	}
	return form
}

// TODO: return errors
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
	return fmt.Sprintf("%v", fieldValue.Interface())
}
