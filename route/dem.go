package route

import (
	"github.com/FlowingSPDG/Got5/controller"
	"github.com/gofiber/fiber/v2"
)

// アップロードされたdemファイルを制御するハンドラ
// DemoUploadHandler POST CS:GO dem file
func DemoUploadHandler(ctrl controller.Controller, bucket string) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// TODO: ヘッダーを検証し不正であれば拒否する
		// Controllerに渡してアップロードを実施
		filename := c.Get("Get5-DemoName")
		matchID := c.Get("Get5-MatchId")
		// mapNumber := c.Get("Get5-MapNumber")
		// serverID := c.Get("Get5-ServerId")

		if err := ctrl.RegisterDemoFile(c.Context(), bucket, matchID, filename, c.Body()); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // カスタムエラーを返したい
		}
		return nil
	})
}