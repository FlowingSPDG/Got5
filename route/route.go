package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
)

func SetupAllGet5Handlers(ctrl controller.Controller, r fiber.Router, bucket string) error {
	if err := SetupMatchLoadHandler(ctrl, r); err != nil {
		return err
	}
	if err := SetupEventHandlers(ctrl, r); err != nil {
		return err
	}
	if err := SetupDemoUploadHandler(ctrl, r, bucket); err != nil {
		return err
	}
	return nil
}

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler(ctrl controller.Controller, r fiber.Router, bucket string) error {
	r.Post("/match/:matchID/demo", DemoUploadHandler(ctrl, bucket))
	return nil
}

// SetupMatchLoadHandlers Setup get5 loadmatch handler
func SetupMatchLoadHandler(ctrl controller.Controller, r fiber.Router) error {
	r.Get("/match/:matchID", LoadMatchHandler(ctrl))
	return nil
}

// SetupRouter Setup get5 handlers to specified fiber.Router
func SetupEventHandlers(ctrl controller.Controller, r fiber.Router) error {
	r.Post("/event", OnEventHandler(ctrl))
	return nil
}
