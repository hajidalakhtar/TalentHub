package service

import (
	"TalentHub/config"
	"TalentHub/entity"
	"TalentHub/exception"
	"TalentHub/helper"
	"TalentHub/model"
	"TalentHub/repository"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"
)

type UserServiceImpl struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
	Configuration  config.Config
}

func NewUserService(db *gorm.DB, userRepository *repository.UserRepository, appConfiguration config.Config) UserService {
	return &UserServiceImpl{DB: db, UserRepository: *userRepository, Configuration: appConfiguration}
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

func (service *UserServiceImpl) Login(email string, password string) (model.LoginResponse, bool) {
	db := service.DB
	user, err := service.UserRepository.FindUserByEmail(db, email)
	isLoginSuccess := false

	if err == nil {
		checkPassword := helper.CheckPasswordHash(password, user.Password)
		if checkPassword {
			isLoginSuccess = true
		}
	}

	signedToken := createToken(user, service.Configuration.Get("JWT_KEY"))

	loginResponse := helper.ToLoginResponse(user, signedToken)
	return loginResponse, isLoginSuccess
}

func (service *UserServiceImpl) FindUserById(id uint) model.UserResponse {
	db := service.DB
	user, err := service.UserRepository.FindUserById(db, id)
	exception.PanicIfNeeded(err)
	userResponse := helper.ToUserResponse(user)
	return userResponse
}

func createToken(user entity.User, jwtKey string) string {
	claims := model.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "TalentHub",
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
		Username: user.Username,
		Email:    user.Email,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		panic(err)
	}
	return signedToken

}
