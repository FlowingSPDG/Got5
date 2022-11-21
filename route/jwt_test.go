package route_test

import (
	"testing"

	"github.com/FlowingSPDG/Got5/route"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	jwt, err := route.GenerateJWT("TEST_MATCHID", "SUPER_SECRET_JWT_STRING")
	assert.NoError(t, err)
	t.Logf("Generated JWT:%s", jwt)
}
