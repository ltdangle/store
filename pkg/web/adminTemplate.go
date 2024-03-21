package web

import (
	"bytes"
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"store/pkg/repo"
	"strconv"
	"strings"
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
	html := LoadTemplate("cart.html")
	var cartItems string

	for _, item := range cartVM.CartItems {
		cartItems += t.cartItem(item) + "\n"

	}
	html = strings.Replace(html, "###cart_items###", cartItems, -1)
	html = strings.Replace(html, "###edit_link###", t.router.UrlInternal(ADMIN_VIEW_ENTITY_ROUTE, "entity", "cart", "uuid", cartVM.Cart.Uuid), -1)
	return html
}

func (t *AdminTmpl) cartItem(item repo.CartItemVM) string {
	html := LoadTemplate("cart_item.html")
	html = strings.Replace(html, "###name###", item.Product.Name, -1)
	html = strings.Replace(html, "###description###", item.Product.Description, -1)
	html = strings.Replace(html, "###price###", "$ "+strconv.Itoa(item.CartItem.Subtotal), -1)
	var fields []string
	html = strings.Replace(html, "###product_fields###", strings.Join(fields, "\n"), -1)
	html = strings.Replace(html, "###remove_link###", t.router.UrlInternal(CART_ITEM_DELETE_ROUTE, "uuid", item.CartItem.Uuid), -1)
	return html
}

// TODO: use a-ha/templ
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
