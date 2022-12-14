package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/examples/logger"
	"github.com/FlowingSPDG/Got5/route"
)

func main() {
	evh := logger.NewEventHandler()
	defer evh.Close()

	// db := logger.NewDatabase()
	loader := logger.NewMatchLoader()
	uploader := logger.NewDemoUploader()
	auth := logger.NewAuth()

	app := fiber.New()

	g5 := app.Group("/get5") // /get5
	route.SetupAllGet5Handlers(evh, loader, uploader, auth, g5)

	app.Listen(":3000")
}
