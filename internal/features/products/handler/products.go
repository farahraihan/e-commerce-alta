package handler

import (
	"TokoGadget/configs"
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/helper"
	"TokoGadget/internal/utils"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	srv products.PServices
	tu  utils.TokenUtilityInterface
}

func NewProductController(s products.PServices, t utils.TokenUtilityInterface) products.PHandler {
	return &ProductController{
		srv: s,
		tu:  t,
	}
}

func (pc *ProductController) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := pc.tu.DecodeToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormatNonData(http.StatusUnauthorized, "Unauthorized", "error"))
		}
		fmt.Println("Ini Depe User Gaes :", userID)

		// Get image from form data
		image, err := c.FormFile("product_picture")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid image file", nil, nil))
		}

		// Open the image file
		src, err := image.Open()
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat("failed", http.StatusInternalServerError, "Failed to open image file", nil, nil))
		}
		defer src.Close()

		// Upload image to Cloudinary
		cld, err := configs.ConnectCloudinary()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat("failed", http.StatusInternalServerError, "Cloudinary configuration error", nil, nil))
		}
		fmt.Println("URL_IMAGE:", src)
		uploadResult, err := cld.Upload.Upload(c.Request().Context(), src, uploader.UploadParams{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat("failed", http.StatusInternalServerError, "Failed to upload image", nil, nil))
		}

		imageURL := uploadResult.SecureURL

		// Bind the request
		var req CreateOrUpdateProductRequest
		err = c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid request body", nil, nil))
		}

		// Convert request to product model
		newProduct := products.Product{
			UserID:         userID,
			ProductName:    req.ProductName,
			Category:       req.Category,
			Description:    req.Description,
			Price:          req.Price,
			Stock:          req.Stock,
			ProductPicture: imageURL,
		}

		// Add product to database
		err = pc.srv.AddProduct(newProduct)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to add product", nil, nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat("success", http.StatusOK, "Product was successfully created", nil, nil))
	}
}

func (pc *ProductController) GetAllProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		// Membaca parameter dari query string
		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		allStr := c.QueryParam("all")
		all, _ := strconv.ParseBool(allStr)

		search := c.QueryParam("search")

		limit := 10
		offset := (page - 1) * limit

		// Panggil layanan untuk mengambil produk
		var products []products.Product
		var totalItems int64
		if all {
			products, err = pc.srv.GetAllProducts(search, limit, offset)
			if err != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to retrieve product data", nil, nil))
			}
			totalItems, err = pc.srv.CountAllProducts(search)
		} else {
			userID := pc.tu.DecodeToken(c.Get("user").(*jwt.Token))
			if userID == 0 {
				return c.JSON(http.StatusUnauthorized, helper.ResponseFormatNonData(http.StatusUnauthorized, "Unauthorized", "error"))
			}

			products, err = pc.srv.GetProductsByUserID(userID, search, limit, offset)
			if err != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to retrieve product data", nil, nil))
			}
			totalItems, err = pc.srv.CountProductsByUserID(userID, search)
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to retrieve total product count", nil, nil))
		}

		// Konversi ke response format yang diinginkan
		response := ToResponseProducts(products)

		totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

		meta := map[string]interface{}{
			"totalItems":   totalItems,
			"itemsPerPage": limit,
			"currentPage":  page,
			"totalPages":   totalPages,
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
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid product ID", nil, nil))
		}

		product, user, err := pc.srv.GetProductByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to retrieve product data", nil, nil))
		}

		response := map[string]interface{}{
			"product":    ToResponseProduct(product),
			"sellerID":   user.ID,
			"sellerName": user.Fullname,
		}

		return c.JSON(http.StatusOK, helper.ResponseFormatWithMeta("success", http.StatusOK, "Product data fetched successfully", response, nil))
	}
}

func (pc *ProductController) UpdateProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("product_id")

		id, err := strconv.ParseUint(productID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid product ID", nil, nil))
		}

		userID := pc.tu.DecodeToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormatNonData(http.StatusUnauthorized, "Unauthorized", "error"))
		}

		var req CreateOrUpdateProductRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid request body", nil, nil))
		}

		updatedProduct := ToModelProduct(req, userID)
		err = pc.srv.UpdateProductByID(uint(id), updatedProduct)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to update the product", nil, nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat("success", http.StatusOK, "Successfully updated the product", nil, nil))
	}
}

func (pc *ProductController) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID := c.Param("product_id")

		id, err := strconv.ParseUint(productID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Invalid product ID", nil, nil))
		}

		userID := pc.tu.DecodeToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormatNonData(http.StatusUnauthorized, "Unauthorized", "error"))
		}

		err = pc.srv.DeleteProduct(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat("failed", http.StatusBadRequest, "Failed to delete the product", nil, nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat("success", http.StatusOK, "Successfully deleted the product", nil, nil))
	}
}
