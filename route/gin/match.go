package ginroute

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/Got5/controller"
)

// CheckMatchLoaderAuth 認証用ハンドラ
func CheckMatchLoaderAuth(auth controller.Auth) func(c *gin.Context) {
	return (func(c *gin.Context) {
		mid := c.Param("matchID")
		reqAuthVal := c.GetHeader("Authorization")
		if err := auth.MatchAuth(c, mid, reqAuthVal); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err) // カスタムエラーを返したい
			return
		}
		c.Next()
	})
}

// LoadMatchHandler GET on get5_loadmatch_url "https://example.com/match_config.json" "Authorization" "Bearer <token>"
func LoadMatchHandler(loader controller.MatchLoader) func(c *gin.Context) {
	return (func(c *gin.Context) {
		// マッチを取得、JSONを組んで返す
		mid := c.Param("matchID")
		m, err := loader.Load(c, mid)
		if err != nil {
			// TODO:エラーハンドリングをする
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.JSON(http.StatusOK, m.ToG5Format())
	})
}
