package discord

import (
	"context"
	"fmt"

	"github.com/FlowingSPDG/Got5/controller"
)

var _ controller.Auth = (*Discord)(nil)

// EventAuth
func (d *Discord) EventAuth(ctx context.Context, mid string, auth string) error {
	// Lets say simple password...
	if auth == d.setting.password {
		return nil
	}
	return fmt.Errorf("Password mismatch")
}

// CheckDemoAuth implements controller.JWTAuth
func (d *Discord) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	// Lets say simple password...
	if auth == d.setting.password {
		return nil
	}
	return fmt.Errorf("Password mismatch")
}

// MatchAuth implements controller.JWTAuth
func (d *Discord) MatchAuth(ctx context.Context, mid string, auth string) error {
	// Lets say simple password...
	if auth == d.setting.password {
		return nil
	}
	return fmt.Errorf("Password mismatch")
}
