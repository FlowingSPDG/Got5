package fiberroute

import (
	"bytes"
	"strconv"

	got5 "github.com/FlowingSPDG/Got5"
	"github.com/gofiber/fiber/v2"
)

// CheckDemoAuth 認証用ハンドラ
func CheckDemoAuth(auth got5.Auth) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error { // Verifyをかける
		filename := c.Get("Get5-FileName")
		matchID := c.Get("Get5-MatchId")

		mapNumStr := c.Get("Get5-MapNumber")
		mapNum := 0
		var err error
		if mapNumStr != "" {
			mapNum, err = strconv.Atoi(mapNumStr)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // カスタムエラーを返したい
			}
		}

		serverID := c.Get("Get5-ServerId")

		reqAuth := c.Get("Authorization")
		if err := auth.CheckDemoAuth(c.Context(), filename, matchID, mapNum, serverID, reqAuth); err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error()) // カスタムエラーを返したい
		}
		return c.Next()
	})
}

// DemoUploadHandler POST CS:GO dem file.
// アップロードされたdemファイルを制御するハンドラ
func DemoUploadHandler(uploader got5.DemoUploader) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// アップロードを実施
		br := bytes.NewBuffer(c.Body())

		filename := c.Get("Get5-FileName")
		matchID := c.Get("Get5-MatchId")

		if err := uploader.Upload(c.Context(), matchID, filename, br); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // カスタムエラーを返したい
		}
		return nil
	})
}
