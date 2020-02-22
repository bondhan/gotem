package dbservice

import (
	"github.com/bondhan/gotem/persistence/repository"
)

type InventoryService interface {
}

type inventoryService struct {
	buyerRepo   *repository.BuyerRepository
	productRepo *repository.ProductsRepository
}

func NewInventoryService(buyerRepo *repository.BuyerRepository, productRepo *repository.ProductsRepository) InventoryService {
	return &inventoryService{
		buyerRepo:   buyerRepo,
		productRepo: productRepo,
	}
}
