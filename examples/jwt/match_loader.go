package jwt

import (
	"context"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
	"github.com/FlowingSPDG/Got5/route"
)

var _ controller.MatchLoader = (*jwtMatchLoader)(nil)

type jwtMatchLoader struct {
	secret string
}

// GetMatch implements controller.MatchLoader
func (j *jwtMatchLoader) Load(ctx context.Context, mid string) (models.G5Match, error) {
	m := models.GetDefaultMatchBO1()
	jwt, err := route.GenerateJWT(mid, j.secret)
	if err != nil {
		return nil, err
	}
	m.Cvars["get5_remote_log_header_value"] = jwt
	return m, nil
}

// NewMatchLoader Get Logger pointer
func NewMatchLoader(secret string) controller.MatchLoader {
	return &jwtMatchLoader{
		secret: secret,
	}
}
