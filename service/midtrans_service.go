package service

import (
	"midtrans-go/model/entity"
	"midtrans-go/model/web"

	"github.com/gin-gonic/gin"
)

type MidtransService interface {
	Create(c *gin.Context, request web.MidtransRequest) web.MidtransResponse
	GetTransactionByOrderID(orderID string) (*entity.Transaction, error)
	GetPaymentStatusByTransactionID(transactionID uint) (*entity.PaymentStatus, error)
	UpdatePaymentStatus(transactionID uint, status string)
}
