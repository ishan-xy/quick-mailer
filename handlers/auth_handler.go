package handlers

import (
	"backend/common"
	"backend/database"
	"fmt"
	"log"
	_ "log"
	_ "time"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)



func SignUp(c fiber.Ctx) error {
	req := c.Locals("auth_request").(database.AuthRequest)

	_, exists, _ := database.UserDB.GetExists(bson.M{"email": req.Email})
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already registered",})
	}

	// hash password
	hash, err := common.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   utils.WithStack(err),
			"message": "Failed to hash password",
		})
	}
	// create user
	user := database.User{
		Email:    req.Email,
		Password: hash,
	}
	_, err = database.UserDB.InsertOne(c.Context(), user)
	if err != nil {
		return utils.WithStack(err)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c fiber.Ctx) error {
	req := c.Locals("auth_request").(database.AuthRequest)

	user, exists, _ := database.UserDB.GetExists(bson.M{"email": req.Email})
	if !exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email not registered"})
	}

	if !common.CheckPasswordHash(req.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect password"})
	}

	tokenString, err := common.GenerateJWT(user.Email)
	if err != nil {
		log.Println("Error generating JWT: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.WithStack(err))
	}
	c.Set("Set-Cookie", fmt.Sprintf("%s=%s; HttpOnly; SameSite=Lax", common.Cfg.CookieName,tokenString))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user":user, "token": fmt.Sprintf("Bearer %s", tokenString)})
}

func Logout(c fiber.Ctx) error {
	c.Set("Set-Cookie", fmt.Sprintf("%s=; HttpOnly; SameSite=Lax", common.Cfg.CookieName))
	return c.JSON(fiber.Map{"message": "Logout successful"})
}