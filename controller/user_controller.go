package controller

import (
	"midtrans-go/model/entity"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(c *gin.Context)
	GetUserWithTransactions(UserID uint) (*entity.User, error)
}
