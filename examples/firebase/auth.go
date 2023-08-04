package fb

import (
	"context"

	"firebase.google.com/go/auth"

	"github.com/FlowingSPDG/Got5/controller"
)

var _ controller.Auth = (*firebaseAuth)(nil)

type firebaseAuth struct {
	a *auth.Client
}

// EventAuth
func (f *firebaseAuth) EventAuth(ctx context.Context, mid string, auth string) error {
	_, err := f.a.VerifyIDToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}

// CheckDemoAuth implements controller.JWTAuth
func (f *firebaseAuth) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID string, auth string) error {
	_, err := f.a.VerifyIDToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}

// MatchAuth implements controller.JWTAuth
func (f *firebaseAuth) MatchAuth(ctx context.Context, mid string, auth string) error {
	_, err := f.a.VerifyIDToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}

// NewAuth Get controller.Auth based firebase auth
func NewAuth(a *auth.Client) controller.Auth {
	return &firebaseAuth{
		a: a,
	}
}
