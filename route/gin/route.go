package ginroute

import (
	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/Got5/controller"
)

// SetupAllGet5Handlers Setup All Get5-related handlers
func SetupAllGet5Handlers(evh controller.EventHandler, loader controller.MatchLoader, uploader controller.DemoUploader, auth controller.Auth, r gin.IRoutes) {
	SetupEventHandlers(evh, auth, r)
	SetupMatchLoadHandler(loader, auth, r)
	SetupDemoUploadHandler(uploader, auth, r)
}

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler(uploader controller.DemoUploader, auth controller.Auth, r gin.IRoutes) {
	r.POST("/demo", CheckDemoAuth(auth), DemoUploadHandler(uploader))
}

// SetupMatchLoadHandler Setup get5 loadmatch handler
func SetupMatchLoadHandler(loader controller.MatchLoader, auth controller.Auth, r gin.IRoutes) {
	r.GET("/match/:matchID", CheckMatchLoaderAuth(auth), LoadMatchHandler(loader))
}

// SetupEventHandlers get5 handlers to specified fiber.Router
func SetupEventHandlers(ctrl controller.EventHandler, auth controller.Auth, r gin.IRoutes) {
	r.POST("/event", CheckEventHandlerAuth(auth), OnEventHandler(ctrl))
}
