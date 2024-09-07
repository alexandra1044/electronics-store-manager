package models

import (
	"github.com/alexandra1044/electronics-store-manager/pkg/config"
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
	db = config.GetDB
	db.AutoMigrate(&Product{})
}
