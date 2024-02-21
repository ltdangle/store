package form

import (
	"fmt"
	"testing"
)

func TestForm(_ *testing.T) {
	form := &Form{}
	form.AddField(&Field{Type: "text", Value: "text value", Required: true})
	form.AddField(&Field{Type: "int", Value: "42", Required: true})
	form.AddField(&Field{Type: "date", Value: "12/12/1987"})
	fmt.Println(form.Render())

}
