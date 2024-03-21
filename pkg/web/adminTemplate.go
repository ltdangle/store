package web

import (
	"bytes"
	"context"
	"store/pkg/repo"
)

type Link struct {
	Url  string
	Text string
}

type AdminTmpl struct {
	router       *AppRouter
	main         string
	leftNavLinks []Link
}

func NewAdminTmpl(router *AppRouter) *AdminTmpl {
	t := &AdminTmpl{router: router}
	return t
}

type Cart struct {
	Uuid      string
	CartItems []CartItem
}
type CartItem struct {
	Uuid     string
	Product  Product
	Subtotal int
}

type Product struct {
	Name        string
	Description string
	Fields      []ProductField
}
type ProductField struct {
	Title string
	Value string
}
type VM struct {
	Cart Cart
}

func (t *AdminTmpl) cart(cartVM *repo.CartVM) string {

	var h bytes.Buffer
	templ := cart(cartVM, t.router)
	_ = templ.Render(context.Background(), &h)

	return h.String()
}

func (t *AdminTmpl) AddNavLink(url string, text string) {
	link := Link{Url: url, Text: text}
	t.leftNavLinks = append(t.leftNavLinks, link)
}

func (t *AdminTmpl) SetMain(html string) {
	t.main = html
}
func (t *AdminTmpl) Render() string {
	templ := Admintmpl(t.leftNavLinks, t.main)

	var html bytes.Buffer
	_ = templ.Render(context.Background(), &html)
	return html.String()
}
