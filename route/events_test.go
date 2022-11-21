package route_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
	"github.com/FlowingSPDG/Got5/route"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var _ controller.EventHandler = (*mockEventHandler)(nil)

type mockEventHandler struct {
	expect any
	parsed any
}

// Close implements controller.EventHandler
func (m *mockEventHandler) Close() error {
	return nil
}

// HandleOnBackupRestore implements controller.EventHandler
func (m *mockEventHandler) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	m.parsed = p
	return nil
}

// HandleOnBombDefused implements controller.EventHandler
func (m *mockEventHandler) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnBombExploded implements controller.EventHandler
func (m *mockEventHandler) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnBombPlanted implements controller.EventHandler
func (m *mockEventHandler) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnDecoyStarted implements controller.EventHandler
func (m *mockEventHandler) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnDemoFinished implements controller.EventHandler
func (m *mockEventHandler) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnDemoUploadEnded implements controller.EventHandler
func (m *mockEventHandler) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnFlashbangDetonated implements controller.EventHandler
func (m *mockEventHandler) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnGameStateChanged implements controller.EventHandler
func (m *mockEventHandler) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnGoingLive implements controller.EventHandler
func (m *mockEventHandler) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	m.parsed = p
	return nil
}

// HandleOnGrenadeThrown implements controller.EventHandler
func (m *mockEventHandler) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	m.parsed = p
	return nil
}

// HandleOnHEGrenadeDetonated implements controller.EventHandler
func (m *mockEventHandler) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnKnifeRoundStarted implements controller.EventHandler
func (m *mockEventHandler) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnKnifeRoundWon implements controller.EventHandler
func (m *mockEventHandler) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	m.parsed = p
	return nil
}

// HandleOnLoadMatchConfigFailed implements controller.EventHandler
func (m *mockEventHandler) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMapPicked implements controller.EventHandler
func (m *mockEventHandler) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMapResult implements controller.EventHandler
func (m *mockEventHandler) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMapVetoed implements controller.EventHandler
func (m *mockEventHandler) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMatchPaused implements controller.EventHandler
func (m *mockEventHandler) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMatchUnpaused implements controller.EventHandler
func (m *mockEventHandler) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMolotovDetonated implements controller.EventHandler
func (m *mockEventHandler) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerBecameMVP implements controller.EventHandler
func (m *mockEventHandler) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerConnected implements controller.EventHandler
func (m *mockEventHandler) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerDeath implements controller.EventHandler
func (m *mockEventHandler) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerDisconnected implements controller.EventHandler
func (m *mockEventHandler) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerSay implements controller.EventHandler
func (m *mockEventHandler) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPreLoadMatchConfig implements controller.EventHandler
func (m *mockEventHandler) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	m.parsed = p
	return nil
}

// HandleOnRoundEnd implements controller.EventHandler
func (m *mockEventHandler) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	m.parsed = p
	return nil
}

// HandleOnRoundStart implements controller.EventHandler
func (m *mockEventHandler) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	m.parsed = p
	return nil
}

// HandleOnRoundStatsUpdated implements controller.EventHandler
func (m *mockEventHandler) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnSeriesInit implements controller.EventHandler
func (m *mockEventHandler) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	m.parsed = p
	return nil
}

// HandleOnSeriesResult implements controller.EventHandler
func (m *mockEventHandler) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	m.parsed = p
	return nil
}

// HandleOnSidePicked implements controller.EventHandler
func (m *mockEventHandler) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnSmokeGrenadeDetonated implements controller.EventHandler
func (m *mockEventHandler) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnTeamReadyStatusChanged implements controller.EventHandler
func (m *mockEventHandler) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	m.parsed = p
	return nil
}

// Hostname implements controller.EventHandler
func (m *mockEventHandler) Hostname() string {
	return ""
}

func TestEventHandleTD(t *testing.T) {
	asserts := assert.New(t)

	cases := []struct {
		title        string
		eventHandler *mockEventHandler
		auth         *mockAuth
		statusCode   int
		input        []byte
		// err        error
	}{
		{
			title: "Example Invalid Event",
			eventHandler: &mockEventHandler{
				expect: nil,
			},
			statusCode: http.StatusBadRequest,
			input: []byte(`{
				"event": "string"
			  }`),
			// err:        nil,
		},
		{
			title: "GameStateChanged",
			eventHandler: &mockEventHandler{
				expect: models.OnGameStateChangedPayload{
					Event:    models.Event{Event: "game_state_changed"},
					NewState: "none",
					OldState: "none",
				},
			},
			statusCode: http.StatusOK,
			// err:        nil,
			input: []byte(`{"event": "game_state_changed","new_state": "none","old_state": "none"}`),
		},
		{
			title: "PreloadMatchConfig",
			eventHandler: &mockEventHandler{
				expect: models.OnPreLoadMatchConfigPayload{
					Event:    models.Event{Event: "preload_match_config"},
					Filename: "TEST_FILENAME",
				},
			},
			statusCode: http.StatusOK,
			// err:        nil,
			input: []byte(`{"event": "preload_match_config","filename": "TEST_FILENAME"}`),
		},
		{
			title: "MatchConfigLoadFail",
			eventHandler: &mockEventHandler{
				expect: models.OnLoadMatchConfigFailedPayload{
					Event:  models.Event{Event: "match_config_load_fail"},
					Reason: "You done goofed.",
				},
			},
			statusCode: http.StatusOK,
			// err:        nil,
			input: []byte(`{"event": "match_config_load_fail","reason": "You done goofed."}`),
		},
		{
			title: "SeriesInit",
			eventHandler: &mockEventHandler{
				expect: models.OnSeriesInitPayload{
					Event:     models.Event{Event: "series_start"},
					Matchid:   "14272",
					Team1Name: "NaVi",
					Team2Name: "Astralis",
				},
			},
			statusCode: http.StatusOK,
			// err:        nil,
			input: []byte(`{"event": "series_start","matchid": "14272","team1_name": "NaVi","team2_name": "Astralis"}`),
		},
		{
			title: "MapResult",
			eventHandler: &mockEventHandler{
				expect: models.OnMapResultPayload{
					Event:      models.Event{Event: "map_result"},
					Matchid:    "14272",
					MapNumber:  0,
					Team1Score: 16,
					Team2Score: 13,
					Winner:     models.Winner{Side: "ct", Team: "team1"},
				},
			},
			statusCode: http.StatusOK,
			// err:        nil,
			input: []byte(`{"event": "map_result","matchid": "14272","map_number": 0,"team1_score": 16,"team2_score": 13,"winner": {"side": "ct","team": "team1"}}`),
		},
		{
			title: "SeriesEnd",
			eventHandler: &mockEventHandler{
				expect: models.OnSeriesResultPayload{
					Event:            models.Event{Event: "series_end"},
					Matchid:          "14272",
					Team1SeriesScore: 2,
					Team2SeriesScore: 0,
					Winner:           models.Winner{Side: "ct", Team: "team1"},
					TimeUntilRestore: 45,
				},
			},
			statusCode: http.StatusOK,
			// err:        nil,
			input: []byte(`{"event": "series_end","matchid": "14272","team1_series_score": 2,"team2_series_score": 0,"winner": {"side": "ct","team": "team1"},"time_until_restore": 45}`),
		},
	}

	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			// Setup fiber
			app := fiber.New()
			g5test := app.Group("/get5testevent") // /test
			err := route.SetupEventHandlers(tt.eventHandler, tt.auth, g5test)
			asserts.NoError(err)

			r := httptest.NewRequest("POST", "/get5testevent/event", bytes.NewBuffer(tt.input))
			r.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(r, -1)
			defer resp.Body.Close()
			b, _ := io.ReadAll(resp.Body)
			t.Logf("b:%s\n", b)
			// asserts.Equal(tt.err, err)
			asserts.Equal(tt.statusCode, resp.StatusCode)
			asserts.Equal(tt.eventHandler.parsed, tt.eventHandler.parsed)
		})
	}
}
