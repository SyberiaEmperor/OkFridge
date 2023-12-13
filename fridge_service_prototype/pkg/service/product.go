package service

import (
	"github.com/SyberiaEmperor/OkFridge/fridge_service_prototype/models"
	"github.com/SyberiaEmperor/OkFridge/fridge_service_prototype/pkg/repository"
)

type ProductStorageService struct {
	repo repository.ProductStorage
}

func NewProductStorageService(repo repository.ProductStorage) *ProductStorageService {
	return &ProductStorageService{repo: repo}
}

func (prods *ProductStorageService) AddProduct(prod models.Product) error {
	return prods.repo.AddProduct(prod)
}

func (prods *ProductStorageService) ChangeProduct(prod models.Product) error {
	return prods.repo.ChangeProduct(prod)
}
