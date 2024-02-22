package form

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

// TODO:  to display first-level related entities, try looping over struct field first and then match them with gorm schema
func inspectStruct(entity any) error {
	value := reflect.ValueOf(entity)

	// Check if the given entity is a struct
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("Provided entity is not a struct.")
	}

	// Loop through the fields of the struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := value.Type().Field(i)

		// Check if the field is an array or a slice
		if field.Kind() == reflect.Array || field.Kind() == reflect.Slice {
			// Get the element type of the array or slice
			elemType := field.Type().Elem()
			fmt.Printf("This field is an array/slice with element type: '%s'\n", elemType)
			// Loop over array or slice elements
			for j := 0; j < field.Len(); j++ {
				elemValue := field.Index(j) // get the element at index j
				fmt.Printf("Index: %d, Element Value: %s\n", j, elemValue.Interface())
			}
		}
		// field.Interface() is used to extract the field value as type `interface{}`
		fmt.Printf("Field Name: '%s', Field Type: '%s', Field Value: '%v', Field kind:'%s'\n",
			fieldType.Name, fieldType.Type.Name(), field.Interface(), field.Kind().String())
	}
	return nil
}

func GormToForm(entity any, db *gorm.DB) *Form {
	err := inspectStruct(entity)
	if err != nil {
		panic(err)
	}
	form := NewForm()
	var anyStruct struct{}
	schema := db.Model(entity).First(&anyStruct).Statement.Schema

	for _, field := range schema.Fields {
		switch field.DataType {
		case "string":
			form.AddField(&Field{Name: field.Name, Type: "text", Required: field.NotNull, Value: GetFieldValueByName(entity, field.Name)})
		case "uint":
			form.AddField(&Field{Name: field.Name, Type: "number", Required: field.NotNull, Value: GetFieldValueByName(entity, field.Name)})
		case "int":
			form.AddField(&Field{Name: field.Name, Type: "number", Required: field.NotNull, Value: GetFieldValueByName(entity, field.Name)})
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

	// TODO: return errors
	if value.Kind() != reflect.Struct {
		panic("provided interface is not a struct")
	}

	fieldValue := value.FieldByName(name)
	// TODO: return errors
	if !fieldValue.IsValid() {
		fmt.Printf("No field with name %s found\n", name)
		return ""
	}

	// Zero value.
	if fieldValue.IsZero() {
		return ""
	}

	return fmt.Sprintf("%v", fieldValue.Interface())
}
