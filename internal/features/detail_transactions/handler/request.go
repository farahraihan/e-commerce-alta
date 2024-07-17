package handler

type RequestCart struct {
	CartID   uint `json:"cart_id"`
	Quantity uint `json:"quantity"`
}
