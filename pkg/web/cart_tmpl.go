package web

import (
	"fmt"
	"store/pkg/models"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type CartTmpl struct {
	router *mux.Router
}

func NewCartTmpl(router *mux.Router) *CartTmpl {
	return &CartTmpl{router: router}
}

func (t *CartTmpl) cart(cartVM CartVM) string {
	html := loadTemplate("cart.html")
	var cartItems string

	for _, item := range cartVM.Cart.CartItems {
		cartItems += t.cartItem(item) + "\n"

	}
	html = strings.Replace(html, "###cart_items###", cartItems, -1)
	html = strings.Replace(html, "###edit_link###", UrlInternal(t.router, CART_EDIT_ROUTE, "uuid", cartVM.Cart.Uuid).Value, -1)
	return html
}

func (t *CartTmpl) cartItem(item *models.CartItem) string {
	html := loadTemplate("cart_item.html")
	html = strings.Replace(html, "###name###", item.Product.Name, -1)
	html = strings.Replace(html, "###description###", item.Product.Description, -1)
	html = strings.Replace(html, "###price###", "$ "+strconv.Itoa(item.Subtotal), -1)
	var fields []string
	for _, field := range item.Product.Fields {
		fieldHtml := fmt.Sprintf(`<li class="mt-1 text-sm text-gray-500"><span style="font-weight:bold">%s</span>:<br /> %s</li>`, field.Title, field.Value)
		fields = append(fields, fieldHtml)
	}
	html = strings.Replace(html, "###product_fields###", strings.Join(fields, "\n"), -1)
	html = strings.Replace(html, "###remove_link###", UrlInternal(t.router, CART_ITEM_DELETE_ROUTE, "uuid", item.Uuid).Value, -1)
	return html
}
