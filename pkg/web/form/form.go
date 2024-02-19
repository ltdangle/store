package form

import (
	"fmt"
	"strings"
)

type Field struct {
	Name        string
	Type        string
	Placeholder string
	Value       string
	Options     []string // For select fields, this would be the list of
	Required    bool
}

type Form struct {
	Method string
	Action string
	Fields []*Field
}

func (form *Form) AddField(field *Field) {
	form.Fields = append(form.Fields, field)
}

func (form *Form) Render() string {
	var fieldsHtml []string
	for _, field := range form.Fields {
		var required string
		if field.Required {
			required = "required"
		}
		formField := fmt.Sprintf(`
  <label for="%s">%s</label>
  <input type="%s" name="%s" placeholder="%s" value="%s" %s >`,
			field.Name, field.Name, field.Type, field.Name, field.Placeholder, field.Value, required)

		fieldsHtml = append(fieldsHtml, formField)
	}

	formHtml := fmt.Sprintf(`
<form action="%s" method="%s" >
	  	%s

  <input type="submit" value="Submit">
</form>`,
		form.Action, form.Method, strings.Join(fieldsHtml, ""),
	)
	return formHtml
}
func (form *Form) Validate() bool {
	return false
}
