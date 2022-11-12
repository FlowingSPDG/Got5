package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller/discord"
	"github.com/FlowingSPDG/Got5/route"
)

var (
	discordToken = ""
)

func main() {
	// Parse -projectID flag
	flag.StringVar(&discordToken, "token", "", "Discord BOT Token")
	flag.Parse()

	// Get Controller for Discord BOT
	ctrl, err := discord.NewDiscordController(discordToken)
	if err != nil {
		log.Fatalln(err)
	}
	defer ctrl.Close()

	// Setup fiber
	app := fiber.New()
	g5 := app.Group("/get5") // /get5
	if err := route.SetupEventHandlers(ctrl, g5); err != nil {
		panic(err)
	}
	if err := route.SetupMatchLoadHandler(ctrl, g5); err != nil {
		panic(err)
	}

	// Start server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
