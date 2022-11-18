package route

import (
	"github.com/FlowingSPDG/Got5/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// DemoUploadHandler POST CS:GO dem file
// アップロードされたdemファイルを制御するハンドラ
func DemoUploadHandler(uploader controller.DemoUploader) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// TODO: ヘッダーを検証し不正であれば拒否する
		// Controllerに渡してアップロードを実施
		filename := utils.CopyString(c.Get("Get5-DemoName"))
		matchID := utils.CopyString(c.Get("Get5-MatchId"))
		// mapNumber := c.Get("Get5-MapNumber")
		// serverID := c.Get("Get5-ServerId")

		if err := uploader.Upload(c.Context(), matchID, filename, c.Body()); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // カスタムエラーを返したい
		}
		return nil
	})
}
