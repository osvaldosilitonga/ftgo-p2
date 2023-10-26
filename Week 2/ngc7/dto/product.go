package dto

type ProductPost struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"image_url" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	StoreID     uint   `json:"store_id" binding:"required"`
}
