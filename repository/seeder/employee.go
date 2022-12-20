// pkg/seeds/users.go
package seeder

import (
	"TalentHub/entity"
	"gorm.io/gorm"
)

func CreateEmployee(db *gorm.DB,
	personal_data_id uint,
	employment_data_id uint,
	salary_data_id uint,
	allowance_id uint,
	company_id uint,

) error {
	return db.Create(&entity.Employee{
		PersonalDataID:   personal_data_id,
		EmploymentDataID: employment_data_id,
		SalaryDataID:     salary_data_id,
		AllowanceID:      allowance_id,
		CompanyID:        company_id,
	}).Error
}
