package router

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v3"
)

func addMailRoutes(r fiber.Router) {
	r.Post("/mail", handlers.SendMail)
}