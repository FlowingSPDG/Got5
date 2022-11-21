package logger

import (
	"context"
	"log"

	"github.com/FlowingSPDG/Got5/controller"
)

var _ controller.Auth = (*loggerAuth)(nil)

type loggerAuth struct{}

// CheckDemoAuth implements controller.Auth
func (*loggerAuth) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	log.Println("mid", mid, "filename", filename, "mapNumber", mapNumber, "serverID", serverID, "auth", auth)
	return nil
}

// EventAuth implements controller.Auth
func (*loggerAuth) EventAuth(ctx context.Context, serverID string, auth string) error {
	log.Println("serverID", serverID, "auth", auth)
	return nil
}

// MatchAuth implements controller.Auth
func (*loggerAuth) MatchAuth(ctx context.Context, mid string, auth string) error {
	log.Println("mid", mid, "auth", auth)
	return nil
}

// NewAuth Get Logger pointer
func NewAuth() controller.Auth {
	return &loggerAuth{}
}
