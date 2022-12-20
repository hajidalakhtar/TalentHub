package repository

import (
	"TalentHub/entity"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) AddUser(db *gorm.DB, user entity.User) (entity.User, error) {
	err := db.Create(&user).Error
	return user, err
}

func (repository *userRepositoryImpl) FindUserById(db *gorm.DB, id uint) (entity.User, error) {
	var user entity.User
	err := db.First(&user, id).Error
	return user, err
}

func (repository *userRepositoryImpl) FindUserByEmail(db *gorm.DB, email string) (entity.User, error) {
	var user entity.User
	err := db.First(&user, "email = ?", email).Error
	return user, err
}
