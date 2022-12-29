package middleware

import (
	"TalentHub/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type Middelware struct {
	configuration config.Config
}

func NewMiddelware(configuration config.Config) Middelware {
	return Middelware{configuration: configuration}
}

func (middleware *Middelware) Protected() func(*fiber.Ctx) error {

	return jwtware.New(jwtware.Config{
		SigningKey:   middleware.configuration.JwtKey(),
		ErrorHandler: jwtError,
	})

}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}
