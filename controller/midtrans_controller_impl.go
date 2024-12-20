package controller

import (
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"midtrans-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MidtransControllerImpl struct {
	MidtransService service.MidtransService
}

func NewMidtransControllerImpl(midtransService service.MidtransService) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		MidtransService: midtransService,
	}
}

func (controller *MidtransControllerImpl) Create(c *gin.Context) {
	var request web.MidtransRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.PanicIfError(err)
	}

	midtransResponse := controller.MidtransService.Create(c, request)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   midtransResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *MidtransControllerImpl) CheckStatus(c *gin.Context) {
	orderID := c.Param("order_id")

	// Dapatkan transaksi berdasarkan orderID
	transaction, err := controller.MidtransService.GetTransactionByOrderID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
		return
	}

	paymentStatus, err := controller.MidtransService.GetPaymentStatusByTransactionID(transaction.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Payment status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orderID": orderID,
		"status":  paymentStatus.Status,
	})
}
