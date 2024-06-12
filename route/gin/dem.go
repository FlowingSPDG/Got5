package ginroute

import (
	"net/http"
	"strconv"

	got5 "github.com/FlowingSPDG/Got5"
	"github.com/gin-gonic/gin"
)

// CheckDemoAuth 認証用ハンドラ
func CheckDemoAuth[TMatchID int](auth got5.Auth[TMatchID]) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Verifyをかける
		filename := c.GetHeader("Get5-FileName")
		matchID := c.GetHeader("Get5-MatchId")

		mid, err := strconv.Atoi(matchID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		mapNumStr := c.GetHeader("Get5-MapNumber")
		mapNum := 0
		if mapNumStr != "" {
			mapNum, err = strconv.Atoi(mapNumStr)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}

		serverID := c.GetHeader("Get5-ServerId")

		reqAuth := c.GetHeader("Authorization")
		if err := auth.CheckDemoAuth(c, TMatchID(mid), filename, mapNum, serverID, reqAuth); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err) // カスタムエラーを返したい
			return
		}
		c.Next()
	}
}

// DemoUploadHandler POST CS:GO dem file.
// アップロードされたdemファイルを制御するハンドラ
func DemoUploadHandler[TMatchID int](uploader got5.DemoUploader[TMatchID]) func(c *gin.Context) {
	return (func(c *gin.Context) {
		// アップロードを実施

		filename := c.GetHeader("Get5-FileName")
		matchID := c.GetHeader("Get5-MatchId")

		mid, err := strconv.Atoi(matchID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := uploader.Upload(c, TMatchID(mid), filename, c.Request.Body); err != nil {
			c.String(http.StatusInternalServerError, err.Error()) // カスタムエラーを返したい
		}
		c.String(http.StatusOK, "OK")
	})
}
