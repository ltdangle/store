package form

import (
	"bytes"
	"io"
	"text/template"
)

type FormField struct {
	Name        string
	Type        string
	Placeholder string
	Value       string
	Options     []string // For select fields, this would be the list of
	Required    bool
}

type DynamicForm struct {
	Action string
	Method string
	Fields []FormField
}

const formTemplateStr = `
    <form action="{{.Action}}" method="{{.Method}}">
        {{range .Fields}}
            {{if eq .Type "text" "password" "email"}}
                <label for="{{.Name}}">{{.Name}}</label>
                <input type="{{.Type}}" name="{{.Name}}"
  placeholder="{{.Placeholder}}" value="{{.Value}}" {{if
  .Required}}required{{end}}>
            {{else if eq .Type "select"}}
                <label for="{{.Name}}">{{.Name}}</label>
                <select name="{{.Name}}" {{if .Required}}required{{end}}>
                    {{range .Options}}
                        <option value="{{.}}">{{.}}</option>
                    {{end}}
                </select>
            {{end}}
            <br>
        {{end}}
        <input type="submit" value="Submit">
    </form>
    `

func gptForm() string {
	dynamicForm := DynamicForm{
		Action: "/submit",
		Method: "POST",
		Fields: []FormField{
			{Name: "username", Type: "text", Placeholder: "Enter username", Required: true},
			{Name: "password", Type: "password", Placeholder: "Enter password", Required: true},
			{Name: "email", Type: "email", Placeholder: "Enter email", Required: true},
			{Name: "gender", Type: "select", Options: []string{"Male", "Female", "Other"}, Required: true},
			// Add more fields as needed
		},
	}

	var buff bytes.Buffer
	renderDynamicForm(&buff, dynamicForm)

	return buff.String()
}
func renderDynamicForm(w io.Writer, form DynamicForm) {
	tmpl, err := template.New("form").Parse(formTemplateStr)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(w, form); err != nil {
		panic(err)
	}
}
