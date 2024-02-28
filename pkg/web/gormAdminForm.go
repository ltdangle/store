package web

import (
	"fmt"
	"reflect"
	forms "store/pkg/web/form"
)

func GormAdminForm(entity any, entityName string, primaryKey string, router *AppRouter) (*forms.Form, error) {
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
		// Check if the field is an array or a slice
		if field.Kind() == reflect.Array || field.Kind() == reflect.Slice {
			f := &forms.Field{
				Label: forms.Label{Name: fieldType.Name},
			}
			form.AddField(f)

			// Loop over array or slice elements
			for j := 0; j < field.Len(); j++ {
				elemValue := field.Index(j) // get the element at index j

				// If elemValue is a pointer to a struct, we need to dereference it.
				if elemValue.Kind() == reflect.Ptr {
					elemValue = elemValue.Elem()
				}

				// Attempt to get the Uuid field
				uuidField := elemValue.FieldByName("Uuid")
				if uuidField.IsValid() { // Check if the Uuid field exists
					uuid := uuidField.Interface()
					url := router.UrlInternal(ADMIN_VIEW_ENTITY_ROUTE, "entity", entityName, primaryKey, fmt.Sprintf("%v", uuid))
					field := &forms.Field{
						// TODO: fix path/to/resource
						Html: fmt.Sprintf(`<a href="%s" style="color:blue;">%s</a>`, url.Value, uuid)}
					form.AddField(field)
					continue // Continue with the next iteration
				} else {
					//TODO: Handle case where Uuid field does not exist
					fmt.Printf("Index: %d, Element does not have a Uuid field\n", j)
				}
			}
		}
		// field.Interface() is used to extract the field value as type `interface{}`
		fmt.Printf("Field Name: '%s', Field Type: '%s', Field Value: '%v', Field kind:'%s'\n", fieldType.Name, fieldType.Type.Name(), field.Interface(), field.Kind().String())
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
