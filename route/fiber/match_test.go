package fiberroute_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	got5 "github.com/FlowingSPDG/Got5"
	fiberroute "github.com/FlowingSPDG/Got5/route/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var _ got5.MatchLoader = (*mockMatchLoader)(nil)
var _ got5.Auth = (*mockAuth)(nil)

type mockAuth struct {
	auth string
}

// CheckDemoAuth implements got5.Auth
func (a *mockAuth) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID string, auth string) error {
	if auth == a.auth {
		return nil
	}
	return fmt.Errorf("Fail")
}

// EventAuth implements got5.Auth
func (a *mockAuth) EventAuth(ctx context.Context, serverID string, auth string) error {
	if auth == a.auth {
		return nil
	}
	return fmt.Errorf("Fail")
}

// MatchAuth implements got5.Auth
func (a *mockAuth) MatchAuth(ctx context.Context, mid string, auth string) error {
	if auth == a.auth {
		return nil
	}
	return fmt.Errorf("Fail")
}

type mockMatchLoader struct {
	match map[string]got5.Match
}

// Load implements got5.MatchLoader
func (m *mockMatchLoader) Load(ctx context.Context, mid string) (got5.G5Match, error) {
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
			loader: &mockMatchLoader{match: map[string]got5.Match{"": {MatchTitle: "", MatchID: "", ClinchSeries: false, NumMaps: 0, PlayersPerTeam: 0, CoachesPerTeam: 0, CoachesMustReady: false, MinPlayersToReady: 0, MinSpectatorsToReady: 0, SkipVeto: false, VetoFirst: "", SideType: "", Spectators: got5.Spectators{}, Maplist: []string{}, MapSides: []string{}, Team1: got5.Team{}, Team2: got5.Team{}, Cvars: map[string]string{}}}},
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
