// pkg/seeds/users.go
package seeder

import (
	"TalentHub/entity"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB,
	username string,
	password string,
	email string,
	company_id uint,
	employee_id uint,
) error {
	return db.Create(&entity.User{
		Username:   username,
		Password:   password,
		Email:      email,
		CompanyId:  company_id,
		EmployeeId: employee_id,
	}).Error
}
