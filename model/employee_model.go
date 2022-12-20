package model

import "time"

type CreateEmployeeRequest struct {
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name" `
	Email         string    `json:"email" `
	MobilePhone   string    `json:"mobile_phone"`
	Phone         string    `json:"phone"`
	PlaceOfBirth  string    `json:"place_of_birth" `
	BirthDate     time.Time `json:"birth_date" `
	Gender        string    `json:"gender" `
	MaritalStatus string    `json:"marital_status"`
	BloodType     string    `json:"blood_type"`
	Religion      string    `json:"religion"`

	IdentityType       string    `json:"identity_type"`
	IdentityNumber     string    `json:"identity_number" `
	IdentityExpiryDate time.Time `json:"identity_expiry_date"`
	PostalCode         string    `json:"postal_code"`
	CitizenIdAddress   string    `json:"citizen_id_address"`
	ResidenceAddress   string    `json:"residence_address"`

	EmployeeNumber string    `json:"employee_number" `
	EmployeeStatus string    `json:"employee_status"`
	JoinDate       time.Time `json:"join_date" `
	EndStatusDate  time.Time `json:"end_status_date"`
	Branch         string    `json:"branch"`        //buat master
	Organization   string    `json:"organization" ` //buat master
	JobPosition    string    `json:"job_position" `
	JobLevel       string    `json:"job_level"`
	Grade          string    `json:"grade"`
	Class          string    `json:"class"`
	Schedule       string    `json:"schedule"` //buat master

	BasicSalary      string `json:"basic_salary" `
	SalaryType       string `json:"salary_type"`
	PaymentSchedule  string `json:"payment_schedule"` //buat master
	ProrateSetting   string `json:"prorate_setting" `
	AllowForOvertime string `json:"allow_for_overtime"`

	BankName          string `json:"bank_name" `
	BankAccountNumber string `json:"bank_account_number" `
	BankAccountName   string `json:"bank_account_name" `

	NPWP              string    `json:"npwp" `
	PTKPStatus        string    `json:"ptkp_status"`
	TaxMethod         string    `json:"tax_method"`
	TaxSalary         string    `json:"tax_salary"`
	TaxableDate       time.Time `json:"taxable_date" `
	EmployeeTaxStatus string    `json:"employee_tax_status"`
	BeginningNetto    string    `json:"beginning_netto"`
	PHH21Paid         string    `json:"phh21_paid"`

	BPJSKetenagaKerjaNumber string    `json:"bpjs_ketenaga_kerja_number"`
	NPPBPJSKetenagaKerja    string    `json:"npp_bpjs_ketenaga_kerja"`
	BPJSKetenagaKerjaDate   time.Time `json:"bpjs_ketenaga_kerja_date"`
	BPJSKesehatanNumber     string    `json:"bpjs_kesehatan_number"`
	BPJSKesehatanFamily     string    `json:"bpjs_kesehatan_family"`
	BPJSKesehatanDate       time.Time `json:"bpjs_kesehatan_date"`
	BPJSKesehatanCost       string    `json:"bpjs_kesehatan_cost" `
	JHTCost                 string    `json:"jht_cost" `
	JaminanPensiunCost      string    `json:"jaminan_pensiun_cost" `
	JaminanPensiunDate      time.Time `json:"jaminan_pensiun_date"`
}

type EmployeeResponse struct {
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name" `
	Email         string    `json:"email" `
	MobilePhone   string    `json:"mobile_phone"`
	Phone         string    `json:"phone"`
	PlaceOfBirth  string    `json:"place_of_birth" `
	BirthDate     time.Time `json:"birth_date" `
	Gender        string    `json:"gender" `
	MaritalStatus string    `json:"marital_status"`
	BloodType     string    `json:"blood_type"`
	Religion      string    `json:"religion"`

	IdentityType       string    `json:"identity_type"`
	IdentityNumber     string    `json:"identity_number" `
	IdentityExpiryDate time.Time `json:"identity_expiry_date"`
	PostalCode         string    `json:"postal_code"`
	CitizenIdAddress   string    `json:"citizen_id_address"`
	ResidenceAddress   string    `json:"residence_address"`

	EmployeeNumber string    `json:"employee_number" `
	EmployeeStatus string    `json:"employee_status"`
	JoinDate       time.Time `json:"join_date" `
	EndStatusDate  time.Time `json:"end_status_date"`
	Branch         string    `json:"branch"`        //buat master
	Organization   string    `json:"organization" ` //buat master
	JobPosition    string    `json:"job_position" `
	JobLevel       string    `json:"job_level"`
	Grade          string    `json:"grade"`
	Class          string    `json:"class"`
	Schedule       string    `json:"schedule"` //buat master

	BasicSalary      string `json:"basic_salary" `
	SalaryType       string `json:"salary_type"`
	PaymentSchedule  string `json:"payment_schedule"` //buat master
	ProrateSetting   string `json:"prorate_setting" `
	AllowForOvertime string `json:"allow_for_overtime"`

	BankName          string `json:"bank_name" `
	BankAccountNumber string `json:"bank_account_number" `
	BankAccountName   string `json:"bank_account_name" `

	NPWP              string    `json:"npwp" `
	PTKPStatus        string    `json:"ptkp_status"`
	TaxMethod         string    `json:"tax_method"`
	TaxSalary         string    `json:"tax_salary"`
	TaxableDate       time.Time `json:"taxable_date" `
	EmployeeTaxStatus string    `json:"employee_tax_status"`
	BeginningNetto    string    `json:"beginning_netto"`
	PHH21Paid         string    `json:"phh21_paid"`

	BPJSKetenagaKerjaNumber string    `json:"bpjs_ketenaga_kerja_number"`
	NPPBPJSKetenagaKerja    string    `json:"npp_bpjs_ketenaga_kerja"`
	BPJSKetenagaKerjaDate   time.Time `json:"bpjs_ketenaga_kerja_date"`
	BPJSKesehatanNumber     string    `json:"bpjs_kesehatan_number"`
	BPJSKesehatanFamily     string    `json:"bpjs_kesehatan_family"`
	BPJSKesehatanDate       time.Time `json:"bpjs_kesehatan_date"`
	BPJSKesehatanCost       string    `json:"bpjs_kesehatan_cost" `
	JHTCost                 string    `json:"jht_cost" `
	JaminanPensiunCost      string    `json:"jaminan_pensiun_cost" `
	JaminanPensiunDate      time.Time `json:"jaminan_pensiun_date"`
}
