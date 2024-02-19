package form

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

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
		fmt.Println("Struct value for " + name + " is " + str)
		return str
	}
	if unsigned_int, ok := fieldValue.Interface().(uint); ok {
		str := strconv.Itoa(int(unsigned_int))
		fmt.Println("Struct value for " + name + " is " + str)
		return str
	}
	if unsigned_int, ok := fieldValue.Interface().(int); ok {
		str := strconv.Itoa(int(unsigned_int))
		fmt.Println("Struct value for " + name + " is " + str)
		return str
	}
	if time_type, ok := fieldValue.Interface().(time.Time); ok {
		str = time_type.String()
		fmt.Println("Struct value for " + name + " is " + str)
		return str
	}

	fmt.Println("Value is not a string")

	return str
}
