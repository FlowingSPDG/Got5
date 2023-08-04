package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"

	webhook "github.com/FlowingSPDG/Got5/examples/discord_webhook"
	fiberroute "github.com/FlowingSPDG/Got5/route/fiber"
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

	// Password password for game server
	Password string
)

func main() {
	flag.Uint64Var(&ID, "id", 0, "Webhook ID")
	flag.StringVar(&Token, "token", "", "Webhook Token")
	// flag.StringVar(&Auth, "auth", "", "Password for game server event")
	flag.StringVar(&Host, "host", "localhost", "hostname")
	flag.IntVar(&Port, "port", 3000, "Port to listen")
	flag.StringVar(&Password, "password", "my_password", "Password for game server")
	flag.Parse()

	auth := webhook.NewAuth(Password)

	evh := webhook.NewEventHandler(ID, Token, fmt.Sprintf("http://%s:%d/get5/event", Host, Port), Auth)
	defer evh.Close()

	app := fiber.New()

	g5 := app.Group("/get5") // /get5
	fiberroute.SetupEventHandlers(evh, auth, g5)

	if err := app.Listen(fmt.Sprintf(":%d", Port)); err != nil {
		panic(err)
	}
}
