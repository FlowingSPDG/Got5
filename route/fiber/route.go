package fiberroute

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
)

// SetupAllGet5Handlers Setup All Get5-related handlers
func SetupAllGet5Handlers(evh controller.EventHandler, loader controller.MatchLoader, uploader controller.DemoUploader, auth controller.Auth, r fiber.Router) {
	SetupEventHandlers(evh, auth, r)
	SetupMatchLoadHandler(loader, auth, r)
	SetupDemoUploadHandler(uploader, auth, r)
}

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler(uploader controller.DemoUploader, auth controller.Auth, r fiber.Router) {
	r.Post("/demo", CheckDemoAuth(auth), DemoUploadHandler(uploader))
}

// SetupMatchLoadHandler Setup get5 loadmatch handler
func SetupMatchLoadHandler(loader controller.MatchLoader, auth controller.Auth, r fiber.Router) {
	r.Get("/match/:matchID", CheckMatchLoaderAuth(auth), LoadMatchHandler(loader))
}

// SetupEventHandlers get5 handlers to specified fiber.Router
func SetupEventHandlers(ctrl controller.EventHandler, auth controller.Auth, r fiber.Router) {
	r.Post("/event", CheckEventHandlerAuth(auth), OnEventHandler(ctrl))
}
