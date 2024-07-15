package products

type Product struct {
	ID             uint
	UserID         uint
	ProductName    string
	Category       string
	Description    string
	Price          int64
	Stock          int32
	ProductPicture string
}
