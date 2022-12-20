// pkg/seeds/users.go
package seeder

import (
	"TalentHub/entity"
	"gorm.io/gorm"
	"time"
)

func CreatePersonalData(db *gorm.DB,
	first_name string,
	last_name string,
	email string,
	mobile_phone string,
	phone string,
	place_of_birth string,
	birth_date time.Time,
	gender string,
	marital_status string,
	blood_type string,
	religion string,
	identity_type string,
	identity_number string,
	identity_expiry_date time.Time,
	postal_code string,
	citizen_id_address string,
	residence_address string,

) error {
	return db.Create(&entity.PersonalData{
		FirstName:          first_name,
		LastName:           last_name,
		Email:              email,
		MobilePhone:        mobile_phone,
		Phone:              phone,
		PlaceOfBirth:       place_of_birth,
		BirthDate:          birth_date,
		Gender:             gender,
		MaritalStatus:      marital_status,
		BloodType:          blood_type,
		Religion:           religion,
		IdentityType:       identity_type,
		IdentityNumber:     identity_number,
		IdentityExpiryDate: identity_expiry_date,
		PostalCode:         postal_code,
		CitizenIdAddress:   citizen_id_address,
		ResidenceAddress:   residence_address,
	}).Error
}
