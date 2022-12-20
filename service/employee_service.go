package service

import (
	"TalentHub/entity"
	"TalentHub/model"
)

type EmployeeService interface {
	CreateEmployee(request model.CreateEmployeeRequest, companyId uint) entity.Employee
	GetEmployeeByCompany(companyId uint) []model.EmployeeResponse
}
