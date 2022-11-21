package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"

	webhook "github.com/FlowingSPDG/Got5/examples/discord_webhook"
	"github.com/FlowingSPDG/Got5/route"
)

var (
	// ID Webhook ID
	ID uint64
	// Token Webhook Token
	Token string
)

func main() {
	flag.Uint64Var(&ID, "id", 0, "Webhook ID")
	flag.StringVar(&Token, "token", "", "Webhook Token")
	flag.Parse()

	evh := webhook.NewEventHandler(ID, Token)
	defer evh.Close()

	app := fiber.New()

	g5 := app.Group("/get5") // /get5
	if err := route.SetupEventHandlers(evh, g5); err != nil {
		panic(err)
	}

	app.Listen(":3000")
}
