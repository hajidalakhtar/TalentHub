package repository

import (
	"TalentHub/entity"
	"TalentHub/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	AddEmployee(db *gorm.DB, employee entity.Employee) (entity.Employee, error)
	AddEmploymentData(db *gorm.DB, employmentData entity.EmploymentData) (entity.EmploymentData, error)
	AddSalaryData(db *gorm.DB, salaryData entity.SalaryData) (entity.SalaryData, error)
	AddPersonalData(db *gorm.DB, personalData entity.PersonalData) (entity.PersonalData, error)
	AddAllowance(db *gorm.DB, allowance entity.Allowance) (entity.Allowance, error)

	FindEmployeeByCompany(db *gorm.DB, company_id uint) ([]model.EmployeeResponse, error)
	FindEmployeeByEmployeeNumber(db *gorm.DB, employee_number string) (*entity.Employee, error)
}
