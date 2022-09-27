package repository

import (
	"context"
	"donasi/model"
	"fmt"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	// "github.com/midtrans/midtrans-go/snap"
)

func (dbRepo *databaseRepository) PostTranscationDonasi(ctx context.Context, payload *model.PayloadTransaksiRequest) (result *snap.Response, err error) {
	// dbRepo.confMidtrans.Snap.c
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "YOUR-ORDER-ID-12345",
			GrossAmt: 100000,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: payload.Nama,
			Email: payload.Email,
			Phone: payload.NoHp,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}
	snapRes, _ := dbRepo.confMidtrans.Snap.CreateTransaction(req)
	if snapRes.ErrorMessages != nil {
		return nil, err
	}
	return snapRes, nil
}

func (dbRepo *databaseRepository) StoreTranscationDonasi(ctx context.Context, payload *model.PayloadTransaksiRequest, donasiId string) (string, error) {

	uuid := uuid.New()
	query := `INSERT INTO transaksi (
		id,
    nama_donatur,
    email_donatur,
    nomor_telepon_donatur,
    nominal,
    donasi_id,
    doa,
    anomin) values 
		($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := dbRepo.confDb.DB.Exec(query, uuid, payload.Nama, payload.Email, payload.NoHp, payload.Nominal, payload.DonasiID, payload.Doa, payload.Status)
	if err != nil {
		return "", err
	}
	// err := pu.Conn.Exec(sqlStatement, a.ID, a.FirstName, a.LastName, a.Email, a.CreatedAt, a.UpdatedAt)
	fmt.Println(query)
	return "", nil
}

func (dbRepo *databaseRepository) NotificationTranscationDonasi()

// func GenerateSnapReq() *snap.Request {
// 	// Initiate Customer address
// 	custAddress := &midtrans.CustomerAddress{
// 		FName: "John",
// 		LName: "Doe",
// 		Phone: "081234567890",
// 		Address: "Baker Street 97th",
// 		City: "Jakarta",
// 		Postcode: "16000",
// 		CountryCode: "IDN",
// 	}

// 	// Initiate Snap Request
// 	snapReq := &snap.Request{
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID: "YOUR-UNIQUE-ORDER-ID-1234",
// 			GrossAmt: 200000,
// 		},
// 		CreditCard: &snap.CreditCardDetails{
// 			Secure: true,
// 		},
// 		CustomerDetail: &midtrans.CustomerDetails{
// 			FName: "John",
// 			LName: "Doe",
// 			Email: "john@doe.com",
// 			Phone: "081234567890",
// 			BillAddr: custAddress,
// 			ShipAddr: custAddress,
// 		},
// 		Items: &[]midtrans.ItemDetails{
// 			midtrans.ItemDetails{
// 				ID: "ITEM1",
// 				Price: 200000,
// 				Qty: 1,
// 				Name: "Someitem",
// 			},
// 		},
// 	}

//  return snapReq
// }
