package main

import (
	"context"
	"flag"
	"log"
	"net"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"

	fsc "github.com/FlowingSPDG/Got5/examples/firebase"
	"github.com/FlowingSPDG/Got5/route"
)

var (
	projectID = ""
	bucket    = ""
	hostName  = "localhost"
	port      = "8080"
)

func main() {
	// Parse -projectID flag
	flag.StringVar(&projectID, "projectID", "", "Firebase project ID")
	flag.StringVar(&bucket, "bucket", "", "Firebase Storage Bucket")
	flag.StringVar(&hostName, "hostname", "localhost", "Web hostname")
	flag.StringVar(&port, "port", "8080", "Port to listen")
	flag.Parse()

	// Get Firebase service
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID:     projectID,
		StorageBucket: bucket,
	}
	fb, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	// Get Event Handler
	evh, err := fsc.NewEventHandler(ctx, fb)
	if err != nil {
		log.Fatalln(err)
	}
	defer evh.Close()

	loader, err := fsc.NewMatchLoader(ctx, fb)
	if err != nil {
		log.Fatalln(err)
	}

	demUploader, err := fsc.NewDemoUploader(ctx, fb)
	if err != nil {
		log.Fatalln(err)
	}

	// Get Match Loader

	// Setup fiber
	app := fiber.New()
	g5 := app.Group("/get5") // /get5
	if err := route.SetupAllGet5Handlers(evh, loader, demUploader, g5); err != nil {
		panic(err)
	}

	p := net.JoinHostPort(hostName, port)

	// Start server
	log.Println("Start listening on:", p)
	if err := app.Listen(p); err != nil {
		panic(err)
	}
}
