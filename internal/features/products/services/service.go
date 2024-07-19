package services

import (
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/users"
	"errors"
)

type ProductServices struct {
	qry      products.PServices
	userRepo users.Query
}

func NewProductService(q products.PServices, userRepo users.Query) products.PServices {
	return &ProductServices{
		qry:      q,
		userRepo: userRepo,
	}
}

func (ps *ProductServices) AddProduct(newProduct products.Product) error {
	err := ps.qry.AddProduct(newProduct)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menambahkan product")
	}
	return nil
}

func (ps *ProductServices) GetProductByID(id uint) (products.Product, users.User, error) {
	// Mengambil data produk dari repository
	product, user, err := ps.qry.GetProductByID(id)
	if err != nil {
		return products.Product{}, users.User{}, err
	}
	// Pastikan user tidak nil, walaupun sudah diset di repository
	if user.ID == 0 {
		return products.Product{}, users.User{}, errors.New("user not found")
	}

	return product, user, nil
}

func (ps *ProductServices) UpdateProductByID(id uint, updatedProduct products.Product) error {
	err := ps.qry.UpdateProductByID(id, updatedProduct)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat memperbarui produk")
	}
	return nil
}

func (ps *ProductServices) DeleteProduct(id uint) error {
	err := ps.qry.DeleteProduct(id)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menghapus produk")
	}
	return nil
}

func (ps *ProductServices) GetAllProducts(term string, limit int, offset int) ([]products.Product, error) {
	products, err := ps.qry.GetAllProducts(term, limit, offset)
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server saat mengambil daftar product")
	}
	return products, nil
}

func (ps *ProductServices) GetProductsByUserID(userID uint, term string, limit int, offset int) ([]products.Product, error) {
	products, err := ps.qry.GetProductsByUserID(userID, term, limit, offset)
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server saat mengambil daftar produk")
	}
	return products, nil
}

func (ps *ProductServices) CountAllProducts(term string) (int64, error) {
	count, err := ps.qry.CountAllProducts(term)
	if err != nil {
		return 0, errors.New("terjadi kesalahan pada server saat menghitung total product")
	}
	return count, nil
}

func (ps *ProductServices) CountProductsByUserID(userID uint, term string) (int64, error) {
	count, err := ps.qry.CountProductsByUserID(userID, term)
	if err != nil {
		return 0, errors.New("terjadi kesalahan pada server saat menghitung total product")
	}
	return count, nil
}
