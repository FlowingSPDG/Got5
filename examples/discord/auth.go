package discord

import (
	"context"
	"fmt"

	got5 "github.com/FlowingSPDG/Got5"
)

var _ got5.Auth = (*Discord)(nil)

// EventAuth
func (d *Discord) EventAuth(ctx context.Context, mid string, auth string) error {
	// Lets say simple password...
	if auth == d.setting.password {
		return nil
	}
	return fmt.Errorf("Password mismatch")
}

// CheckDemoAuth implements got5.JWTAuth
func (d *Discord) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID string, auth string) error {
	// Lets say simple password...
	if auth == d.setting.password {
		return nil
	}
	return fmt.Errorf("Password mismatch")
}

// MatchAuth implements got5.JWTAuth
func (d *Discord) MatchAuth(ctx context.Context, mid string, auth string) error {
	// Lets say simple password...
	if auth == d.setting.password {
		return nil
	}
	return fmt.Errorf("Password mismatch")
}
