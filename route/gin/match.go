package ginroute

import (
	"net/http"
	"strconv"

	got5 "github.com/FlowingSPDG/Got5"
	"github.com/gin-gonic/gin"
)

// CheckMatchLoaderAuth 認証用ハンドラ
func CheckMatchLoaderAuth(auth got5.Auth, headerKey string) func(c *gin.Context) {
	return (func(c *gin.Context) {
		matchID := c.Param("matchID")
		headerValue := c.GetHeader(headerKey)

		mid, err := strconv.Atoi(matchID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := auth.MatchAuth(c, mid, headerValue); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err) // カスタムエラーを返したい
			return
		}
		c.Next()
	})
}

// LoadMatchHandler GET on get5_loadmatch_url "https://example.com/match_config.json" "Authorization" "Bearer <token>"
func LoadMatchHandler(loader got5.MatchLoader) func(c *gin.Context) {
	return (func(c *gin.Context) {
		// マッチを取得、JSONを組んで返す
		matchID := c.Param("matchID")
		mid, err := strconv.Atoi(matchID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		m, err := loader.Load(c, mid)
		if err != nil {
			// TODO:エラーハンドリングをする
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.JSON(http.StatusOK, m)
	})
}
