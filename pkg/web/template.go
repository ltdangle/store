package web

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

)

type Link struct {
	url  string
	text string
}

type Tmpl struct {
	router       *AppRouter
	main         string
	leftNavLinks []Link
}

func NewTmpl(router *AppRouter) *Tmpl {
	t := &Tmpl{router: router}
	// TODO: fix
	url := router.UrlInternal( CART_VIEW_ROUTE, "uuid", "someuuid").Value
	t.AddNavLink(url, "Cart")
	t.AddNavLink(url, "CartItem")
	t.AddNavLink(url, "Product")
	t.AddNavLink(url, "Product fields")
	return t
}

// TODO: use "html/template"
func LoadTemplate(tmpl string) string {
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

func (t *Tmpl) AddNavLink(url string, text string) {
	link := Link{url: url, text: text}
	t.leftNavLinks = append(t.leftNavLinks, link)
}

func (t *Tmpl) buildLeftNav() string {
	var links []string
	for _, link := range t.leftNavLinks {
		html := fmt.Sprintf(`
								<li>
									<a href="%s" class="text-gray-400 hover:text-white hover:bg-gray-800 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold">
									   %s	
									</a>
								</li>
            `,
			link.url, link.text,
		)
		links = append(links, html)
	}
	return strings.Join(links, "\n")
}

func (t *Tmpl) SetMain(html string) {
	t.main = html
}
func (t *Tmpl) Render() string {

	html := LoadTemplate("template.html")
	html = strings.Replace(html, "###left-nav###", t.buildLeftNav(), 1)
	html = strings.Replace(html, "###cart###", t.main, 1)

	return html
}
