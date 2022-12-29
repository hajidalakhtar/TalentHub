package main

import (
	"TalentHub/config"
	"TalentHub/controller"
	"TalentHub/exception"
	"TalentHub/middleware"
	"TalentHub/repository"
	"TalentHub/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {

	// Setup Configuration
	configuration := config.New()
	database := config.ConnectDB(configuration)
	middelware := middleware.NewMiddelware(configuration)
	//middleware := mi
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
	userService := service.NewUserService(database, &userRepository, configuration)
	attendanceService := service.NewAttendanceService(database, &attendanceRepository)

	//// Setup Controller
	employeeController := controller.NewEmployeeController(&employeeService, middelware)
	companyController := controller.NewCompanyController(companyService, middelware)
	userController := controller.NewUserController(userService, middelware)
	attendanceController := controller.NewAttendanceController(&attendanceService, middelware)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Setup Routing
	employeeController.Route(app)
	companyController.Route(app)
	userController.Route(app)
	attendanceController.Route(app)

	// Start App

	err := app.Listen(getPort())
	exception.PanicIfNeeded(err)

}
