package service

import (
	"TalentHub/entity"
	"TalentHub/exception"
	"TalentHub/helper"
	"TalentHub/model"
	"TalentHub/repository"
	"gorm.io/gorm"
)

type employeeService struct {
	DB                 *gorm.DB
	EmployeeRepository repository.EmployeeRepository
}

func NewEmployeeService(db *gorm.DB, employeeRepository *repository.EmployeeRepository) EmployeeService {
	return &employeeService{DB: db, EmployeeRepository: *employeeRepository}
}

func (service *employeeService) CreateEmployee(request model.CreateEmployeeRequest, companyId uint) entity.Employee {

	db := service.DB

	//companyId
	personalData, employmentData, salaryData, allowance := helper.ToAllEmployeeModelFromRequest(request)

	employmentDataResult, err := service.EmployeeRepository.AddEmploymentData(db, employmentData)
	exception.PanicIfNeeded(err)

	salaryDataResult, err := service.EmployeeRepository.AddSalaryData(db, salaryData)
	exception.PanicIfNeeded(err)

	personalDataResult, err := service.EmployeeRepository.AddPersonalData(db, personalData)
	exception.PanicIfNeeded(err)

	allowanceResult, err := service.EmployeeRepository.AddAllowance(db, allowance)
	exception.PanicIfNeeded(err)

	employee := entity.Employee{
		CompanyID:        companyId,
		PersonalDataID:   personalDataResult.ID,
		EmploymentDataID: employmentDataResult.ID,
		SalaryDataID:     salaryDataResult.ID,
		AllowanceID:      allowanceResult.ID,
	}

	employeeResult, err := service.EmployeeRepository.AddEmployee(db, employee)
	exception.PanicIfNeeded(err)

	return employeeResult

}

func (service *employeeService) GetEmployeeByCompany(companyId uint) []model.EmployeeResponse {
	db := service.DB
	employeeResult, err := service.EmployeeRepository.FindEmployeeByCompany(db, companyId)
	exception.PanicIfNeeded(err)

	return employeeResult
}
