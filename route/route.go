package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
)

// SetupAllGet5Handlers Setup All Get5-related handlers
func SetupAllGet5Handlers(evh controller.EventHandler, loader controller.MatchLoader, uploader controller.DemoUploader, auth controller.Auth, r fiber.Router) error {
	if err := SetupEventHandlers(evh, auth, r); err != nil {
		return err
	}
	if err := SetupMatchLoadHandler(loader, auth, r); err != nil {
		return err
	}
	if err := SetupDemoUploadHandler(uploader, auth, r); err != nil {
		return err
	}
	return nil
}

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler(uploader controller.DemoUploader, auth controller.Auth, r fiber.Router) error {
	r.Post("/demo", CheckDemoAuth(auth), DemoUploadHandler(uploader))
	return nil
}

// SetupMatchLoadHandler Setup get5 loadmatch handler
func SetupMatchLoadHandler(loader controller.MatchLoader, auth controller.Auth, r fiber.Router) error {
	r.Get("/match/:matchID", CheckMatchLoaderAuth(auth), LoadMatchHandler(loader))
	return nil
}

// SetupEventHandlers get5 handlers to specified fiber.Router
func SetupEventHandlers(ctrl controller.EventHandler, auth controller.Auth, r fiber.Router) error {
	r.Post("/event", CheckEventHandlerAuth(auth), OnEventHandler(ctrl))
	return nil
}
