package main

import (
	"context"
	"flag"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"

	fsc "github.com/FlowingSPDG/Got5/controller/firestore"
	"github.com/FlowingSPDG/Got5/route"
)

var (
	projectID = ""
)

func main() {
	// Parse -projectID flag
	flag.StringVar(&projectID, "projectID", "", "Firestore project ID")
	flag.Parse()

	// Get Firebase service
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: projectID,
	}
	fb, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	// Get firestore
	client, err := fb.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// Get Controller connected to firestore
	ctrl := fsc.NewFirestoreController(client)

	// Setup fiber
	app := fiber.New()
	g5 := app.Group("/get5") // /get5
	if err := route.SetupRouter(ctrl, g5); err != nil {
		panic(err)
	}

	// Start server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
