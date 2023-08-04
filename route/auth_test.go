package route_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FlowingSPDG/Got5/route"
)

func TestGenerateJWT(t *testing.T) {
	jwt, err := route.GenerateJWT("TEST_MATCHID", "SUPER_SECRET_JWT_STRING")
	t.Logf("jwt:%#v", jwt)
	assert.NoError(t, err)
}
