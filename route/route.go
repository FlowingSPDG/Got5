package route

import (
	"fmt"

	"github.com/FlowingSPDG/Got5/models"
	"github.com/gofiber/fiber/v2"
)

// OnEventHandler POST on /Get5_OnEvent
func OnEventHandler() func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnEventPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		fmt.Println("Body:", p)
		return nil
	})
}
