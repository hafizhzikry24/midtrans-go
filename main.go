package main

import (
	"log"
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

	db, err := initializer.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	validate := validator.New()
	midtransService := service.NewMidtransServiceImpl(validate, db)
	midtransController := controller.NewMidtransControllerImpl(midtransService)

	router := gin.Default()
	router.Use(middleware.ErrorHandle())
	midtrans := router.Group("/midtrans")
	{
		midtrans.POST("/create", midtransController.Create)
		midtrans.GET("/status/:orderID", midtransController.CheckStatus)

	}
	router.Run()
}
