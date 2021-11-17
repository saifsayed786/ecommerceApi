package models

import "time"

type Product struct {
	Id          int    `json:"id,string"`
	Sku         string `json:"sku,"`
	Name        string `json:"Name" gorm:"unique"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	CreatedBy   int
	User        User      `gorm:"foreignKey:CreatedBy" json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type ProductCategories struct {
	Id         int `json:"id,string"`
	ProductId  int
	Product    Product `gorm:"foreignKey:ProductId"`
	CategoryId int
	Category   Category `gorm:"foreignKey:CategoryID"`
}
