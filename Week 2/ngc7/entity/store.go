package entity

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey"`
	StoreEmail string `json:"store_email" gorm:"not null; unique"`
	Password   string `json:"password" gorm:"not null"`
	StoreName  string `json:"store_name" gorm:"not null"`
	StoreType  string `json:"store_type" gorm:"not null"`
	Products   []Product
}
