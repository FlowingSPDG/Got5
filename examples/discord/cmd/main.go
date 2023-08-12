package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/examples/discord"
	fiberroute "github.com/FlowingSPDG/Got5/route/fiber"
)

var (
	discordToken = ""
	hostName     = "localhost"
	port         = "8080"
	password     = ""
)

func main() {
	// Parse -projectID flag
	flag.StringVar(&discordToken, "token", "", "Discord BOT Token")
	flag.StringVar(&hostName, "hostname", "", "Web hostname")
	flag.StringVar(&port, "port", "8080", "Port to listen")
	flag.StringVar(&password, "password", "password", "Password to auth")
	flag.Parse()

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	// Get Controller for Discord BOT
	ctx := context.Background()
	ctrl, err := discord.NewDiscord(ctx, discordToken, discord.ControllerSetting{
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
	fiberroute.SetupEventHandlers(ctrl, ctrl, g5)
	fiberroute.SetupMatchLoadHandler(ctrl, ctrl, g5)

	p := net.JoinHostPort(hostName, port)

	// Start server
	log.Println("Start listening on:", p)
	if err := app.Listen(p); err != nil {
		panic(err)
	}
}
