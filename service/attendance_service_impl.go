package service

import (
	"TalentHub/entity"
	"TalentHub/exception"
	"TalentHub/helper"
	"TalentHub/model"
	"TalentHub/repository"
	"gorm.io/gorm"
	"time"
)

type AttendanceServiceImpl struct {
	DB                   *gorm.DB
	AttendanceRepository repository.AttendanceRepository
}

func NewAttendanceService(db *gorm.DB, attendanceRepository *repository.AttendanceRepository) AttendanceService {
	return &AttendanceServiceImpl{
		DB:                   db,
		AttendanceRepository: *attendanceRepository,
	}
}

func (service *AttendanceServiceImpl) CreateAttendance(request entity.Attendance) entity.Attendance {

	db := service.DB
	attendance, err := service.AttendanceRepository.AddAttendance(db, request)
	exception.PanicIfNeeded(err)
	return attendance

}

func (service *AttendanceServiceImpl) GetAttendanceByEmployeeId(employeeId uint) []model.AttendanceResponse {

	db := service.DB
	attendance, err := service.AttendanceRepository.FindAttendanceByEmployeeId(db, employeeId)
	exception.PanicIfNeeded(err)
	result := helper.ToAttendanceResponses(attendance)
	return result

}

func (service *AttendanceServiceImpl) GetAttendanceByDate(date time.Time) []model.AttendanceResponse {

	db := service.DB
	attendance, err := service.AttendanceRepository.FindAttendanceByDate(db, date)
	exception.PanicIfNeeded(err)
	result := helper.ToAttendanceResponses(attendance)
	return result

}

func (service *AttendanceServiceImpl) GetAttendanceByEmployeeIdAndDate(employeeId uint, date time.Time) model.AttendanceResponse {

	db := service.DB
	attendance, err := service.AttendanceRepository.FindAttendanceByEmployeeIdAndDate(db, employeeId, date)
	exception.PanicIfNeeded(err)
	result := helper.ToAttendanceResponse(attendance)
	return result

}

func (service *AttendanceServiceImpl) GetAttendanceByFilter(employeeId uint, date time.Time) []model.AttendanceResponse {

	db := service.DB
	attendance, err := service.AttendanceRepository.FindAttendanceByFilter(db, employeeId, date)
	exception.PanicIfNeeded(err)
	result := helper.ToAttendanceResponses(attendance)
	return result

}
