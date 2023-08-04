package jwt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FlowingSPDG/Got5/examples/jwt"
)

func TestGenerateJWT(t *testing.T) {
	jwt, err := jwt.GenerateJWT("TEST_MATCHID", "SUPER_SECRET_JWT_STRING")
	t.Logf("jwt:%#v", jwt)
	assert.NoError(t, err)
}
