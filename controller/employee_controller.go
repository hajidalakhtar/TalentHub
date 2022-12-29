package controller

import (
	"TalentHub/exception"
	"TalentHub/middleware"
	"TalentHub/model"
	"TalentHub/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
	Middelware      middleware.Middelware
}

func NewEmployeeController(employeeService *service.EmployeeService, middelware middleware.Middelware) EmployeeController {
	return EmployeeController{EmployeeService: *employeeService, Middelware: middelware}
}

func (controller *EmployeeController) Route(app *fiber.App) {
	app.Post("/api/employee/:company_id", controller.Middelware.Protected(), controller.CreateEmployee)
	app.Get("/api/employee/:company_id", controller.Middelware.Protected(), controller.GetEmployeeByCompany)

}

func (controller *EmployeeController) CreateEmployee(c *fiber.Ctx) error {
	var request model.CreateEmployeeRequest

	companyId := c.Params("company_id")
	companyIdInt, err := strconv.Atoi(companyId)

	err = c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	_ = controller.EmployeeService.CreateEmployee(request, uint(companyIdInt))

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "success",
	})

}

func (controller *EmployeeController) GetEmployeeByCompany(c *fiber.Ctx) error {
	companyId := c.Params("company_id")
	companyIdInt, err := strconv.Atoi(companyId)
	exception.PanicIfNeeded(err)
	response := controller.EmployeeService.GetEmployeeByCompany(uint(companyIdInt))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
