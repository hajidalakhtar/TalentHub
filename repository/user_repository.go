package repository

import (
	"TalentHub/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(db *gorm.DB, user entity.User) (entity.User, error)
	FindUserById(db *gorm.DB, id uint) (entity.User, error)
	FindUserByEmail(db *gorm.DB, email string) (entity.User, error)
}
