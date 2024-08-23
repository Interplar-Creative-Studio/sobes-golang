package task3

import "github.com/google/uuid"

type Product struct {
	UUID        uuid.UUID
	Title       string
	SKU         string
	Description string
	Price       float64
	Status      string
}
