package handler

import (
	"TokoGadget/internal/features/sales"
	"TokoGadget/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SaleController struct {
	srv sales.SServices
	tu  utils.TokenUtilityInterface
}

func NewSaleController(s sales.SServices, t utils.TokenUtilityInterface) sales.SHandler {
	return &SaleController{
		srv: s,
		tu:  t,
	}
}

func (sc *SaleController) GetSalesByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid UserID")
		}

		_, products, transactions, details, err := sc.srv.GetSalesByUserID(uint(userID))
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fetch sales data")
		}

		// Mapping data to the desired response format
		var sales []map[string]interface{}

		for _, transaction := range transactions {
			salesItems := make([]map[string]interface{}, 0)

			for _, detail := range details {
				if detail.TransactionID == transaction.ID {
					product := products[detail.ProductID]
					item := map[string]interface{}{
						"product_name":    product.ProductName,
						"product_picture": product.ProductPicture,
						"quantity":        detail.Quantity,
						"sub_total":       product.Price * uint64(detail.Quantity),
					}
					salesItems = append(salesItems, item)
				}
			}

			sale := map[string]interface{}{
				"transaction_id": transaction.ID,
				"buyer_id":       transaction.UserID,
				"status":         transaction.Status,
				"sales_items":    salesItems,
				"grand_total":    transaction.GrandTotal,
				"created_at":     transaction.CreatedAt.Format("2006-01-02T15:04:05Z"),
				"updated_at":     transaction.UpdatedAt.Format("2006-01-02T15:04:05Z"),
				"deleted_at":     transaction.DeletedAt,
			}

			sales = append(sales, sale)
		}

		response := map[string]interface{}{
			"data": sales,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (sc *SaleController) GetSalesByTransactionID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid UserID")
		}

		transactionID, err := strconv.ParseUint(c.Param("transactionID"), 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid TransactionID")
		}

		_, products, transaction, details, err := sc.srv.GetSalesByTransactionID(uint(userID), uint(transactionID))
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fetch sales data")
		}

		salesItems := make([]map[string]interface{}, 0)

		for _, detail := range details {
			product := products[detail.ProductID]
			item := map[string]interface{}{
				"product_name":    product.ProductName,
				"product_picture": product.ProductPicture,
				"quantity":        detail.Quantity,
				"sub_total":       product.Price * uint64(detail.Quantity),
			}
			salesItems = append(salesItems, item)
		}

		response := map[string]interface{}{
			"transaction_id": transaction.ID,
			"buyer_id":       transaction.UserID,
			"status":         transaction.Status,
			"sales_items":    salesItems,
			"grand_total":    transaction.GrandTotal,
			"created_at":     transaction.CreatedAt.Format("2006-01-02T15:04:05Z"),
			"updated_at":     transaction.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			"deleted_at":     transaction.DeletedAt,
		}

		return c.JSON(http.StatusOK, response)
	}
}
