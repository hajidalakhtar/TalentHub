package entity

import (
	"gorm.io/gorm"
	"time"
)

type Company struct {
	gorm.Model
	Name                    string    `gorm:"unique" json:"name"`
	Description             string    `json:"description"`
	HQInitial               string    `json:"hq_initial"`
	Email                   string    `gorm:"unique" json:"email"`
	Phone                   string    `json:"phone"`
	Address                 string    `json:"address"`
	PostalCode              string    `json:"postal_code"`
	City                    string    `json:"city"`
	Province                string    `json:"province"`
	BPJSKetenagaKerjaNumber string    `json:"bpjs_ketenaga_kerja_number"`
	JKK                     string    `json:"jkk" gorm:"type:jkk_options"`
	NPWP                    string    `json:"npwp"`
	TaxableDate             time.Time `json:"taxable_date"`
}
