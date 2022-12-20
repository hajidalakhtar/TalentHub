// pkg/seeds/users.go
package seeder

import (
	"TalentHub/entity"
	"gorm.io/gorm"
	"time"
)

func CreateEmploymentData(db *gorm.DB,
	employee_number string,
	employee_status string,
	join_date time.Time,
	end_status_date time.Time,
	branch string,
	organization string,
	job_position string,
	job_level string,
	grade string,
	class string,
	schedule string,

) error {
	return db.Create(&entity.EmploymentData{
		EmployeeNumber: employee_number,
		EmployeeStatus: employee_status,
		JoinDate:       join_date,
		EndStatusDate:  end_status_date,
		Branch:         branch,
		Organization:   organization,
		JobPosition:    job_position,
		JobLevel:       job_level,
		Grade:          grade,
		Class:          class,
		Schedule:       schedule,
	}).Error
}
