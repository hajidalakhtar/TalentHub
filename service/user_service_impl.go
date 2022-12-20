package service

import (
	"TalentHub/entity"
	"TalentHub/exception"
	"TalentHub/helper"
	"TalentHub/model"
	"TalentHub/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
}

func NewUserService(db *gorm.DB, userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{DB: db, UserRepository: *userRepository}
}

func (service *UserServiceImpl) CreateUser(user entity.User) model.UserResponse {
	db := service.DB
	hash, err := helper.HashPassword(user.Password)
	exception.PanicIfNeeded(err)
	user.Password = hash
	service.UserRepository.AddUser(db, user)
	userResponse := helper.ToUserResponse(user)
	return userResponse
}

func (service *UserServiceImpl) Login(email string, password string) (model.UserResponse, bool) {
	db := service.DB
	user, err := service.UserRepository.FindUserByEmail(db, email)
	isLoginSuccess := false

	if err == nil {
		checkPassword := helper.CheckPasswordHash(password, user.Password)
		if checkPassword {
			isLoginSuccess = true
		}
	}

	userResponse := helper.ToUserResponse(user)
	return userResponse, isLoginSuccess
}

func (service *UserServiceImpl) FindUserById(id uint) model.UserResponse {
	db := service.DB
	user, err := service.UserRepository.FindUserById(db, id)
	exception.PanicIfNeeded(err)
	userResponse := helper.ToUserResponse(user)
	return userResponse
}
