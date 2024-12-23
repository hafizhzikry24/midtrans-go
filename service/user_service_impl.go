package service

import (
	"midtrans-go/helper"
	"midtrans-go/model/entity"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	Validate *validator.Validate
	DB       *gorm.DB
}

func NewUserServiceImpl(validate *validator.Validate, db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{
		Validate: validate,
		DB:       db,
	}
}

func (service *UserServiceImpl) CreateUser(request entity.User) entity.User {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	err = service.DB.Create(&request).Error
	helper.PanicIfError(err)
	return request
}

func (service *UserServiceImpl) GetUserWithTransaction(UserID uint) (*entity.User, error) {
	var user entity.User
	err := service.DB.Preload("Transaction").Where("id = ?", UserID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
