package models

import "github.com/google/uuid"

type Uuid struct {
	Uuid string
}

func NewUuid() Uuid {
	return Uuid{Uuid: uuid.New().String()}

}
