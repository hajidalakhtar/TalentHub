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

type UserController struct {
	UserService service.UserService
	Middelware  middleware.Middelware
}

func NewUserController(userService service.UserService, middelware middleware.Middelware) *UserController {
	return &UserController{UserService: userService, Middelware: middelware}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/api/user", controller.Middelware.Protected(), controller.CreateUser)
	app.Post("/api/user/login", controller.Login)
	app.Get("/api/user/:user_id", controller.Middelware.Protected(), controller.GetUserById)
}

func (controller *UserController) CreateUser(c *fiber.Ctx) error {
	var request entity.User
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := controller.UserService.CreateUser(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var request entity.User
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, isLoginSuccess := controller.UserService.Login(request.Email, request.Password)

	if isLoginSuccess {
		return c.JSON(model.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   response,
		})
	} else {
		return c.JSON(model.WebResponse{
			Code:   401,
			Status: "UNAUTHORIZED",
			Data:   "Email or Password is wrong",
		})
	}

}

func (controller *UserController) GetUserById(c *fiber.Ctx) error {

	userId := c.Params("user_id")
	userIdInt, err := strconv.Atoi(userId)
	exception.PanicIfNeeded(err)
	response := controller.UserService.FindUserById(uint(userIdInt))

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
