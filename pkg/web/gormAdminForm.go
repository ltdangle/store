package web

import (
	"fmt"
	"reflect"
	forms "store/pkg/web/form"
)

func GormAdminForm(entity any) (*forms.Form, error) {
	value := reflect.ValueOf(entity)

	// Check if the given entity is a struct
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Provided entity is not a struct.")
	}

	form := forms.NewForm()

	// Loop through the fields of the struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := value.Type().Field(i)

		// fmt.Printf("Field Name: '%s', Field Type: '%s', Field Value: '%v', Field kind:'%s'\n", fieldType.Name, fieldType.Type.Name(), field.Interface(), field.Kind().String())
		switch fieldType.Type.Name() {
		case "string":
			field := &forms.Field{
				Label: forms.Label{Name: fieldType.Name},
				Input: &forms.Input{Name: fieldType.Name, Type: "text", Value: fmt.Sprintf("%v", field.Interface())},
			}
			form.AddField(field)
		case "uint", "int":
			field := &forms.Field{
				Label: forms.Label{Name: fieldType.Name},
				Input: &forms.Input{Name: fieldType.Name, Type: "number", Value: fmt.Sprintf("%v", field.Interface())},
			}
			form.AddField(field)
		}
	}
	return form, nil
}
