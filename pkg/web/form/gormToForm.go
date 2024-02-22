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
				fmt.Printf("Index: %d, Element Value: %s\n", j, elemValue.Interface())
				field := &Field{
					Html:  fmt.Sprintf(`<a href="%s" style="color:blue;">%s</a>`, elemValue.Interface(), elemValue.Interface()),
				}
				form.AddField(field)
				continue
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
