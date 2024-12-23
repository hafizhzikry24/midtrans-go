package service

import (
	"midtrans-go/model/entity"
)

type UserService interface {
	CreateUser(request entity.User) entity.User
	GetUserWithTransaction(UserID uint) (*entity.User, error)
}
