package controller

import (
	"TalentHub/entity"
	"TalentHub/exception"
	"TalentHub/middleware"
	"TalentHub/model"
	"TalentHub/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type AttendanceController struct {
	AttendanceService service.AttendanceService
	Middelware        middleware.Middelware
}

func NewAttendanceController(attendanceService *service.AttendanceService, middelware middleware.Middelware) AttendanceController {
	return AttendanceController{
		AttendanceService: *attendanceService,
		Middelware:        middelware,
	}
}

func (controller *AttendanceController) Route(app *fiber.App) {
	app.Post("/api/attendance", controller.Middelware.Protected(), controller.CreateAttendance)
	app.Get("/api/attendance/employee/:employee_id", controller.Middelware.Protected(), controller.GetAttendanceByEmployeeId)
	app.Get("/api/attendance/date/:date", controller.Middelware.Protected(), controller.GetAttendanceByDate)
	app.Get("/api/attendance/employee/:employee_id/date/:date", controller.Middelware.Protected(), controller.GetAttendanceByEmployeeIdAndDate)
	app.Get("/api/attendance/filter", controller.Middelware.Protected(), controller.GetAttendanceByFilter)
}

func (controller *AttendanceController) CreateAttendance(c *fiber.Ctx) error {
	var request entity.Attendance
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	_ = controller.AttendanceService.CreateAttendance(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "success",
	})
}

func (controller *AttendanceController) GetAttendanceByEmployeeId(c *fiber.Ctx) error {
	employeeId := c.Params("employee_id")
	employeeIdInt, err := strconv.Atoi(employeeId)
	exception.PanicIfNeeded(err)

	attendanceResult := controller.AttendanceService.GetAttendanceByEmployeeId(uint(employeeIdInt))

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   attendanceResult,
	})
}

func (controller *AttendanceController) GetAttendanceByDate(c *fiber.Ctx) error {
	//date := c.Params("date")

	//attendanceResult := controller.AttendanceService.GetAttendanceByDate(date)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "attendanceResult",
	})
}

func (controller *AttendanceController) GetAttendanceByEmployeeIdAndDate(c *fiber.Ctx) error {
	//employeeId := c.Params("employee_id")
	//employeeIdInt, err := strconv.Atoi(employeeId)
	//exception.PanicIfNeeded(err)
	//
	//date := c.Params("date")

	//attendanceResult := controller.AttendanceService.GetAttendanceByEmployeeIdAndDate(uint(employeeIdInt), date)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "attendanceResult",
	})
}

func (controller *AttendanceController) GetAttendanceByFilter(c *fiber.Ctx) error {
	//var request entity.Attendance
	//err := c.BodyParser(&request)
	//exception.PanicIfNeeded(err)
	//
	//attendanceResult := controller.AttendanceService.GetAttendanceByFilter(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "attendanceResult",
	})
}
