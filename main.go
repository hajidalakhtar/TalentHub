package main

import (
	"TalentHub/config"
	"TalentHub/controller"
	"TalentHub/exception"
	"TalentHub/repository"
	"TalentHub/service"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// Setup Configuration
	configuration := config.New()
	database := config.ConnectDB(configuration)
	//err := config.Migration(database)
	//exception.PanicIfNeeded(err)
	//
	//for _, seed := range seeder.All() {
	//	if err := seed.Run(database); err != nil {
	//		log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
	//	}
	//}

	//// Setup Repository
	employeeRepository := repository.NewEmployeeRepository()
	companayRepository := repository.NewCompanyRepository()
	userRepository := repository.NewUserRepository()
	attendanceRepository := repository.NewAttendanceRepository()
	//
	//// Setup Service
	employeeService := service.NewEmployeeService(database, &employeeRepository)
	companyService := service.NewCompanyService(database, &companayRepository)
	userService := service.NewUserService(database, &userRepository)
	attendanceService := service.NewAttendanceService(database, &attendanceRepository)

	//// Setup Controller
	employeeController := controller.NewEmployeeController(&employeeService)
	companyController := controller.NewCompanyController(companyService)
	userController := controller.NewUserController(userService)
	attendanceController := controller.NewAttendanceController(&attendanceService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	employeeController.Route(app)
	companyController.Route(app)
	userController.Route(app)
	attendanceController.Route(app)

	// Start App

	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	portStr := ":" + configuration.Get("PORT")
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		err := app.Listen(portStr)
		exception.PanicIfNeeded(err)

	}

}
