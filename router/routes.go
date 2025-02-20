package router

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gofiber/fiber/v3"
)

func addMailRoutes(r fiber.Router) {
	r.Post("/mail", handlers.SendMail, middleware.VerifyAPIKey())
}

func addAuthRoutes(r fiber.Router) {
	r.Post("/signup", handlers.SignUp, middleware.ValidateAuthRequest())
	r.Post("/login", handlers.Login, middleware.ValidateAuthRequest())
	r.Get("/logout", handlers.Logout, middleware.JWTProtected())
	r.Post("/register-client", handlers.MakeNewClient, middleware.JWTProtected() ,middleware.ValidateClientRegisterRequest())
}
