package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string   `gorm:"unique" json:"username"`
	Password   string   `json:"password"`
	Email      string   `gorm:"unique" json:"email"`
	EmployeeId uint     `json:"employee_id"`
	Employee   Employee `json:"employee"`
	CompanyId  uint     `json:"company_id"`
	Company    Company  `json:"company"`
}
