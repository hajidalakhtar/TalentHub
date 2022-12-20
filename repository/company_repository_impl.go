package repository

import (
	"TalentHub/entity"
	"gorm.io/gorm"
)

type companyRepositoryImpl struct {
}

func NewCompanyRepository() CompanyRepository {
	return &companyRepositoryImpl{}
}

func (repository *companyRepositoryImpl) AddCompany(db *gorm.DB, company entity.Company) (entity.Company, error) {

	err := db.Create(&company).Error
	return company, err

}

func (repository *companyRepositoryImpl) FindCompanyById(db *gorm.DB, id uint) (entity.Company, error) {

	var company entity.Company
	err := db.First(&company, id).Error
	return company, err

}
