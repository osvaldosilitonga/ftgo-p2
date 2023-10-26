package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       int    `json:"price"`
	StoreID     uint   `json:"store_id"`
}
