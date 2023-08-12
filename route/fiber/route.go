package fiberroute

import (
	got5 "github.com/FlowingSPDG/Got5"
	"github.com/gofiber/fiber/v2"
)

// SetupAllGet5Handlers Setup All Get5-related handlers
func SetupAllGet5Handlers(evh got5.EventHandler, loader got5.MatchLoader, uploader got5.DemoUploader, auth got5.Auth, r fiber.Router) {
	SetupEventHandlers(evh, auth, r)
	SetupMatchLoadHandler(loader, auth, r)
	SetupDemoUploadHandler(uploader, auth, r)
}

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler(uploader got5.DemoUploader, auth got5.Auth, r fiber.Router) {
	r.Post("/demo", CheckDemoAuth(auth), DemoUploadHandler(uploader))
}

// SetupMatchLoadHandler Setup get5 loadmatch handler
func SetupMatchLoadHandler(loader got5.MatchLoader, auth got5.Auth, r fiber.Router) {
	r.Get("/match/:matchID", CheckMatchLoaderAuth(auth), LoadMatchHandler(loader))
}

// SetupEventHandlers get5 handlers to specified fiber.Router
func SetupEventHandlers(ctrl got5.EventHandler, auth got5.Auth, r fiber.Router) {
	r.Post("/event", CheckEventHandlerAuth(auth), OnEventHandler(ctrl))
}
