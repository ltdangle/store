package models

import "github.com/google/uuid"

func NewUuid() string {
	return uuid.New().String()

}
