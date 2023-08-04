package webhook

import (
	"context"
	"fmt"

	"github.com/FlowingSPDG/Got5/controller"
)

type webhookAuth struct {
	password string
}

var _ controller.Auth = (*webhookAuth)(nil)

// CheckAuth implements controller.EventHandler
func (wh *webhookAuth) EventAuth(ctx context.Context, mid string, auth string) error {
	// Lets say simple password...
	if auth == wh.password {
		return nil
	}
	return fmt.Errorf("Password mismatch")
}

// CheckDemoAuth implements controller.JWTAuth
func (*webhookAuth) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID string, jwt string) error {
	panic("unimplemented")
}

// MatchAuth implements controller.JWTAuth
func (*webhookAuth) MatchAuth(ctx context.Context, mid string, jwt string) error {
	panic("unimplemented")
}

// NewAuth Get new auth pointer
func NewAuth(password string) controller.Auth {
	return &webhookAuth{
		password: password,
	}
}
