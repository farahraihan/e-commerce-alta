package utils

import (
    "log"

    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/snap"
)

type MidtransInterface interface {
    RequestPayment(orderId string, amount int) (string, error)
}

type midtransPayment struct {
    snapClient snap.Client
}

func NewMidtransPayment(serverKey string) MidtransInterface {
    return &midtransPayment{
        snapClient: snap.Client{ServerKey: serverKey, Env: midtrans.Sandbox},
    }
}

func (mp *midtransPayment) RequestPayment(orderId string, amount int) (string, error) {
    req := snap.Request{
        TransactionDetails: midtrans.TransactionDetails{
            OrderID:  orderId,
            GrossAmt: int64(amount),
        },
    }

    res, err := mp.snapClient.CreateTransaction(&req)
    if err != nil {
        log.Println("midtrans error:", err.Error())
        return "", nil
    }

    return res.RedirectURL, nil
}