package service

import (
	"midtrans-go/helper"
	"midtrans-go/model/entity"
	"midtrans-go/model/web"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type MidtransServiceImpl struct {
	Validate *validator.Validate
	DB       *gorm.DB
}

func NewMidtransServiceImpl(validate *validator.Validate, db *gorm.DB) *MidtransServiceImpl {
	return &MidtransServiceImpl{
		Validate: validate,
		DB:       db,
	}
}

func (service *MidtransServiceImpl) Create(c *gin.Context, request web.MidtransRequest) web.MidtransResponse {
	// Validasi request
	err := service.Validate.Struct(request)
	if err != nil {
		helper.PanicIfError(err)
	}

	// Konfigurasi Snap Client
	var snapClient = snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// ID pengguna sebagai string
	userID := strconv.Itoa(request.UserId)

	// Detail transaksi untuk Snap
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-User-" + userID + "-" + request.ItemID,
			GrossAmt: request.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: request.ItemName,        // Sebaiknya data ini diambil dari request
			Email: "hafizhzik24@gmail.com", // Ganti dengan data dinamis jika tersedia
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property-" + request.ItemID,
				Qty:   1,
				Price: request.Amount,
				Name:  request.ItemName,
			},
		},
	}

	// Buat transaksi melalui Snap API
	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		helper.PanicIfError(errSnap.GetRawError())
	}

	// Simpan transaksi ke database
	transaction := entity.Transaction{
		OrderID:     req.TransactionDetails.OrderID,
		UserID:      request.UserId,
		ItemID:      request.ItemID,
		ItemName:    request.ItemName,
		Amount:      request.Amount,
		Token:       response.Token,
		RedirectUrl: response.RedirectURL,
	}
	err = service.DB.Create(&transaction).Error
	if err != nil {
		helper.PanicIfError(err)
	}

	// Simpan status pembayaran
	paymentStatus := entity.PaymentStatus{
		TransactionID: transaction.ID,
		Status:        "Pending",
	}
	err = service.DB.Create(&paymentStatus).Error
	if err != nil {
		helper.PanicIfError(err)
	}

	// Return response ke controller
	return web.MidtransResponse{
		Token:       response.Token,
		RedirectUrl: response.RedirectURL,
	}
}

func (service *MidtransServiceImpl) GetTransactionByOrderID(orderID string) (*entity.Transaction, error) {
	var transaction entity.Transaction
	err := service.DB.Where("order_id = ?", orderID).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (service *MidtransServiceImpl) GetPaymentStatusByTransactionID(transactionID uint) (*entity.PaymentStatus, error) {
	var paymentStatus entity.PaymentStatus
	err := service.DB.Where("transaction_id = ?", transactionID).First(&paymentStatus).Error
	if err != nil {
		return nil, err
	}
	return &paymentStatus, nil
}

func (service *MidtransServiceImpl) UpdatePaymentStatus(transactionID uint, status string) {
	var paymentStatus entity.PaymentStatus
	err := service.DB.Where("transaction_id = ?", transactionID).First(&paymentStatus).Error
	if err != nil {
		panic(err)
	}

	paymentStatus.Status = status
	err = service.DB.Save(&paymentStatus).Error
	if err != nil {
		panic(err)
	}
}
