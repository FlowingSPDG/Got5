package webhook

import (
	"context"
	"fmt"

	got5 "github.com/FlowingSPDG/Got5"
)

type webhookAuth struct {
	password string
}

var _ got5.Auth = (*webhookAuth)(nil)

// CheckAuth implements got5.EventHandler
func (wh *webhookAuth) EventAuth(ctx context.Context, mid string, auth string) error {
	// Lets say simple password...
	if auth == wh.password {
		return nil
	}
	return fmt.Errorf("Password mismatch")
}

// CheckDemoAuth implements got5.JWTAuth
func (*webhookAuth) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID string, jwt string) error {
	panic("unimplemented")
}

// MatchAuth implements got5.JWTAuth
func (*webhookAuth) MatchAuth(ctx context.Context, mid string, jwt string) error {
	panic("unimplemented")
}

// NewAuth Get new auth pointer
func NewAuth(password string) got5.Auth {
	return &webhookAuth{
		password: password,
	}
}
