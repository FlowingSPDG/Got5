package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller/logger"
	"github.com/FlowingSPDG/Got5/route"
)

func main() {
	ctrl := logger.NewLoggerController()
	defer ctrl.Close()

	app := fiber.New()

	g5 := app.Group("/get5") // /get5
	if err := route.SetupAllGet5Handlers(ctrl, g5, ""); err != nil {
		panic(err)
	}

	app.Listen(":3000")
}
