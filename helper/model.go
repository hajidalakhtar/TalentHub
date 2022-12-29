package helper

import (
	"TalentHub/entity"
	"TalentHub/model"
	"time"
)

func ToAllEmployeeModelFromRequest(request model.CreateEmployeeRequest) (entity.PersonalData, entity.EmploymentData, entity.SalaryData, entity.Allowance) {

	dateFormatReference := "2006-01-02T15:04:05Z07:00"

	BirthDateInDateFormat, err := time.Parse(dateFormatReference, request.BirthDate)
	IdentityExpiryDateInDateFormat, err := time.Parse(dateFormatReference, request.IdentityExpiryDate)
	JoinDateInDateFormat, err := time.Parse(dateFormatReference, request.JoinDate)
	EndStatusDateInDateFormat, err := time.Parse(dateFormatReference, request.EndStatusDate)
	TaxableDateInDateFormat, err := time.Parse(dateFormatReference, request.TaxableDate)
	BPJSKetenagaKerjaDateInDateFormat, err := time.Parse(dateFormatReference, request.BPJSKetenagaKerjaDate)
	BPJSKesehatanDateInDateFormat, err := time.Parse(dateFormatReference, request.BPJSKesehatanDate)
	JaminanPensiunDateInDateFormat, err := time.Parse(dateFormatReference, request.JaminanPensiunDate)
	if err != nil {
		panic(err)
	}

	personalData := entity.PersonalData{
		FirstName:          request.FirstName,
		LastName:           request.LastName,
		Email:              request.Email,
		MobilePhone:        request.MobilePhone,
		Phone:              request.Phone,
		PlaceOfBirth:       request.PlaceOfBirth,
		BirthDate:          BirthDateInDateFormat,
		Gender:             request.Gender,
		MaritalStatus:      request.MaritalStatus,
		BloodType:          request.BloodType,
		Religion:           request.Religion,
		IdentityType:       request.IdentityType,
		IdentityNumber:     request.IdentityNumber,
		IdentityExpiryDate: IdentityExpiryDateInDateFormat,
		PostalCode:         request.PostalCode,
		CitizenIdAddress:   request.CitizenIdAddress,
		ResidenceAddress:   request.ResidenceAddress,
	}

	employmentData := entity.EmploymentData{
		EmployeeNumber: request.EmployeeNumber,
		EmployeeStatus: request.EmployeeStatus,
		JoinDate:       JoinDateInDateFormat,
		EndStatusDate:  EndStatusDateInDateFormat,
		Branch:         request.Branch,
		Organization:   request.Organization,
		JobPosition:    request.JobPosition,
		JobLevel:       request.JobLevel,
		Grade:          request.Grade,
		Class:          request.Class,
		Schedule:       request.Schedule,
	}

	salaryData := entity.SalaryData{
		BasicSalary:             request.BasicSalary,
		SalaryType:              request.SalaryType,
		PaymentSchedule:         request.PaymentSchedule,
		ProrateSetting:          request.ProrateSetting,
		AllowForOvertime:        request.AllowForOvertime,
		BankName:                request.BankName,
		BankAccountNumber:       request.BankAccountNumber,
		BankAccountName:         request.BankAccountName,
		NPWP:                    request.NPWP,
		PTKPStatus:              request.PTKPStatus,
		TaxMethod:               request.TaxMethod,
		TaxSalary:               request.TaxSalary,
		TaxableDate:             TaxableDateInDateFormat,
		EmployeeTaxStatus:       request.EmployeeTaxStatus,
		BeginningNetto:          request.BeginningNetto,
		PHH21Paid:               request.PHH21Paid,
		BPJSKetenagaKerjaNumber: request.BPJSKetenagaKerjaNumber,
		NPPBPJSKetenagaKerja:    request.NPPBPJSKetenagaKerja,
		BPJSKetenagaKerjaDate:   BPJSKetenagaKerjaDateInDateFormat,
		BPJSKesehatanNumber:     request.BPJSKesehatanNumber,
		BPJSKesehatanFamily:     request.BPJSKesehatanFamily,
		BPJSKesehatanDate:       BPJSKesehatanDateInDateFormat,
		BPJSKesehatanCost:       request.BPJSKesehatanCost,
		JHTCost:                 request.JHTCost,
		JaminanPensiunCost:      request.JaminanPensiunCost,
		JaminanPensiunDate:      JaminanPensiunDateInDateFormat,
	}

	allowance := entity.Allowance{
		TunjanganLainlain: 0,
		TunjanganJabatan:  0,
		TunjanganBulanan:  0,
	}

	return personalData, employmentData, salaryData, allowance

}

func ToEmployeeResponse(employee entity.Employee) model.EmployeeResponse {

	employeeResponse := model.EmployeeResponse{
		FirstName:               employee.PersonalData.FirstName,
		LastName:                employee.PersonalData.LastName,
		Email:                   employee.PersonalData.Email,
		MobilePhone:             employee.PersonalData.MobilePhone,
		Phone:                   employee.PersonalData.Phone,
		PlaceOfBirth:            employee.PersonalData.PlaceOfBirth,
		BirthDate:               employee.PersonalData.BirthDate,
		Gender:                  employee.PersonalData.Gender,
		MaritalStatus:           employee.PersonalData.MaritalStatus,
		BloodType:               employee.PersonalData.BloodType,
		Religion:                employee.PersonalData.Religion,
		IdentityType:            employee.PersonalData.IdentityType,
		IdentityNumber:          employee.PersonalData.IdentityNumber,
		IdentityExpiryDate:      employee.PersonalData.IdentityExpiryDate,
		PostalCode:              employee.PersonalData.PostalCode,
		CitizenIdAddress:        employee.PersonalData.CitizenIdAddress,
		ResidenceAddress:        employee.PersonalData.ResidenceAddress,
		EmployeeNumber:          employee.EmploymentData.EmployeeNumber,
		EmployeeStatus:          employee.EmploymentData.EmployeeStatus,
		JoinDate:                employee.EmploymentData.JoinDate,
		EndStatusDate:           employee.EmploymentData.EndStatusDate,
		Branch:                  employee.EmploymentData.Branch,
		Organization:            employee.EmploymentData.Organization,
		JobPosition:             employee.EmploymentData.JobPosition,
		JobLevel:                employee.EmploymentData.JobLevel,
		Grade:                   employee.EmploymentData.Grade,
		Class:                   employee.EmploymentData.Class,
		Schedule:                employee.EmploymentData.Schedule,
		BasicSalary:             employee.SalaryData.BasicSalary,
		SalaryType:              employee.SalaryData.SalaryType,
		PaymentSchedule:         employee.SalaryData.PaymentSchedule,
		ProrateSetting:          employee.SalaryData.ProrateSetting,
		AllowForOvertime:        employee.SalaryData.AllowForOvertime,
		BankName:                employee.SalaryData.BankName,
		BankAccountNumber:       employee.SalaryData.BankAccountNumber,
		BankAccountName:         employee.SalaryData.BankAccountName,
		NPWP:                    employee.SalaryData.NPWP,
		PTKPStatus:              employee.SalaryData.PTKPStatus,
		TaxMethod:               employee.SalaryData.TaxMethod,
		TaxSalary:               employee.SalaryData.TaxSalary,
		TaxableDate:             employee.SalaryData.TaxableDate,
		EmployeeTaxStatus:       employee.SalaryData.EmployeeTaxStatus,
		BeginningNetto:          employee.SalaryData.BeginningNetto,
		PHH21Paid:               employee.SalaryData.PHH21Paid,
		BPJSKetenagaKerjaNumber: employee.SalaryData.BPJSKetenagaKerjaNumber,
		NPPBPJSKetenagaKerja:    employee.SalaryData.NPPBPJSKetenagaKerja,
		BPJSKetenagaKerjaDate:   employee.SalaryData.BPJSKetenagaKerjaDate,
		BPJSKesehatanNumber:     employee.SalaryData.BPJSKesehatanNumber,
		BPJSKesehatanFamily:     employee.SalaryData.BPJSKesehatanFamily,
		BPJSKesehatanDate:       employee.SalaryData.BPJSKesehatanDate,
		BPJSKesehatanCost:       employee.SalaryData.BPJSKesehatanCost,
		JHTCost:                 employee.SalaryData.JHTCost,
		JaminanPensiunCost:      employee.SalaryData.JaminanPensiunCost,
		JaminanPensiunDate:      employee.SalaryData.JaminanPensiunDate,
	}

	return employeeResponse

}

func ToEmployeeResponses(employee []entity.Employee) []model.EmployeeResponse {
	var response []model.EmployeeResponse
	for _, employee := range employee {
		response = append(response, ToEmployeeResponse(employee))
	}
	return response
}

func ToUserResponse(user entity.User) model.UserResponse {
	userResponse := model.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return userResponse
}

func ToLoginResponse(user entity.User, token string) model.LoginResponse {
	loginResponse := model.LoginResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	return loginResponse
}

func ToAttendanceResponse(attendance entity.Attendance) model.AttendanceResponse {
	attendanceResponse := model.AttendanceResponse{
		Date:         attendance.Date,
		CheckIn:      attendance.CheckIn,
		CheckOut:     attendance.CheckOut,
		Shift:        attendance.Shift,
		OverTime:     attendance.OverTime,
		Status:       attendance.Status,
		EmployeeName: attendance.Employee.PersonalData.FirstName + " " + attendance.Employee.PersonalData.LastName,
		BasicSalary:  attendance.Employee.SalaryData.BasicSalary,
	}
	return attendanceResponse
}

func ToAttendanceResponses(attendance []entity.Attendance) []model.AttendanceResponse {
	var response []model.AttendanceResponse
	for _, attendance := range attendance {
		response = append(response, ToAttendanceResponse(attendance))
	}
	return response
}
