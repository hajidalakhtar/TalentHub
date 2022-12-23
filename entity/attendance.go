package entity

import (
	"gorm.io/gorm"
	"time"
)

type Attendance struct {
	gorm.Model
	EmployeeId uint      `json:"employee_id"`
	Employee   Employee  `json:"employee"`
	Date       time.Time `json:"date"`
	CheckIn    time.Time `json:"start_time"`
	CheckOut   time.Time `json:"end_time"`
	Shift      string    `json:"shift"`
	OverTime   string    `json:"overtime"`
	Status     string    `json:"status"`
}
