package repository

import (
	"TalentHub/entity"
	"gorm.io/gorm"
	"time"
)

type AttendanceRepository interface {
	AddAttendance(db *gorm.DB, attendance entity.Attendance) (entity.Attendance, error)
	FindAttendanceByFilter(db *gorm.DB, employee_id uint, date time.Time) ([]entity.Attendance, error)

	FindAttendanceByEmployeeId(db *gorm.DB, employee_id uint) ([]entity.Attendance, error)
	FindAttendanceByDate(db *gorm.DB, date time.Time) ([]entity.Attendance, error)
	FindAttendanceByEmployeeIdAndDate(db *gorm.DB, employee_id uint, date time.Time) (entity.Attendance, error)
}
