package web

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	cartItem := loadTemplate("cart_item.html")
	cart = strings.Replace(cart, "###cart_item###", cartItem, -1)

	return cart
}
