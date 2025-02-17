package router

import (
	"errors"
	"fmt"
	"log"
	"time"

	"backend/common"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v3/middleware/recover"
)

func Init() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(utils.WithStack(errors.New("Error initializing router: " + fmt.Sprint(err))))
		}
	}()

	utils.SetErrorStackTrace(common.IsDebug)

	app := fiber.New(fiber.Config{
		CaseSensitive:      true,
		Concurrency:        1024 * 1024,
		IdleTimeout:        30 * time.Second,
		DisableDefaultDate: true,
		JSONEncoder:        json.Marshal,
		JSONDecoder:        json.Unmarshal,
	})

	app.Use(cors.New())
	app.Use(fiberRecover.New(fiberRecover.Config{EnableStackTrace: true}))
	app.Use(logger.New())
	log.Println("Default logging enabled")

	addMailRoutes(app)
	addAuthRoutes(app)

	log.Fatal(
		app.Listen("127.0.0.1:8080", fiber.ListenConfig{
			EnablePrintRoutes: true,
		}),
	)
}