package controller

import (
	"midtrans-go/helper"
	"midtrans-go/model/entity"
	"midtrans-go/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserServiceImpl
}

func NewUserControllerImpl(userService service.UserServiceImpl) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) CreateUser(c *gin.Context) {
	var request entity.User
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.PanicIfError(err)
	}

	user := controller.UserService.CreateUser(request)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    user,
	})

}

func (controller *UserControllerImpl) GetUserWithTransaction(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	user, err := controller.UserService.GetUserWithTransaction(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User retrieved successfully",
		"data":    user,
	})
}
