package service

import "TalentHub/entity"

type CompanyService interface {
	CreateCompany(request entity.Company) entity.Company
	GetCompanyById(id uint) entity.Company
}
