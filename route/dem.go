package route

import (
	"bytes"
	"strconv"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/gofiber/fiber/v2"
)

// CheckDemoAuth 認証用ハンドラ
func CheckDemoAuth(uploader controller.DemoUploader) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error { // Verifyをかける
		filename := c.Get("Get5-DemoName")
		matchID := c.Get("Get5-MatchId")

		mapNumStr := c.Get("Get5-MapNumber")
		mapNum, err := strconv.Atoi(mapNumStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // カスタムエラーを返したい
		}

		serverIDstr := c.Get("Get5-ServerId")
		serverID, err := strconv.Atoi(serverIDstr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // カスタムエラーを返したい
		}

		auth := c.Get("Get5-Authorization")
		if err := uploader.CheckDemoAuth(c.Context(), filename, matchID, mapNum, serverID, auth); err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Not verified") // カスタムエラーを返したい
		}
		return c.Next()
	})
}

// DemoUploadHandler POST CS:GO dem file.
// アップロードされたdemファイルを制御するハンドラ
func DemoUploadHandler(uploader controller.DemoUploader) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// アップロードを実施
		br := bytes.NewBuffer(c.Body())

		filename := c.Get("Get5-DemoName")
		matchID := c.Get("Get5-MatchId")

		if err := uploader.Upload(c.Context(), matchID, filename, br); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // カスタムエラーを返したい
		}
		return nil
	})
}
