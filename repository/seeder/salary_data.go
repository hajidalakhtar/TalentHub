// pkg/seeds/users.go
package seeder

import (
	"TalentHub/entity"
	"gorm.io/gorm"
	"time"
)

func CreateSalaryData(db *gorm.DB,
	basic_salary string,
	salary_type string,
	payment_schedule string,
	prorate_setting string,
	allow_for_overtime string,
	bank_name string,
	bank_account_number string,
	bank_account_name string,
	npwp string,
	ptkp_status string,
	tax_method string,
	tax_salary string,
	taxable_date time.Time,
	employee_tax_status string,
	beginning_netto string,
	phh21_paid string,
	bpjs_ketenaga_kerja_number string,
	npp_bpjs_ketenaga_kerja string,
	bpjs_ketenaga_kerja_date time.Time,
	bpjs_kesehatan_number string,
	bpjs_kesehatan_family string,
	bpjs_kesehatan_date time.Time,
	bpjs_kesehatan_cost string,
	jht_cost string,
	jaminan_pensiun_cost string,
	jaminan_pensiun_date time.Time,

) error {
	return db.Create(&entity.SalaryData{
		BasicSalary:             basic_salary,
		SalaryType:              salary_type,
		PaymentSchedule:         payment_schedule,
		ProrateSetting:          prorate_setting,
		AllowForOvertime:        allow_for_overtime,
		BankName:                bank_name,
		BankAccountNumber:       bank_account_number,
		BankAccountName:         bank_account_name,
		NPWP:                    npwp,
		PTKPStatus:              ptkp_status,
		TaxMethod:               tax_method,
		TaxSalary:               tax_salary,
		TaxableDate:             taxable_date,
		EmployeeTaxStatus:       employee_tax_status,
		BeginningNetto:          beginning_netto,
		PHH21Paid:               phh21_paid,
		BPJSKetenagaKerjaNumber: bpjs_ketenaga_kerja_number,
		NPPBPJSKetenagaKerja:    npp_bpjs_ketenaga_kerja,
		BPJSKetenagaKerjaDate:   bpjs_ketenaga_kerja_date,
		BPJSKesehatanNumber:     bpjs_kesehatan_number,
		BPJSKesehatanFamily:     bpjs_kesehatan_family,
		BPJSKesehatanDate:       bpjs_kesehatan_date,
		BPJSKesehatanCost:       bpjs_kesehatan_cost,
		JHTCost:                 jht_cost,
		JaminanPensiunCost:      jaminan_pensiun_cost,
		JaminanPensiunDate:      jaminan_pensiun_date,
	}).Error
}
