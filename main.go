package main

import (
	"midtrans-go/controller"
	"midtrans-go/initializer"
	"midtrans-go/middleware"
	"midtrans-go/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() {
	initializer.LoadEnv()
}

func main() {
	validate := validator.New()
	midtransService := service.NewMidtransServiceImpl(validate)
	midtransController := controller.NewMidtransControllerImpl(midtransService)

	router := gin.Default()
	router.Use(middleware.ErrorHandle())
	midtrans := router.Group("/midtrans")
	{
		midtrans.POST("/create", midtransController.Create)
	}
	router.Run()
}
