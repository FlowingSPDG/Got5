package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"

	fsc "github.com/FlowingSPDG/Got5/controller/firebase"
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

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

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
	ctrl, err := fsc.NewFirebaseController(ctx, fb, fsc.ControllerSetting{
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
	if err := route.SetupAllGet5Handlers(ctrl, g5, bucket); err != nil {
		panic(err)
	}

	p := net.JoinHostPort(hostName, port)

	// Start server
	log.Println("Start listening on:", p)
	if err := app.Listen(p); err != nil {
		panic(err)
	}
}
