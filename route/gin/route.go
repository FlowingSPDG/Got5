package ginroute

import (
	got5 "github.com/FlowingSPDG/Got5"
	"github.com/gin-gonic/gin"
)

// SetupAllGet5Handlers Setup All Get5-related handlers
func SetupAllGet5Handlers[TMatchID int](evh got5.EventHandler[TMatchID], loader got5.MatchLoader[TMatchID], uploader got5.DemoUploader[TMatchID], auth got5.Auth[TMatchID], r gin.IRoutes) {
	SetupEventHandlers(evh, auth, r)
	SetupMatchLoadHandler(loader, auth, r)
	SetupDemoUploadHandler(uploader, auth, r)
}

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler[TMatchID int](uploader got5.DemoUploader[TMatchID], auth got5.Auth[TMatchID], r gin.IRoutes) {
	r.POST("/demo", CheckDemoAuth(auth), DemoUploadHandler(uploader))
}

// SetupMatchLoadHandler Setup get5 loadmatch handler
func SetupMatchLoadHandler[TMatchID int](loader got5.MatchLoader[TMatchID], auth got5.Auth[TMatchID], r gin.IRoutes) {
	r.GET("/match/:matchID", CheckMatchLoaderAuth(auth), LoadMatchHandler(loader))
}

// SetupEventHandlers get5 handlers to specified fiber.Router
func SetupEventHandlers[TMatchID got5.MatchID](ctrl got5.EventHandler[TMatchID], auth got5.Auth[TMatchID], r gin.IRoutes) {
	r.POST("/event", CheckEventHandlerAuth(auth), OnEventHandler(ctrl))
}
