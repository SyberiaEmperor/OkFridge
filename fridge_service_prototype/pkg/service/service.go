package service

import (
	"github.com/SyberiaEmperor/OkFridge/fridge_service_prototype/models"
	"github.com/SyberiaEmperor/OkFridge/fridge_service_prototype/pkg/repository"
)

type ProductStorage interface {
	AddProduct(prod models.Product) error
	changeProduct(prod models.Product) error
}

type Service struct {
	ProductStorage
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ProductStorage: NewProductStorageService(repo.ProductStorage),
	}
}