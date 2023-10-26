package dto

type StoreRegister struct {
	StoreEmail string `json:"store_email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	StoreName  string `json:"store_name" binding:"required"`
	StoreType  string `json:"store_type" binding:"required"`
}

type StoreLogin struct {
	StoreEmail string `json:"store_email" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
