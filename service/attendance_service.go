package service

import (
	"TalentHub/entity"
	"TalentHub/model"
	"time"
)

type AttendanceService interface {
	CreateAttendance(request entity.Attendance) entity.Attendance
	GetAttendanceByEmployeeId(employeeId uint) []model.AttendanceResponse
	GetAttendanceByDate(date time.Time) []model.AttendanceResponse
	GetAttendanceByEmployeeIdAndDate(employeeId uint, date time.Time) model.AttendanceResponse
	GetAttendanceByFilter(employeeId uint, date time.Time) []model.AttendanceResponse
}
                        