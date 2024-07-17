package services

import (
	"TokoGadget/internal/features/products"
	"errors"
)

type ProductServices struct {
	qry products.PServices
}

func NewProductService(q products.PServices) products.PServices {
	return &ProductServices{
		qry: q,
	}
}

func (ps *ProductServices) AddProduct(newProduct products.Product) error {
	err := ps.qry.AddProduct(newProduct)
	if err != nil {
		return errors.New("terjadi kesalahan pada server saat menambahkan product")
	}
	return nil
}

func (ps *ProductServices) GetAllProducts() ([]products.Product, error) {
	products, err := ps.qry.GetAllProducts()
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server saat mengambil daftar product")
	}
	return products, nil
}

func (ps *ProductServices) GetAllProductsByUserID(userID uint) ([]products.Product, error) {
	products, err := ps.qry.GetAllProductsByUserID(userID)
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server saat mengambil daftar produk")
	}
	return products, nil
}

func (ps *ProductServices) GetProductByID(id uint) (*products.Product, error) {
	product, err := ps.qry.GetProductByID(id)
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server saat mengambil artikel")
	}
	return product, nil
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

func (ps *ProductServices) GetProductsBySearch(search string) ([]products.Product, error) {
	// Implementasi logika pencarian di sini
	products, err := ps.qry.GetProductsBySearch(search)
	if err != nil {
		return nil, err
	}
	return products, nil
}
