package task1

import (
	"fmt"
	"github.com/google/uuid"
)

/*
 Функции в этом файле реализовывать и менять не нужно, они лишь показывают принцип применения Товара и DTO
*/

func UpdateProduct(dto UpdateProductDto) error {
	if err := dto.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	product := FindProduct(dto.UUID)

	product.Name = dto.Name
	product.Description = dto.Description
	product.Price = dto.Price
	product.Status = dto.Status

	if err := SaveProduct(product); err != nil {
		return fmt.Errorf("save product error: %w", err)
	}

	return nil
}

func FindProduct(uuid uuid.UUID) Product {
	return Product{} // этот метод реализовывать не нужно
}

func SaveProduct(product Product) error {
	// этот метод реализовывать не нужно
	return nil
}
