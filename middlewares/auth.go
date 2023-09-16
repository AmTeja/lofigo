package middlewares

import (
	"github.com/amteja/lofig/env"
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   []byte(env.Get("JWT_SECRET")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Missing or malformed JWT",
		})
	}

	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"message": "Invalid or expired JWT",
	})
}
