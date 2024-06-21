package ginroute

import (
	got5 "github.com/FlowingSPDG/Got5"
	"github.com/gin-gonic/gin"
)

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler(uploader got5.DemoUploader, auth got5.Auth, r gin.IRoutes) {
	r.POST("/demo", CheckDemoAuth(auth), DemoUploadHandler(uploader))
}

// SetupMatchLoadHandler Setup get5 loadmatch handler
func SetupMatchLoadHandler(loader got5.MatchLoader, auth got5.Auth, authHeaderkey string, r gin.IRoutes) {
	r.GET("/match/:matchID", CheckMatchLoaderAuth(auth, authHeaderkey), LoadMatchHandler(loader))
}

// SetupEventHandlers get5 handlers to specified fiber.Router
func SetupEventHandlers(ctrl got5.EventHandler, auth got5.Auth, authHeaderkey string, r gin.IRoutes) {
	r.POST("/event", CheckEventHandlerAuth(auth, authHeaderkey), OnEventHandler(ctrl))
}
