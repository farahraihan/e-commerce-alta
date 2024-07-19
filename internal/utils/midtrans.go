package utils

import (
	"TokoGadget/configs"
	"fmt"

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
	s := snap.Client{}
    s.New(configs.ImportserverKey(), midtrans.Sandbox)
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(amount), // Amount in Rupiah (IDR)
		},
	}
	fmt.Println("Requesnya :", req.TransactionDetails)

    // res, err :=mp.snapClient.CreateTransaction(req)
    res, err :=s.CreateTransaction(req)
    fmt.Println("isi Respon midtrans: ", res)
	// if err != nil {
    //     log.Println("midtrans error:", err.Error())
    //     return "", nil
    // }
	fmt.Println("Erornya atuh :",err)
	fmt.Println("Urlnya :",res.RedirectURL)
	urlMidtrans := res.RedirectURL

    return urlMidtrans, err
}