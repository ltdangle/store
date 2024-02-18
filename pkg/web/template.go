package web

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"store/pkg/models"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Tmpl struct {
	router *mux.Router
}

func NewTmpl(router *mux.Router) *Tmpl {
	return &Tmpl{router: router}
}

// TODO: use "html/template"
func (t *Tmpl) loadTemplate(tmpl string) string {
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
func (t *Tmpl) template(cartVM CartVM) string {
	html := t.loadTemplate("template.html")
	html = strings.Replace(html, "###cart###", t.cart(cartVM), 1)

	return html
}

func (t *Tmpl) cart(cartVM CartVM) string {
	cart := t.loadTemplate("cart.html")
	var cartItems string

	for _, item := range cartVM.Cart.CartItems {
		cartItems += t.cartItem(item) + "\n"

	}
	cart = strings.Replace(cart, "###cart_items###", cartItems, -1)

	return cart
}

func (t *Tmpl) cartItem(item *models.CartItem) string {
	html := t.loadTemplate("cart_item.html")
	html = strings.Replace(html, "###name###", item.Product.Name, -1)
	html = strings.Replace(html, "###description###", item.Product.Description, -1)
	html = strings.Replace(html, "###price###", "$ "+strconv.Itoa(item.Subtotal), -1)
	var fields []string
	for _, field := range item.Product.Fields {
		fieldHtml := fmt.Sprintf(`<li class="mt-1 text-sm text-gray-500"><span style="font-weight:bold">%s</span>:<br /> %s</li>`, field.Title, field.Value)
		fields = append(fields, fieldHtml)
	}
	html = strings.Replace(html, "###product_fields###", strings.Join(fields, "\n"), -1)
	return html
}
