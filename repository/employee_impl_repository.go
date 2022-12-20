package repository

import (
	"TalentHub/entity"
	"TalentHub/helper"
	"TalentHub/model"
	"gorm.io/gorm"
)

type employeeReposiotryImpl struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeReposiotryImpl{}
}
func (repository *employeeReposiotryImpl) AddEmployee(db *gorm.DB, employee entity.Employee) (entity.Employee, error) {
	err := db.Create(&employee).Error
	return employee, err
}

func (repository *employeeReposiotryImpl) AddEmploymentData(db *gorm.DB, employmentData entity.EmploymentData) (entity.EmploymentData, error) {
	err := db.Create(&employmentData).Error
	return employmentData, err
}

func (repository *employeeReposiotryImpl) AddSalaryData(db *gorm.DB, salaryData entity.SalaryData) (entity.SalaryData, error) {
	err := db.Create(&salaryData).Error
	return salaryData, err
}

func (repository *employeeReposiotryImpl) AddPersonalData(db *gorm.DB, personalData entity.PersonalData) (entity.PersonalData, error) {
	err := db.Create(&personalData).Error
	return personalData, err
}

func (repository *employeeReposiotryImpl) AddAllowance(db *gorm.DB, allowance entity.Allowance) (entity.Allowance, error) {
	err := db.Create(&allowance).Error
	return allowance, err
}

func (repository *employeeReposiotryImpl) FindEmployeeByCompany(db *gorm.DB, company_id uint) ([]model.EmployeeResponse, error) {

	var employee []entity.Employee
	err := db.Model(entity.Employee{}).Preload("PersonalData").Preload("EmploymentData").Preload("SalaryData").Find(&employee, "company_id = ?", company_id).Error
	employeeResponse := helper.ToEmployeeResponses(employee)
	return employeeResponse, err

}

func (repository *employeeReposiotryImpl) FindEmployeeByEmployeeNumber(db *gorm.DB, employee_number string) (*entity.Employee, error) {
	//TODO implement me
	panic("implement me")
}
