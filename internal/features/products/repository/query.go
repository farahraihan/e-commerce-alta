package repository

import (
	"TokoGadget/internal/features/products"

	"gorm.io/gorm"
)

type ProductModel struct {
	db *gorm.DB
}

func NewProductModel(connection *gorm.DB) products.PQuery {
	return &ProductModel{
		db: connection,
	}
}

func (pm *ProductModel) AddProduct(newProduct products.Product) error {
	return pm.db.Create(&newProduct).Error
}

func (pm *ProductModel) GetAllProducts() ([]products.Product, error) {
	var productList []products.Product
	err := pm.db.Where("deleted_at IS NULL").Find(&productList).Error
	return productList, err
}

func (pm *ProductModel) GetAllProductsByUserID(userID uint) ([]products.Product, error) {
	var productList []products.Product
	err := pm.db.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&productList).Error
	return productList, err
}

func (pm *ProductModel) GetProductByID(id uint) (*products.Product, error) {
	var article products.Product
	err := pm.db.Where("id = ? AND deleted_at IS NULL", id).First(&article).Error
	return &article, err
}

func (pm *ProductModel) DeleteProduct(id uint) error {
	var product products.Product
	err := pm.db.First(&product, id).Error
	if err != nil {
		return err
	}
	return pm.db.Delete(&product).Error
}

func (pm *ProductModel) UpdateProductByID(productID uint, updatedProduct products.Product) error {
	var product products.Product
	err := pm.db.Where("id = ? AND deleted_at IS NULL", productID).First(&product).Error
	if err != nil {
		return err
	}
	updatedProduct.ID = productID
	return pm.db.Save(&updatedProduct).Error
}

func (pm *ProductModel) GetProductsBySearch(search string) ([]products.Product, error) {
	var productList []products.Product
	// Implementasi logika pencarian berdasarkan kategori atau nama produk di sini
	err := pm.db.Where("category LIKE ? OR product_name LIKE ?", "%"+search+"%", "%"+search+"%").Find(&productList).Error
	return productList, err
}
