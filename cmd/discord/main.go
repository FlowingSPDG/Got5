package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller/discord"
	"github.com/FlowingSPDG/Got5/route"
)

var (
	discordToken = ""
	hostName     = "localhost"
	port         = "8080"
)

func main() {
	// Parse -projectID flag
	flag.StringVar(&discordToken, "token", "", "Discord BOT Token")
	flag.StringVar(&hostName, "hostname", "", "Web hostname")
	flag.StringVar(&port, "port", "8080", "Port to listen")
	flag.Parse()

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	// Get Controller for Discord BOT
	ctx := context.Background()
	ctrl, err := discord.NewDiscordController(ctx, discordToken, discord.ControllerSetting{
		Hostname: hostName,
		Port:     portInt,
	})
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

	p := net.JoinHostPort(hostName, port)

	// Start server
	log.Println("Start listening on:", p)
	if err := app.Listen(p); err != nil {
		panic(err)
	}
}
