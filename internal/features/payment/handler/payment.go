package payment

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentHandler struct {
	MidtransClient snap.Client
}

func NewPaymentHandler(midtransKey string) *PaymentHandler {
	client := snap.Client{}
	client.New(midtransKey, midtrans.Sandbox)

	return &PaymentHandler{
		MidtransClient: client,
	}
}

func (ph *PaymentHandler) CreateSnapTransaction(c echo.Context) error {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "ID-12345",
			GrossAmt: 100000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
		},
	}

	snapResp, err := ph.MidtransClient.CreateTransaction(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error creating Snap transaction: %v", err),
		})
	}

	resp := struct {
		RedirectURL string `json:"redirect_url"`
		Token       string `json:"token"`
	}{
		RedirectURL: snapResp.RedirectURL,
		Token:       snapResp.Token,
	}

	return c.JSON(http.StatusOK, resp)
}
