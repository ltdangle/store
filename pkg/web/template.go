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
)

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
func template(cartVM CartVM) string {
	html := loadTemplate("template.html")
	html = strings.Replace(html, "###cart###", cart(cartVM), 1)

	return html
}

func cart(cartVM CartVM) string {
	cart := loadTemplate("cart.html")
	var cartItems string

	for _, item := range cartVM.Cart.CartItems {
		cartItems += cartItem(item) + "\n"

	}
	cart = strings.Replace(cart, "###cart_items###", cartItems, -1)

	return cart
}

func cartItem(item *models.CartItem) string {
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
	return html
}
