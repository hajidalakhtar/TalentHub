// pkg/seeds/seeds.go
package seeder

import (
	"TalentHub/model"
	"gorm.io/gorm"
	"time"
)

func All() []model.Seed {
	return []model.Seed{
		model.Seed{
			Name: "Create Personal Data",
			Run: func(db *gorm.DB) error {
				return CreatePersonalData(db,
					"Hajid",
					"Al Akhtar",
					"hajidalakhtar@gmail.com",
					"081234567890",
					"081234567890",
					"Jakarta",
					time.Now(),
					"Male",
					"Single",
					"A",
					"Islam",
					"KTP",
					"1231231231231231",
					time.Now(),
					"12345",
					"Jl. Kebon Jeruk",
					"Jl. Kebon Jeruk",
				)

			},
		},
		model.Seed{
			Name: "Create Employee Data",
			Run: func(db *gorm.DB) error {
				return CreateEmploymentData(db,
					"123123123",
					"Permanent",
					time.Now(),
					time.Now(),
					"Jakarta",
					"IT",
					"Manager",
					"Junior",
					"A",
					"1",
					"Normal",
				)

			},
		},
		model.Seed{
			Name: "Create Salary Data",
			Run: func(db *gorm.DB) error {
				return CreateSalaryData(db,
					"10000",
					"Monthly",
					"notmal",
					"asdasd",
					"Yes",
					"BCA",
					"123123123",
					"asdasd",
					"asdasd",
					"TK0",
					"Gross",
					"Taxable",
					time.Now(),
					"Pegawai tetap",
					"asdasd",
					"asdasd",
					"asdasd",
					"asdasd",
					time.Now(),
					"asdasd",
					"asdasd",
					time.Now(),
					"Company",
					"Company",
					"Company",
					time.Now(),
				)

			},
		},
		model.Seed{
			Name: "Create Allowance ",
			Run: func(db *gorm.DB) error {
				return CreateAllowance(db,
					1000000,
					0,
					0,
				)

			},
		},
		model.Seed{
			Name: "Create Company ",
			Run: func(db *gorm.DB) error {
				return CreateCompany(db,
					"PT. Talenthub Indonesia",
					"081234567890",
					"081234567890",
					"Jakarta",
					"123123123",
					"123123123",
					"123123123",
					"Depok",
					"Jawa Barat",
					"123123123",
					"1 - 0.24%",
					"123123123",
					time.Now(),
				)

			},
		},
		model.Seed{
			Name: "Create Employee ",
			Run: func(db *gorm.DB) error {
				return CreateEmployee(db,
					1,
					1,
					1,
					1,
					1,
				)

			},
		},
	}
}
