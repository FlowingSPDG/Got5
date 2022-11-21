package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"

	webhook "github.com/FlowingSPDG/Got5/examples/discord_webhook"
	"github.com/FlowingSPDG/Got5/route"
)

var (
	// ID Webhook ID
	ID uint64
	// Token Webhook Token
	Token string

	// Auth Auth for game server
	// Currently unavailable
	Auth string = ""

	// Host e.g. "localhost"
	Host string

	// Port port to listen
	Port int
)

func main() {
	flag.Uint64Var(&ID, "id", 0, "Webhook ID")
	flag.StringVar(&Token, "token", "", "Webhook Token")
	// flag.StringVar(&Auth, "auth", "", "Password for game server event")
	flag.StringVar(&Host, "host", "localhost", "hostname")
	flag.IntVar(&Port, "port", 3000, "Port to listen")
	flag.Parse()

	evh := webhook.NewEventHandler(ID, Token, fmt.Sprintf("http://%s:%d/get5/event", Host, Port), Auth)
	defer evh.Close()

	app := fiber.New()

	g5 := app.Group("/get5") // /get5
	if err := route.SetupEventHandlers(evh, g5); err != nil {
		panic(err)
	}

	if err := app.Listen(fmt.Sprintf(":%d", Port)); err != nil {
		panic(err)
	}
}
