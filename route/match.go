package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
)

// LoadMatchHandler GET on get5_loadmatch_url "https://example.com/match_config.json" "Authorization" "Bearer <token>"
func LoadMatchHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// TODO: 必要であればJWTを検証する

		// マッチを取得、JSONを組んで返す
		// 注意: c.Params はリクエストの間でしか有効ではないので注意
		// 参照: https://docs.gofiber.io/#zero-allocation
		mid := c.Params("matchID")
		m, err := ctrl.GetMatch(c.Context(), mid)
		if err != nil {
			// TODO:エラーハンドリングをする
			return c.Status(fiber.StatusNotFound).SendString("match not found")
		}
		return c.JSON(m.ToG5Format())
	})
}
