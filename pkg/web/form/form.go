package form

import (
	"fmt"
	"strings"
)

type Label struct {
	Name string
}

type Input struct {
	Name        string
	Type        string
	Placeholder string
	Value       string
	Options     []string // For select fields, this would be the list of
	Required    bool
	Error       string
}
type Field struct {
	Label Label
	Input Input
	Html  string
}

func NewInput() *Input {
	return &Input{}
}

type Form struct {
	Method string
	Action string
	Fields []*Field
}

func NewForm() *Form {
	return &Form{Method: "POST"}
}

func (form *Form) AddField(field *Field) {
	form.Fields = append(form.Fields, field)
}

func (form *Form) Render() string {
	var fieldsHtml []string
	for _, field := range form.Fields {
		label := fmt.Sprintf(`
		  <label for="%s" class="block text-sm font-medium leading-6 text-gray-900">%s</label>
		`, field.Input.Name, field.Input.Name)
		var required string
		if field.Input.Required {
			required = "required"
		}

		formField := fmt.Sprintf(`
	 <div>
	 		%s
		  <div class="mt-2" style="margin-bottom:1em">
		    <input type="%s" name="%s" class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" placeholder="%s" value="%s" %s>
      <p class="mt-2 text-sm text-red-600" style="padding:0;margin:0">%s</p>
		  </div>
		</div>`,
			label, field.Input.Type, field.Input.Name, field.Input.Placeholder, field.Input.Value, required, field.Input.Error)

		fieldsHtml = append(fieldsHtml, formField)
	}

	formHtml := fmt.Sprintf(`
<form action="%s" method="%s" >
  <div class="space-y-12">
    <div class="border-b border-gray-900/10 pb-12">
      <h2 class="text-base font-semibold leading-7 text-gray-900">Profile</h2>
      <p class="mt-1 text-sm leading-6 text-gray-600">This information will be displayed publicly so be careful what you share.</p>

	  	%s
    </div>
 </div>
<div class="mt-6 flex items-center justify-end gap-x-6">
    <button type="button" class="text-sm font-semibold leading-6 text-gray-900">Cancel</button>
    <button type="submit" class="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Save</button>
</div>
</form>`,
		form.Action, form.Method, strings.Join(fieldsHtml, ""),
	)
	return formHtml
}
func (form *Form) Validate() bool {
	return false
}
