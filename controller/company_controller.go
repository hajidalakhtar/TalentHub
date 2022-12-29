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

type CompanyController struct {
	CompanyService service.CompanyService
	Middelware     middleware.Middelware
}

func NewCompanyController(companyService service.CompanyService, middleware middleware.Middelware) CompanyController {
	return CompanyController{
		CompanyService: companyService,
		Middelware:     middleware,
	}
}

func (controller *CompanyController) Route(app *fiber.App) {
	app.Post("/api/company", controller.Middelware.Protected(), controller.CreateCompany)
	app.Get("/api/company/:company_id", controller.Middelware.Protected(), controller.GetCompanyById)

}

func (controller *CompanyController) CreateCompany(c *fiber.Ctx) error {
	var request entity.Company
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	_ = controller.CompanyService.CreateCompany(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "success",
	})

}

func (controller *CompanyController) GetCompanyById(c *fiber.Ctx) error {

	companyId := c.Params("company_id")
	companyIdInt, err := strconv.Atoi(companyId)
	exception.PanicIfNeeded(err)

	companyResult := controller.CompanyService.GetCompanyById(uint(companyIdInt))

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   companyResult,
	})
}
