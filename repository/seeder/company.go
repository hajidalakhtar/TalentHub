// pkg/seeds/users.go
package seeder

import (
	"TalentHub/entity"
	"gorm.io/gorm"
	"time"
)

func CreateCompany(db *gorm.DB,
	name string,
	description string,
	hq_initial string,
	email string,
	phone string,
	address string,
	postal_code string,
	city string,
	province string,
	bpjs_ketenaga_kerja_number string,
	jkk string,
	npwp string,
	taxable_date time.Time,

) error {
	return db.Create(&entity.Company{
		Name:                    name,
		Description:             description,
		HQInitial:               hq_initial,
		Email:                   email,
		Phone:                   phone,
		Address:                 address,
		PostalCode:              postal_code,
		City:                    city,
		Province:                province,
		BPJSKetenagaKerjaNumber: bpjs_ketenaga_kerja_number,
		JKK:                     jkk,
		NPWP:                    npwp,
		TaxableDate:             taxable_date,
	}).Error
}
