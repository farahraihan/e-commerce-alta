package repository

import (
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/users"

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

func (pm *ProductModel) GetProductByID(id uint) (products.Product, users.User, error) {
	var product products.Product
	var user users.User

	err := pm.db.Preload("User").Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, user, err
	}

	return product, product.User, nil
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

func (pm *ProductModel) GetAllProducts(term string, limit int, offset int) ([]products.Product, error) {
	var products []products.Product
	query := pm.db.Model(&products).Limit(limit).Offset(offset)

	if term != "" {
		query = query.Where("product_name LIKE ? OR category LIKE ?", "%"+term+"%", "%"+term+"%")
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pm *ProductModel) GetProductsByUserID(userID uint, term string, limit int, offset int) ([]products.Product, error) {
	var products []products.Product
	query := pm.db.Model(&products).Where("user_id = ?", userID).Limit(limit).Offset(offset)

	// Jika term tidak kosong, gunakan untuk pencarian
	if term != "" {
		query = query.Where("(product_name LIKE ? OR category LIKE ?)", "%"+term+"%", "%"+term+"%")
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pm *ProductModel) CountAllProducts(term string) (int64, error) {
	var count int64
	err := pm.db.Model(&products.Product{}).Where("product_name LIKE ? OR category LIKE ?", "%"+term+"%", "%"+term+"%").Count(&count).Error
	return count, err
}

func (pm *ProductModel) CountProductsByUserID(userID uint, term string) (int64, error) {
	var count int64
	err := pm.db.Model(&products.Product{}).Where("(product_name LIKE ? OR category LIKE ?) AND user_id = ?", "%"+term+"%", "%"+term+"%", userID).Count(&count).Error
	return count, err
}
