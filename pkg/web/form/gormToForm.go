package form

import (
	"fmt"
	"reflect"
)

// TODO:  to display first-level related entities, try looping over struct field first and then match them with gorm schema
func GormToForm(entity any) *Form {
	value := reflect.ValueOf(entity)

	// Check if the given entity is a struct
	if value.Kind() != reflect.Struct {
		panic(fmt.Errorf("Provided entity is not a struct."))
	}

	form := NewForm()

	// Loop through the fields of the struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := value.Type().Field(i)

		// Check if the field is an array or a slice
		if field.Kind() == reflect.Array || field.Kind() == reflect.Slice {
			f := &Field{
				Label: Label{Name: fieldType.Name},
			}
			form.AddField(f)

			// Get the element type of the array or slice
			elemType := field.Type().Elem()
			fmt.Printf("This field is an array/slice with element type: '%s'\n", elemType)
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
					fmt.Printf("Index: %d, UUID Value: %s\n", j, uuid)

					// Use `uuid` to construct the link
					field := &Field{
						// TODO: fix path/to/resource
						Html: fmt.Sprintf(`<a href="/path/to/resource/%s" style="color:blue;">%s</a>`, uuid, uuid)}
					form.AddField(field)
					continue // Continue with the next iteration
				} else {
					// Handle case where Uuid field does not exist
					fmt.Printf("Index: %d, Element does not have a Uuid field\n", j)
				}
			}
		}
		// field.Interface() is used to extract the field value as type `interface{}`
		fmt.Printf("Field Name: '%s', Field Type: '%s', Field Value: '%v', Field kind:'%s'\n", fieldType.Name, fieldType.Type.Name(), field.Interface(), field.Kind().String())
		switch fieldType.Type.Name() {
		case "string":
			field := &Field{
				Label: Label{Name: fieldType.Name},
				Input: &Input{Name: fieldType.Name, Type: "text", Value: fmt.Sprintf("%v", field.Interface())},
			}
			form.AddField(field)
		case "uint", "int":
			field := &Field{
				Label: Label{Name: fieldType.Name},
				Input: &Input{Name: fieldType.Name, Type: "number", Value: fmt.Sprintf("%v", field.Interface())},
			}
			form.AddField(field)
		}
	}
	return form
}
