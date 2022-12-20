package repository

import (
	"TalentHub/entity"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	AddCompany(db *gorm.DB, company entity.Company) (entity.Company, error)
	FindCompanyById(db *gorm.DB, id uint) (entity.Company, error)
}
