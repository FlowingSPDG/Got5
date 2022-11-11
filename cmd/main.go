package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/route"
)

func main() {
	app := fiber.New()

	g5 := app.Group("/get5") // /get5
	g5.Post("/Get5_OnEvent", route.OnEventHandler())

	app.Listen(":3000")
}
