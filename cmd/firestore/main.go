package main

import (
	"context"
	"flag"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"

	fsc "github.com/FlowingSPDG/Got5/controller/firebase"
	"github.com/FlowingSPDG/Got5/route"
)

var (
	projectID = ""
	bucket    = ""
)

func main() {
	// Parse -projectID flag
	flag.StringVar(&projectID, "projectID", "", "Firebase project ID")
	flag.StringVar(&bucket, "bucket", "", "Firebase Storage Bucket")
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

	// Get Controller connected to firestore
	ctrl, err := fsc.NewFirebaseController(ctx, fb)
	if err != nil {
		log.Fatalln(err)
	}
	defer ctrl.Close()

	// Setup fiber
	app := fiber.New()
	g5 := app.Group("/get5") // /get5
	if err := route.SetupAllGet5Handlers(ctrl, g5, bucket); err != nil {
		panic(err)
	}

	// Start server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}