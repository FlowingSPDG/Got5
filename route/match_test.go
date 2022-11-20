package route_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
	"github.com/FlowingSPDG/Got5/route"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var _ controller.MatchLoader = (*mockMatchLoader)(nil)

type mockMatchLoader struct {
	auths map[string]string
	match map[string]models.Match
}

// CheckMatchAuth implements controller.MatchLoader
func (m *mockMatchLoader) CheckMatchAuth(ctx context.Context, mid string, auth string) error {
	if a, ok := m.auths[mid]; ok {
		if a == auth {
			return nil
		}
	}
	return fmt.Errorf("Fail")
}

// Load implements controller.MatchLoader
func (m *mockMatchLoader) Load(ctx context.Context, mid string) (models.G5Match, error) {
	if m, ok := m.match[mid]; ok {
		return m, nil
	}
	return nil, fmt.Errorf("Fail")
}

// TestLoadMatchSuccess 成功のみのパターン
func TestLoadMatchSuccess(t *testing.T) {
	asserts := assert.New(t)

	cases := []struct {
		title  string
		loader *mockMatchLoader
		mid    string
		auth   string
	}{
		{
			title:  "Fail auth",
			loader: &mockMatchLoader{auths: map[string]string{"TEST_MATCH": "ABC"}, match: map[string]models.Match{"": {MatchTitle: "", MatchID: "", ClinchSeries: false, NumMaps: 0, PlayersPerTeam: 0, CoachesPerTeam: 0, CoachesMustReady: false, MinPlayersToReady: 0, MinSpectatorsToReady: 0, SkipVeto: false, VetoFirst: "", SideType: "", Spectators: models.Spectators{}, Maplist: []string{}, MapSides: []string{}, Team1: models.Team{}, Team2: models.Team{}, Cvars: map[string]string{}}}},
			mid:    "TEST_MATCH",
			auth:   "ABC",
		},
	}

	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			// Setup fiber
			app := fiber.New()
			g5test := app.Group("/get5testloadmatch") // /test
			err := route.SetupMatchLoadHandler(tt.loader, g5test)
			asserts.NoError(err)

			r := httptest.NewRequest("GET", "/get5testloadmatch/event", nil)
			r.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(r, -1)
			defer resp.Body.Close()
			asserts.NoError(err)
		})
	}
}
