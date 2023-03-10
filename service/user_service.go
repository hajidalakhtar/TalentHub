package service

import (
	"TalentHub/entity"
	"TalentHub/model"
)

type UserService interface {
	CreateUser(request entity.User) model.UserResponse
	Login(email string, password string) (model.LoginResponse, bool)
	FindUserById(id uint) model.UserResponse
}
