package entity

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model
	PersonalDataID   uint `gorm:"not null" json:"personal_data_id"`
	PersonalData     PersonalData
	EmploymentDataID uint `gorm:"not null" json:"employment_data_id"`
	EmploymentData   EmploymentData
	SalaryDataID     uint `gorm:"not null" json:"salary_data_id"`
	SalaryData       SalaryData
	CompanyID        uint `json:"company_id"`
	Company          Company
	AllowanceID      uint `json:"allowance_id"`
	Allowance        Allowance
}

type PersonalData struct {
	gorm.Model
	FirstName     string    `json:"first_name" gorm:"not null"`
	LastName      string    `json:"last_name" `
	Email         string    `json:"email" gorm:"not null"`
	MobilePhone   string    `json:"mobile_phone"`
	Phone         string    `json:"phone"`
	PlaceOfBirth  string    `json:"place_of_birth" gorm:"not null"`
	BirthDate     time.Time `json:"birth_date" gorm:"not null"`
	Gender        string    `json:"gender" gorm:"type:gender_options;not null"`
	MaritalStatus string    `json:"marital_status" gorm:"type:marital_status_options"`
	BloodType     string    `json:"blood_type" gorm:"type:blood_type_options"`
	Religion      string    `json:"religion" gorm:"type:religion_options"`

	IdentityType       string    `json:"identity_type" gorm:"type:identity_type_options;not null"`
	IdentityNumber     string    `json:"identity_number" gorm:"not null"`
	IdentityExpiryDate time.Time `json:"identity_expiry_date"`
	PostalCode         string    `json:"postal_code"`
	CitizenIdAddress   string    `json:"citizen_id_address"`
	ResidenceAddress   string    `json:"residence_address"`
}

type EmploymentData struct {
	gorm.Model
	EmployeeNumber string    `json:"employee_number" gorm:"not null"`
	EmployeeStatus string    `json:"employee_status" gorm:"type:employee_status_options;not null"`
	JoinDate       time.Time `json:"join_date" gorm:"not null"`
	EndStatusDate  time.Time `json:"end_status_date"`
	Branch         string    `json:"branch"`                       //buat master
	Organization   string    `json:"organization" gorm:"not null"` //buat master
	JobPosition    string    `json:"job_position" gorm:"not null"`
	JobLevel       string    `json:"job_level" gorm:"type:job_level_options;not null"`
	Grade          string    `json:"grade" gorm:"type:grade_options;"`
	Class          string    `json:"class" gorm:"type:class_options;"`
	Schedule       string    `json:"schedule"` //buat master
}

type SalaryData struct {
	gorm.Model
	BasicSalary      string `json:"basic_salary" gorm:"not null"`
	SalaryType       string `json:"salary_type" gorm:"type:salary_type_options;not null"`
	PaymentSchedule  string `json:"payment_schedule"` //buat master
	ProrateSetting   string `json:"prorate_setting" `
	AllowForOvertime string `json:"allow_for_overtime" gorm:"type:allow_for_overtime_options;not null"`

	BankName          string `json:"bank_name" gorm:"not null"`
	BankAccountNumber string `json:"bank_account_number" gorm:"not null"`
	BankAccountName   string `json:"bank_account_name" gorm:"not null"`

	NPWP              string    `json:"npwp" gorm:"not null"`
	PTKPStatus        string    `json:"ptkp_status" gorm:"type:ptkp_status_options;not null"`
	TaxMethod         string    `json:"tax_method" gorm:"type:tax_method_options;not null"`
	TaxSalary         string    `json:"tax_salary" gorm:"type:tax_salary_options;not null"`
	TaxableDate       time.Time `json:"taxable_date" `
	EmployeeTaxStatus string    `json:"employee_tax_status" gorm:"type:employee_tax_status_options;not null"`
	BeginningNetto    string    `json:"beginning_netto"`
	PHH21Paid         string    `json:"phh21_paid"`

	BPJSKetenagaKerjaNumber string    `json:"bpjs_ketenaga_kerja_number"`
	NPPBPJSKetenagaKerja    string    `json:"npp_bpjs_ketenaga_kerja"`
	BPJSKetenagaKerjaDate   time.Time `json:"bpjs_ketenaga_kerja_date"`
	BPJSKesehatanNumber     string    `json:"bpjs_kesehatan_number"`
	BPJSKesehatanFamily     string    `json:"bpjs_kesehatan_family"`
	BPJSKesehatanDate       time.Time `json:"bpjs_kesehatan_date"`
	BPJSKesehatanCost       string    `json:"bpjs_kesehatan_cost" gorm:"type:bpjs_kesehatan_cost_options;"`
	JHTCost                 string    `json:"jht_cost" gorm:"type:jht_cost_options;"`
	JaminanPensiunCost      string    `json:"jaminan_pensiun_cost" gorm:"type:jaminan_pensiun_cost_options;"`
	JaminanPensiunDate      time.Time `json:"jaminan_pensiun_date"`
}

type Allowance struct {
	gorm.Model
	TunjanganLainlain int `json:"tunjangan_lainlain" gorm:"default:0"`
	TunjanganJabatan  int `json:"tunjangan_jabatan" gorm:"default:0"`
	TunjanganBulanan  int `json:"tunjangan_bulanan" gorm:"default:0"`
}
