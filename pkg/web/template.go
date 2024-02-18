package web

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
)

type Tmpl struct {
	router *mux.Router
	main   string
}

func NewTmpl(router *mux.Router) *Tmpl {
	return &Tmpl{router: router}
}

// TODO: use "html/template"
func loadTemplate(tmpl string) string {
	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}

	currentDir := filepath.Dir(currentFilePath)

	targetFilePath := filepath.Join(currentDir+"/tmpl/", tmpl)
	content, err := os.ReadFile(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
func (t *Tmpl) setMain(html string) {
	t.main = html
}
func (t *Tmpl) render() string {
	html := loadTemplate("template.html")
	html = strings.Replace(html, "###cart###", t.main, 1)

	return html
}
