package route_test

import (
	"testing"

	"github.com/FlowingSPDG/Got5/route"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	jwt, err := route.GenerateJWT("TEST_MATCHID", "SUPER_SECRET_JWT_STRING")
	t.Logf("jwt:%#v", jwt)
	assert.NoError(t, err)
}
