package handler

import (
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/helper"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	srv products.PServices
}

func NewProductController(s products.PServices) products.PHandler {
	return &ProductController{
		srv: s,
	}
}

func (pc *ProductController) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := getUserIDFromToken(c) // Get user ID from token
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat("failed", http.StatusUnauthorized, "Unauthorized access", nil))
		}

		var req CreateOrUpdateProductRequest
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid request body", nil))
		}

		newProduct := ToModelProduct(req, userID)
		err = pc.srv.AddProduct(newProduct)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to add product", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat("success", http.StatusOK, "Product was successfully created", nil))
	}
}

func (pc *ProductController) GetAllProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		// Membaca parameter dari query string
		pageStr := c.QueryParam("page")
		page, _ := strconv.Atoi(pageStr) // konversi ke integer, default: 0

		allStr := c.QueryParam("all")
		all, _ := strconv.ParseBool(allStr) // konversi ke boolean, default: false

		search := c.QueryParam("search") // tambahan untuk pencarian

		userID := getUserIDFromToken(c) // Get user ID from token
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat("failed", http.StatusUnauthorized, "Unauthorized access", nil))
		}

		// Panggil layanan untuk mengambil produk
		var products []products.Product
		if all {
			products, err = pc.srv.GetAllProducts()
		} else {
			// Jika tidak meminta semua produk, gunakan pencarian
			products, err = pc.srv.GetProductsBySearch(search)
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to retrieve product data", nil))
		}

		// Konversi ke response format yang diinginkan
		response := ToResponseProducts(products)

		// Menambahkan metadata, misalnya jumlah total item, halaman saat ini, dll.
		meta := map[string]interface{}{
			"totalItems":   len(products),
			"itemsPerPage": 10, // misalnya 10 item per halaman
			"currentPage":  page + 1,
			// totalPages bisa dihitung berdasarkan totalItems dan itemsPerPage
		}

		// Mengembalikan response JSON dengan meta data
		return c.JSON(http.StatusOK, helper.ResponseFormatWithMeta("success", http.StatusOK, "Successfully retrieved all products data", response, meta))
	}
}

func (pc *ProductController) GetProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("product_id")

		id, err := strconv.ParseUint(productID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid product ID", nil))
		}

		userID := getUserIDFromToken(c) // Get user ID from token
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat("failed", http.StatusUnauthorized, "Unauthorized access", nil))
		}

		product, err := pc.srv.GetProductByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to retrieve product data", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat("success", http.StatusOK, "Product data fetched successfully", ToResponseProduct(*product)))
	}
}

func (pc *ProductController) UpdateProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("product_id")

		id, err := strconv.ParseUint(productID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid product ID", nil))
		}

		userID := getUserIDFromToken(c) // Get user ID from token
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat("failed", http.StatusUnauthorized, "Unauthorized access", nil))
		}

		var req CreateOrUpdateProductRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid request body", nil))
		}

		updatedProduct := ToModelProduct(req, userID)
		err = pc.srv.UpdateProductByID(uint(id), updatedProduct)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to update the product", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat("success", http.StatusOK, "Successfully updated the product", nil))
	}
}

func (pc *ProductController) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("product_id")

		id, err := strconv.ParseUint(productID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid product ID", nil))
		}

		userID := getUserIDFromToken(c) // Get user ID from token
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat("failed", http.StatusUnauthorized, "Unauthorized access", nil))
		}

		err = pc.srv.DeleteProduct(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to delete the product", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat("success", http.StatusOK, "Successfully deleted the product", nil))
	}
}

func getUserIDFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))
	return userID
}
