package dbservice

import (
	"github.com/bondhan/gotem/persistence/repository"
)

// InventoryService ...
type InventoryService interface {
}

type inventoryService struct {
	productRepo *repository.ProductsRepository
	variantRepo *repository.VariantsRepository
}

// NewInventoryService ..
func NewInventoryService(variantRepo *repository.VariantsRepository, productRepo *repository.ProductsRepository) InventoryService {
	return &inventoryService{
		variantRepo: variantRepo,
		productRepo: productRepo,
	}
}
