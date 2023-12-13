package repository

import (
	"github.com/SyberiaEmperor/OkFridge/fridge_service_prototype/models"
	"gorm.io/gorm"
)

type ProductStorageService struct {
	db *gorm.DB
}

func NewProductStorageService(db *gorm.DB) *ProductStorageService {
	return &ProductStorageService{
		db: db,
	}
}

func (prods *ProductStorageService) AddProduct(prod models.Product) error {
	res := prods.db.Save(&prod)

	return res.Error
}

func (prods *ProductStorageService) changeProduct(prod models.Product) error {
	var product models.Product
	res := prods.db.First(&product, prod.Id)

	return res.Error
}
