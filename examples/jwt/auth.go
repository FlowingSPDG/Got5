package jwt

import (
	"context"
	"fmt"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/route"
	"github.com/golang-jwt/jwt"
)

var _ controller.Auth = (*firebaseAuth)(nil)

type firebaseAuth struct {
	secret string
}

// CheckDemoAuth implements controller.Auth
func (f *firebaseAuth) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	token, err := jwt.Parse(auth, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return f.secret, nil
	})

	if claims, ok := token.Claims.(route.G5JWT); ok && token.Valid {
		fmt.Println(claims)
	} else {
		return err
	}
	return nil
}

// EventAuth implements controller.Auth
func (f *firebaseAuth) EventAuth(ctx context.Context, serverID string, auth string) error {
	token, err := jwt.Parse(auth, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return f.secret, nil
	})

	if claims, ok := token.Claims.(route.G5JWT); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
	return nil
}

// MatchAuth implements controller.Auth
func (f *firebaseAuth) MatchAuth(ctx context.Context, mid string, auth string) error {
	token, err := jwt.Parse(auth, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return f.secret, nil
	})

	if claims, ok := token.Claims.(route.G5JWT); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
	return nil
}

// NewAuth Get Logger pointer
func NewAuth(secret string) controller.Auth {
	return &firebaseAuth{
		secret: secret,
	}
}
