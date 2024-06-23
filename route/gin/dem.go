package ginroute

import (
	"net/http"
	"strconv"

	got5 "github.com/FlowingSPDG/Got5"
	"github.com/gin-gonic/gin"
)

// CheckDemoAuth 認証用ハンドラ
func CheckDemoAuth(auth got5.Auth) func(c *gin.Context) {
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
		if err := auth.CheckDemoAuth(c, int(mid), filename, mapNum, serverID, reqAuth); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err) // カスタムエラーを返したい
			return
		}
		c.Next()
	}
}

// DemoUploadHandler POST CS2 dem file.
// アップロードされたdemファイルを制御するハンドラ
func DemoUploadHandler(uploader got5.DemoUploader) func(c *gin.Context) {
	return (func(c *gin.Context) {
		// アップロードを実施

		// old headers: Get5-FileName, Get5-MatchId, Get5-MapNumber.

		filename := c.GetHeader("MatchZy-FileName")
		matchID := c.GetHeader("MatchZy-MatchId")
		mapNumber := c.GetHeader("MatchZy-MapNumber")

		mid, err := strconv.Atoi(matchID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		mNumber, err := strconv.Atoi(mapNumber)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := uploader.Upload(c, mid, filename, mNumber, c.Request.Body); err != nil {
			c.String(http.StatusInternalServerError, err.Error()) // カスタムエラーを返したい
		}
		c.String(http.StatusOK, "OK")
	})
}
