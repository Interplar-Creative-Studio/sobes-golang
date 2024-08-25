package task1

import (
	"fmt"
	"github.com/google/uuid"
	"tasks/task1/product_status"
	"time"
)

type Product struct {
	UUID        uuid.UUID
	Name        string
	Description string
	Price       float64
	Status      string // FIXME со строчкой работать неудобно
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p Product) ReadableStatus() string {
	return product_status.ReadableString(p.Status)
}

type UpdateProductDto struct {
	UUID        uuid.UUID
	Name        string
	Description string
	Price       float64
	Status      string
}

func (dto UpdateProductDto) Validate() error {
	if dto.Name == "" {
		return fmt.Errorf("product name is required")
	}
	if dto.Description == "" {
		return fmt.Errorf("product description is required")
	}
	if dto.Price <= 0 {
		return fmt.Errorf("product price is required and must be greater than zero")
	}

	// FIXME хотелось бы иметь dto.Status.Valid() bool
	if !product_status.Exists(dto.Status) {
		return fmt.Errorf("product status is unknown")
	}
	return nil
}
