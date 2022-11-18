package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
)

// SetupAllGet5Handlers Setup All Get5-related handlers
func SetupAllGet5Handlers(evh controller.EventHandler, loader controller.MatchLoader, uploader controller.DemoUploader, r fiber.Router) error {
	if err := SetupEventHandlers(evh, r); err != nil {
		return err
	}
	if err := SetupMatchLoadHandler(loader, r); err != nil {
		return err
	}
	if err := SetupDemoUploadHandler(uploader, r); err != nil {
		return err
	}
	return nil
}

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler(uploader controller.DemoUploader, r fiber.Router) error {
	r.Post("/demo", DemoUploadHandler(uploader))
	return nil
}

// SetupMatchLoadHandler Setup get5 loadmatch handler
func SetupMatchLoadHandler(loader controller.MatchLoader, r fiber.Router) error {
	r.Get("/match/:matchID", LoadMatchHandler(loader))
	return nil
}

// SetupEventHandlers get5 handlers to specified fiber.Router
func SetupEventHandlers(ctrl controller.EventHandler, r fiber.Router) error {
	r.Post("/event", OnEventHandler(ctrl))
	return nil
}
