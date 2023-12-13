package repository

import (
	"github.com/SyberiaEmperor/OkFridge/fridge_service_prototype/models"
	"gorm.io/gorm"
)

type ProductStorage interface {
	AddProduct(prod models.Product) error
	changeProduct(prod models.Product) error
}

type Repository struct {
	ProductStorage
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ProductStorage: NewProductStorageService(db),
	}
}
