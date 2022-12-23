package config

import (
	"TalentHub/entity"
	"TalentHub/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(configuration Config) *gorm.DB {

	dsn := configuration.Get("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicIfNeeded(err)
	return db

}

func Migration(db *gorm.DB) error {

	db.Migrator().DropTable(&entity.User{})
	db.Migrator().DropTable(&entity.Employee{})
	db.Migrator().DropTable(&entity.SalaryData{})
	db.Migrator().DropTable(&entity.Company{})
	db.Migrator().DropTable(&entity.PersonalData{})
	db.Migrator().DropTable(&entity.EmploymentData{})
	db.Migrator().DropTable(&entity.Allowance{})

	db.Exec("CREATE TYPE gender_options AS ENUM ('Male','Female');")
	db.Exec("CREATE TYPE marital_status_options AS ENUM ('Single','Married','Divorced');")
	db.Exec("CREATE TYPE blood_type_options AS ENUM ('A','B','AB','O');")
	db.Exec("CREATE TYPE religion_options AS ENUM ('Islam','Protestan','Katolik','Hindu','Budha');")
	db.Exec("CREATE TYPE identity_type_options AS ENUM ('KTP','Passport');")

	db.Exec("CREATE TYPE employee_status_options AS ENUM ('Permanent', 'Contract', 'Probation');")
	db.Exec("CREATE TYPE job_level_options AS ENUM ('Junior', 'Middle', 'Senior');")
	db.Exec("CREATE TYPE grade_options AS ENUM ('A', 'B', 'C', 'D');")
	db.Exec("CREATE TYPE class_options AS ENUM ('1', '2', '3', '4');")

	db.Exec("CREATE TYPE salary_type_options AS ENUM ('Monthly', 'Daily');")
	db.Exec("CREATE TYPE allow_for_overtime_options AS ENUM ('Yes', 'No');")
	db.Exec("CREATE TYPE ptkp_status_options AS ENUM ('TK0', 'TK1', 'TK2', 'TK3');")
	db.Exec("CREATE TYPE tax_method_options AS ENUM ('Gross', 'Gross Up','netto');")
	db.Exec("CREATE TYPE tax_salary_options AS ENUM ('Taxable', 'Non-Taxable');")
	db.Exec("CREATE TYPE employee_tax_status_options AS ENUM ('Pegawai tetap', 'Pegawai tidak tetap');")
	db.Exec("CREATE TYPE bpjs_kesehatan_cost_options AS ENUM ('Company', 'Employee');")
	db.Exec("CREATE TYPE jaminan_pensiun_cost_options AS ENUM ('Company', 'Employee');")
	db.Exec("CREATE TYPE jht_cost_options AS ENUM ('Company', 'Employee');")
	db.Exec("CREATE TYPE jkk_options AS ENUM ('1 - 0.24%', '2 - 0.54%', '3 - 0.89%','4 - 1.27%','5 - 1.24%');")

	err := db.AutoMigrate(&entity.PersonalData{}, &entity.EmploymentData{}, &entity.SalaryData{}, &entity.Allowance{}, &entity.Employee{}, &entity.User{}, &entity.Company{}, &entity.Attendance{})
	return err
}
