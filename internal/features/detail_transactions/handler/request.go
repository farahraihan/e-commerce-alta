package handler

type RequestCart struct {
	ProductID uint `json:"product_id"`
	CartID    uint `json:"cart_id"`
	Quantity  uint `json:"quantity"`
}
