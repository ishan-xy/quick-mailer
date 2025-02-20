package handlers

import (
	"backend/common"
	"backend/database"
	_ "log"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func MakeNewClient(c fiber.Ctx) error {
	req := c.Locals("client_request").(database.RegisterClientRequest)

	_, exists, _ := database.ClientDB.GetExists(bson.M{"email": req.Email})
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already registered",})
	}

	_, exists, _ = database.ClientDB.GetExists(bson.M{"name": req.Name})
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name already registered",})
	}
	apikey, serialnumber, err := common.GenerateAPIKey(common.Cfg.API_Secret)
	if err!=nil{
		return utils.WithStack(err)
	}

	client := database.Client{
		Name: req.Name,
		Email: req.Email,
		SerialNumber: serialnumber,
		ClientSecret: apikey,
	}

	_, err = database.ClientDB.InsertOne(c.Context(), client)
	if err != nil {
		return utils.WithStack(err)
	}

	return c.Status(fiber.StatusCreated).JSON(client)
}