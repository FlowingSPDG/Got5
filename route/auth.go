package route

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// G5JWT Got5 JWT. CVAR only supports 128 characters so shrink as much as you can
type G5JWT struct {
	MatchID string `json:"m"`
}

// Valid Check if JWT valids. Impelement for jwt.Claims
func (j G5JWT) Valid() error {
	if j.MatchID == "" {
		return fmt.Errorf("Empty")
	}
	return nil
}

// GenerateJWT Generate JWT from specified matchID
func GenerateJWT(matchID string, secret string) (string, error) {
	// Create the Claims
	g := G5JWT{
		MatchID: matchID,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, g)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
