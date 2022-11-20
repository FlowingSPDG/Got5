package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
)

// CheckMatchLoaderAuth 認証用ハンドラ
func CheckMatchLoaderAuth(loader controller.MatchLoader) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		mid := c.Params("matchID")
		reqAuthVal := c.Get("Authorization")
		if err := loader.CheckMatchAuth(c.Context(), mid, reqAuthVal); err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Authorization Invalid")
		}
		return c.Next()
	})
}

// LoadMatchHandler GET on get5_loadmatch_url "https://example.com/match_config.json" "Authorization" "Bearer <token>"
func LoadMatchHandler(loader controller.MatchLoader) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// マッチを取得、JSONを組んで返す
		mid := c.Params("matchID")
		m, err := loader.Load(c.Context(), mid)
		if err != nil {
			// TODO:エラーハンドリングをする
			return c.Status(fiber.StatusNotFound).SendString("match not found")
		}
		return c.JSON(m.ToG5Format())
	})
}
