package main

import (
	"TalentHub/config"
	"TalentHub/controller"
	"TalentHub/exception"
	"TalentHub/repository"
	"TalentHub/repository/seeder"
	"TalentHub/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {

	// Setup Configuration
	configuration := config.New()
	database := config.ConnectDB(configuration)
	err := config.Migration(database)
	exception.PanicIfNeeded(err)

	for _, seed := range seeder.All() {
		if err := seed.Run(database); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	//// Setup Repository
	employeeRepository := repository.NewEmployeeRepository()
	companayRepository := repository.NewCompanyRepository()
	userRepository := repository.NewUserRepository()
	//
	//// Setup Service
	employeeService := service.NewEmployeeService(database, &employeeRepository)
	companyService := service.NewCompanyService(database, &companayRepository)
	userService := service.NewUserService(database, &userRepository)
	//
	//// Setup Controller
	employeeController := controller.NewEmployeeController(&employeeService)
	companyController := controller.NewCompanyController(companyService)
	userController := controller.NewUserController(userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	employeeController.Route(app)
	companyController.Route(app)
	userController.Route(app)

	// Start App
	err = app.Listen(":" + configuration.Get("PORT"))
	exception.PanicIfNeeded(err)
}
