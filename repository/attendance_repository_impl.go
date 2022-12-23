package repository

import (
	"TalentHub/entity"
	"gorm.io/gorm"
	"time"
)

type AttendanceRepositoryImpl struct {
}

func NewAttendanceRepository() AttendanceRepository {
	return &AttendanceRepositoryImpl{}
}

func (repository *AttendanceRepositoryImpl) AddAttendance(db *gorm.DB, attendance entity.Attendance) (entity.Attendance, error) {
	err := db.Create(&attendance).Error
	return attendance, err

}

func (repository *AttendanceRepositoryImpl) FindAttendanceByEmployeeId(db *gorm.DB, employee_id uint) ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	err := db.Model(entity.Attendance{}).Preload("Employee").Find(&attendance, "employee_id = ?", employee_id).Error
	return attendance, err
}

func (repository *AttendanceRepositoryImpl) FindAttendanceByDate(db *gorm.DB, date time.Time) ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	err := db.Model(entity.Attendance{}).Find(&attendance, "date = ?", date).Error
	return attendance, err
}

func (repository *AttendanceRepositoryImpl) FindAttendanceByEmployeeIdAndDate(db *gorm.DB, employee_id uint, date time.Time) (entity.Attendance, error) {
	var attendance entity.Attendance
	err := db.Model(entity.Attendance{}).Find(&attendance, "employee_id = ? AND date = ?", employee_id, date).Error
	return attendance, err
}

func (repository *AttendanceRepositoryImpl) FindAttendanceByFilter(db *gorm.DB, employee_id uint, date time.Time) ([]entity.Attendance, error) {
	var attendance []entity.Attendance
	query := db.Model(entity.Attendance{})
	if employee_id != 0 {
		query = query.Where("employee_id = ?", employee_id)
	}
	if !date.IsZero() {
		query = query.Where("date = ?", date)
	}
	err := query.Find(&attendance).Error
	return attendance, err
}
