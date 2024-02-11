package main

import (
	"store/pkg/repo"
)

func main() {
	repo.Migrate(".env")
}
