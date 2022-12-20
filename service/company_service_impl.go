package service

import (
	"TalentHub/entity"
	"TalentHub/exception"
	"TalentHub/repository"
	"gorm.io/gorm"
)

type companyServiceImpl struct {
	DB                *gorm.DB
	CompanyRepository repository.CompanyRepository
}

func NewCompanyService(db *gorm.DB, companyRepository *repository.CompanyRepository) CompanyService {
	return &companyServiceImpl{
		DB:                db,
		CompanyRepository: *companyRepository,
	}
}

func (service *companyServiceImpl) CreateCompany(request entity.Company) entity.Company {

	db := service.DB
	company, err := service.CompanyRepository.AddCompany(db, request)
	exception.PanicIfNeeded(err)
	return company

}

func (service *companyServiceImpl) GetCompanyById(id uint) entity.Company {

	db := service.DB
	company, err := service.CompanyRepository.FindCompanyById(db, id)
	exception.PanicIfNeeded(err)

	return company

}
