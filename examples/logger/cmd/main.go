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

	app := fiber.New()

	g5 := app.Group("/get5") // /get5
	if err := route.SetupAllGet5Handlers(evh, loader, uploader, g5); err != nil {
		panic(err)
	}

	app.Listen(":3000")
}
