package model

import "time"

type AttendanceResponse struct {
	id           uint      `json:"id"`
	Date         time.Time `json:"date"`
	CheckIn      time.Time `json:"check_in"`
	CheckOut     time.Time `json:"check_out"`
	Shift        string    `json:"shift" gorm:"default:'WEEKDAY'"`
	OverTime     string    `json:"overtime"`
	Status       string    `json:"status"`
	EmployeeName string    `json:"employee_name"`
	BasicSalary  string    `json:"basic_salary"`
}
