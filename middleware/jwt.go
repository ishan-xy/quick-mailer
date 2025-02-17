package middleware

import (
	"backend/common"

	"log"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/gofiber/fiber/v3"
)

func JWTProtected() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := c.Cookies(common.Cfg.CookieName)
		log.Println("Token: ", tokenString)
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		token, err := common.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			log.Printf("Error validating token: %v", utils.WithStack(err))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		c.Locals("user", token)
		return c.Next()
	}
}
