package fiberroute_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
	fiberroute "github.com/FlowingSPDG/Got5/route/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var _ controller.MatchLoader = (*mockMatchLoader)(nil)
var _ controller.Auth = (*mockAuth)(nil)

type mockAuth struct {
	auth string
}

// CheckDemoAuth implements controller.Auth
func (a *mockAuth) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	if auth == a.auth {
		return nil
	}
	return fmt.Errorf("Fail")
}

// EventAuth implements controller.Auth
func (a *mockAuth) EventAuth(ctx context.Context, serverID string, auth string) error {
	if auth == a.auth {
		return nil
	}
	return fmt.Errorf("Fail")
}

// MatchAuth implements controller.Auth
func (a *mockAuth) MatchAuth(ctx context.Context, mid string, auth string) error {
	if auth == a.auth {
		return nil
	}
	return fmt.Errorf("Fail")
}

type mockMatchLoader struct {
	match map[string]models.Match
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
		auth   *mockAuth
		mid    string
	}{
		{
			title:  "Fail auth",
			loader: &mockMatchLoader{match: map[string]models.Match{"": {MatchTitle: "", MatchID: "", ClinchSeries: false, NumMaps: 0, PlayersPerTeam: 0, CoachesPerTeam: 0, CoachesMustReady: false, MinPlayersToReady: 0, MinSpectatorsToReady: 0, SkipVeto: false, VetoFirst: "", SideType: "", Spectators: models.Spectators{}, Maplist: []string{}, MapSides: []string{}, Team1: models.Team{}, Team2: models.Team{}, Cvars: map[string]string{}}}},
			mid:    "TEST_MATCH",
		},
	}

	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			// Setup fiber
			app := fiber.New()
			g5test := app.Group("/get5testloadmatch") // /test
			fiberroute.SetupMatchLoadHandler(tt.loader, tt.auth, g5test)

			r := httptest.NewRequest("GET", "/get5testloadmatch/event", nil)
			r.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(r, -1)
			defer resp.Body.Close()
			asserts.NoError(err)
		})
	}
}
