package models

import (
	"alexandra1044/electronics-store-manager/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Name     string `gorm:""json:"name"`
	Price    int    `gorm:""json:"price"`
	Category string `gorm:""json:"category"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (p *Product) CreateProduct() *Product {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func getAllProducts() []Product {
	var Product []Product
	db.Find(&Product)
	return Product
}

func GetProductById(Id int64) (*Product, *gorm.DB) {
	var getProduct Product
	db := db.Where("ID=?id", Id).Find(&getProduct)
	return &getProduct, db
}

func DeleteProduct(Id int64) Product {
	var product Product
	db.Where("ID=?id", Id).Delete(product)
	return product
}
