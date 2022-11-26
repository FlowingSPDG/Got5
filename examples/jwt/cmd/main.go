package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/najeira/randstr"

	"github.com/FlowingSPDG/Got5/examples/jwt"
	"github.com/FlowingSPDG/Got5/route"
)

func main() {
	secret := randstr.String(512) // 512bit-long secret
	loader := jwt.NewMatchLoader(secret)
	auth := jwt.NewAuth(secret)

	app := fiber.New()

	g5 := app.Group("/get5") // /get5
	route.SetupMatchLoadHandler(loader, auth, g5)

	app.Listen(":3000")
}
