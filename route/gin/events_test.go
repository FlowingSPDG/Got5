package ginroute_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	got5 "github.com/FlowingSPDG/Got5"
	ginroute "github.com/FlowingSPDG/Got5/route/gin"
)

func getStringPointer(s string) *string {
	return &s
}

var _ got5.EventHandler = (*mockEventHandler)(nil)

type mockEventHandler struct {
	expect any
	parsed any
}

// HandleOnPauseBegan implements got5.EventHandler.
func (m *mockEventHandler) HandleOnPauseBegan(ctx context.Context, p got5.OnPauseBeganPayload) error {
	m.parsed = p
	return nil
}

// Close implements got5.EventHandler
func (m *mockEventHandler) Close() error {
	return nil
}

// HandleOnBackupRestore implements got5.EventHandler
func (m *mockEventHandler) HandleOnBackupRestore(ctx context.Context, p got5.OnBackupRestorePayload) error {
	m.parsed = p
	return nil
}

// HandleOnBombDefused implements got5.EventHandler
func (m *mockEventHandler) HandleOnBombDefused(ctx context.Context, p got5.OnBombDefusedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnBombExploded implements got5.EventHandler
func (m *mockEventHandler) HandleOnBombExploded(ctx context.Context, p got5.OnBombExplodedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnBombPlanted implements got5.EventHandler
func (m *mockEventHandler) HandleOnBombPlanted(ctx context.Context, p got5.OnBombPlantedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnDecoyStarted implements got5.EventHandler
func (m *mockEventHandler) HandleOnDecoyStarted(ctx context.Context, p got5.OnDecoyStartedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnDemoFinished implements got5.EventHandler
func (m *mockEventHandler) HandleOnDemoFinished(ctx context.Context, p got5.OnDemoFinishedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnDemoUploadEnded implements got5.EventHandler
func (m *mockEventHandler) HandleOnDemoUploadEnded(ctx context.Context, p got5.OnDemoUploadEndedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnFlashbangDetonated implements got5.EventHandler
func (m *mockEventHandler) HandleOnFlashbangDetonated(ctx context.Context, p got5.OnFlashbangDetonatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnGameStateChanged implements got5.EventHandler
func (m *mockEventHandler) HandleOnGameStateChanged(ctx context.Context, p got5.OnGameStateChangedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnGoingLive implements got5.EventHandler
func (m *mockEventHandler) HandleOnGoingLive(ctx context.Context, p got5.OnGoingLivePayload) error {
	m.parsed = p
	return nil
}

// HandleOnGrenadeThrown implements got5.EventHandler
func (m *mockEventHandler) HandleOnGrenadeThrown(ctx context.Context, p got5.OnGrenadeThrownPayload) error {
	m.parsed = p
	return nil
}

// HandleOnHEGrenadeDetonated implements got5.EventHandler
func (m *mockEventHandler) HandleOnHEGrenadeDetonated(ctx context.Context, p got5.OnHEGrenadeDetonatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnKnifeRoundStarted implements got5.EventHandler
func (m *mockEventHandler) HandleOnKnifeRoundStarted(ctx context.Context, p got5.OnKnifeRoundStartedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnKnifeRoundWon implements got5.EventHandler
func (m *mockEventHandler) HandleOnKnifeRoundWon(ctx context.Context, p got5.OnKnifeRoundWonPayload) error {
	m.parsed = p
	return nil
}

// HandleOnLoadMatchConfigFailed implements got5.EventHandler
func (m *mockEventHandler) HandleOnLoadMatchConfigFailed(ctx context.Context, p got5.OnLoadMatchConfigFailedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMapPicked implements got5.EventHandler
func (m *mockEventHandler) HandleOnMapPicked(ctx context.Context, p got5.OnMapPickedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMapResult implements got5.EventHandler
func (m *mockEventHandler) HandleOnMapResult(ctx context.Context, p got5.OnMapResultPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMapVetoed implements got5.EventHandler
func (m *mockEventHandler) HandleOnMapVetoed(ctx context.Context, p got5.OnMapVetoedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMatchPaused implements got5.EventHandler
func (m *mockEventHandler) HandleOnMatchPaused(ctx context.Context, p got5.OnMatchPausedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMatchUnpaused implements got5.EventHandler
func (m *mockEventHandler) HandleOnMatchUnpaused(ctx context.Context, p got5.OnMatchUnpausedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnMolotovDetonated implements got5.EventHandler
func (m *mockEventHandler) HandleOnMolotovDetonated(ctx context.Context, p got5.OnMolotovDetonatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerBecameMVP implements got5.EventHandler
func (m *mockEventHandler) HandleOnPlayerBecameMVP(ctx context.Context, p got5.OnPlayerBecameMVPPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerConnected implements got5.EventHandler
func (m *mockEventHandler) HandleOnPlayerConnected(ctx context.Context, p got5.OnPlayerConnectedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerDeath implements got5.EventHandler
func (m *mockEventHandler) HandleOnPlayerDeath(ctx context.Context, p got5.OnPlayerDeathPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerDisconnected implements got5.EventHandler
func (m *mockEventHandler) HandleOnPlayerDisconnected(ctx context.Context, p got5.OnPlayerDisconnectedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPlayerSay implements got5.EventHandler
func (m *mockEventHandler) HandleOnPlayerSay(ctx context.Context, p got5.OnPlayerSayPayload) error {
	m.parsed = p
	return nil
}

// HandleOnPreLoadMatchConfig implements got5.EventHandler
func (m *mockEventHandler) HandleOnPreLoadMatchConfig(ctx context.Context, p got5.OnPreLoadMatchConfigPayload) error {
	m.parsed = p
	return nil
}

// HandleOnRoundEnd implements got5.EventHandler
func (m *mockEventHandler) HandleOnRoundEnd(ctx context.Context, p got5.OnRoundEndPayload) error {
	m.parsed = p
	return nil
}

// HandleOnRoundStart implements got5.EventHandler
func (m *mockEventHandler) HandleOnRoundStart(ctx context.Context, p got5.OnRoundStartPayload) error {
	m.parsed = p
	return nil
}

// HandleOnRoundStatsUpdated implements got5.EventHandler
func (m *mockEventHandler) HandleOnRoundStatsUpdated(ctx context.Context, p got5.OnRoundStatsUpdatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnSeriesInit implements got5.EventHandler
func (m *mockEventHandler) HandleOnSeriesInit(ctx context.Context, p got5.OnSeriesInitPayload) error {
	m.parsed = p
	return nil
}

// HandleOnSeriesResult implements got5.EventHandler
func (m *mockEventHandler) HandleOnSeriesResult(ctx context.Context, p got5.OnSeriesResultPayload) error {
	m.parsed = p
	return nil
}

// HandleOnSidePicked implements got5.EventHandler
func (m *mockEventHandler) HandleOnSidePicked(ctx context.Context, p got5.OnSidePickedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnSmokeGrenadeDetonated implements got5.EventHandler
func (m *mockEventHandler) HandleOnSmokeGrenadeDetonated(ctx context.Context, p got5.OnSmokeGrenadeDetonatedPayload) error {
	m.parsed = p
	return nil
}

// HandleOnTeamReadyStatusChanged implements got5.EventHandler
func (m *mockEventHandler) HandleOnTeamReadyStatusChanged(ctx context.Context, p got5.OnTeamReadyStatusChangedPayload) error {
	m.parsed = p
	return nil
}

// Hostname implements got5.EventHandler
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
			title:        "Example Invalid Event",
			eventHandler: &mockEventHandler{expect: nil, parsed: nil},
			auth:         &mockAuth{},
			statusCode:   http.StatusBadRequest,
			input:        []byte(`{"event": "string"}`),
		},
		{
			title: "GameStateChanged",
			eventHandler: &mockEventHandler{
				expect: got5.OnGameStateChangedPayload{Event: got5.Event{Event: "game_state_changed"}, NewState: "none", OldState: "none"},
			},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "game_state_changed",
				"new_state": "none",
				"old_state": "none"
			  }`),
		},
		{
			title: "PreloadMatchConfig",
			eventHandler: &mockEventHandler{
				expect: got5.OnPreLoadMatchConfigPayload{Event: got5.Event{Event: "preload_match_config"}, Filename: "string"},
			},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "preload_match_config",
				"filename": "string"
			  }`),
		},
		{
			title:        "MatchConfigLoadFail",
			eventHandler: &mockEventHandler{expect: got5.OnLoadMatchConfigFailedPayload{Event: got5.Event{Event: "match_config_load_fail"}, Reason: "You done goofed."}},
			auth:         &mockAuth{},
			statusCode:   http.StatusOK,
			input: []byte(`{
				"event": "match_config_load_fail",
				"reason": "You done goofed."
			  }`),
		},
		{
			title: "SeriesInit",
			eventHandler: &mockEventHandler{expect: got5.OnSeriesInitPayload{
				Event:   got5.Event{Event: "series_start"},
				MatchID: "14272",
				NumMaps: 1,
				Team1: struct {
					ID   string "json:\"id\""
					Name string "json:\"name\""
				}{
					ID:   "2843",
					Name: "Natus Vincere",
				},
				Team2: struct {
					ID   string "json:\"id\""
					Name string "json:\"name\""
				}{
					ID:   "2843",
					Name: "Natus Vincere",
				},
			},
			},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "series_start",
				"matchid": "14272",
				"num_maps": 1,
				"team1": {
				  "id": "2843",
				  "name": "Natus Vincere"
				},
				"team2": {
				  "id": "2843",
				  "name": "Natus Vincere"
				}
			  }`),
		},
		{
			title: "MapResult",
			eventHandler: &mockEventHandler{
				expect: got5.OnMapResultPayload{
					Event:     got5.Event{Event: "map_result"},
					MatchID:   "14272",
					MapNumber: 0,
					Team1: got5.Get5StatsTeam{
						ID:          "2843",
						Name:        "Natus Vincere",
						SeriesScore: 0,
						Score:       14,
						ScoreCt:     10,
						ScoreT:      14,
						Players: []got5.Get5StatsPlayer{
							{
								SteamID: "76561198279375306",
								Name:    "s1mple",
								Stats: struct {
									Kills             int "json:\"kills\""
									Deaths            int "json:\"deaths\""
									Assists           int "json:\"assists\""
									FlashAssists      int "json:\"flash_assists\""
									TeamKills         int "json:\"team_kills\""
									Suicides          int "json:\"suicides\""
									Damage            int "json:\"damage\""
									UtilityDamage     int "json:\"utility_damage\""
									EnemiesFlashed    int "json:\"enemies_flashed\""
									FriendliesFlashed int "json:\"friendlies_flashed\""
									KnifeKills        int "json:\"knife_kills\""
									HeadshotKills     int "json:\"headshot_kills\""
									RoundsPlayed      int "json:\"rounds_played\""
									BombDefuses       int "json:\"bomb_defuses\""
									BombPlants        int "json:\"bomb_plants\""
									OneK              int "json:\"1k\""
									TwoK              int "json:\"2k\""
									ThreeK            int "json:\"3k\""
									FourK             int "json:\"4k\""
									FiveK             int "json:\"5k\""
									OneV1             int "json:\"1v1\""
									OneV2             int "json:\"1v2\""
									OneV3             int "json:\"1v3\""
									OneV4             int "json:\"1v4\""
									OneV5             int "json:\"1v5\""
									FirstKillsT       int "json:\"first_kills_t\""
									FirstKillsCt      int "json:\"first_kills_ct\""
									FirstDeathsT      int "json:\"first_deaths_t\""
									FirstDeathsCt     int "json:\"first_deaths_ct\""
									TradeKills        int "json:\"trade_kills\""
									KAST              int "json:\"kast\""
									Score             int "json:\"score\""
									Mvp               int "json:\"mvp\""
								}{
									Kills:             34,
									Deaths:            8,
									Assists:           5,
									FlashAssists:      3,
									TeamKills:         0,
									Suicides:          0,
									Damage:            2948,
									UtilityDamage:     173,
									EnemiesFlashed:    6,
									FriendliesFlashed: 2,
									KnifeKills:        1,
									HeadshotKills:     19,
									RoundsPlayed:      27,
									BombDefuses:       4,
									BombPlants:        3,
									OneK:              3,
									TwoK:              2,
									ThreeK:            3,
									FourK:             0,
									FiveK:             1,
									OneV1:             1,
									OneV2:             3,
									OneV3:             2,
									OneV4:             0,
									OneV5:             1,
									FirstKillsT:       6,
									FirstKillsCt:      5,
									FirstDeathsT:      1,
									FirstDeathsCt:     1,
									TradeKills:        3,
									KAST:              23,
									Score:             45,
									Mvp:               4,
								},
							},
						},
						Side:         "ct",
						StartingSide: "ct",
					},
					Team2: got5.Get5StatsTeam{
						ID:          "2843",
						Name:        "Natus Vincere",
						SeriesScore: 0,
						Score:       14,
						ScoreCt:     10,
						ScoreT:      14,
						Players: []got5.Get5StatsPlayer{
							{
								SteamID: "76561198279375306",
								Name:    "s1mple",
								Stats: struct {
									Kills             int "json:\"kills\""
									Deaths            int "json:\"deaths\""
									Assists           int "json:\"assists\""
									FlashAssists      int "json:\"flash_assists\""
									TeamKills         int "json:\"team_kills\""
									Suicides          int "json:\"suicides\""
									Damage            int "json:\"damage\""
									UtilityDamage     int "json:\"utility_damage\""
									EnemiesFlashed    int "json:\"enemies_flashed\""
									FriendliesFlashed int "json:\"friendlies_flashed\""
									KnifeKills        int "json:\"knife_kills\""
									HeadshotKills     int "json:\"headshot_kills\""
									RoundsPlayed      int "json:\"rounds_played\""
									BombDefuses       int "json:\"bomb_defuses\""
									BombPlants        int "json:\"bomb_plants\""
									OneK              int "json:\"1k\""
									TwoK              int "json:\"2k\""
									ThreeK            int "json:\"3k\""
									FourK             int "json:\"4k\""
									FiveK             int "json:\"5k\""
									OneV1             int "json:\"1v1\""
									OneV2             int "json:\"1v2\""
									OneV3             int "json:\"1v3\""
									OneV4             int "json:\"1v4\""
									OneV5             int "json:\"1v5\""
									FirstKillsT       int "json:\"first_kills_t\""
									FirstKillsCt      int "json:\"first_kills_ct\""
									FirstDeathsT      int "json:\"first_deaths_t\""
									FirstDeathsCt     int "json:\"first_deaths_ct\""
									TradeKills        int "json:\"trade_kills\""
									KAST              int "json:\"kast\""
									Score             int "json:\"score\""
									Mvp               int "json:\"mvp\""
								}{
									Kills:             34,
									Deaths:            8,
									Assists:           5,
									FlashAssists:      3,
									TeamKills:         0,
									Suicides:          0,
									Damage:            2948,
									UtilityDamage:     173,
									EnemiesFlashed:    6,
									FriendliesFlashed: 2,
									KnifeKills:        1,
									HeadshotKills:     19,
									RoundsPlayed:      27,
									BombDefuses:       4,
									BombPlants:        3,
									OneK:              3,
									TwoK:              2,
									ThreeK:            3,
									FourK:             0,
									FiveK:             1,
									OneV1:             1,
									OneV2:             3,
									OneV3:             2,
									OneV4:             0,
									OneV5:             1,
									FirstKillsT:       6,
									FirstKillsCt:      5,
									FirstDeathsT:      1,
									FirstDeathsCt:     1,
									TradeKills:        3,
									KAST:              23,
									Score:             45,
									Mvp:               4,
								},
							},
						},
						Side:         "ct",
						StartingSide: "ct",
					},
					Winner: got5.Winner{Side: "ct", Team: "team1"}},
			},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "map_result",
				"matchid": "14272",
				"map_number": 0,
				"team1": {
				  "id": "2843",
				  "name": "Natus Vincere",
				  "series_score": 0,
				  "score": 14,
				  "score_ct": 10,
				  "score_t": 14,
				  "players": [
					{
					  "steamid": "76561198279375306",
					  "name": "s1mple",
					  "stats": {
						"kills": 34,
						"deaths": 8,
						"assists": 5,
						"flash_assists": 3,
						"team_kills": 0,
						"suicides": 0,
						"damage": 2948,
						"utility_damage": 173,
						"enemies_flashed": 6,
						"friendlies_flashed": 2,
						"knife_kills": 1,
						"headshot_kills": 19,
						"rounds_played": 27,
						"bomb_defuses": 4,
						"bomb_plants": 3,
						"1k": 3,
						"2k": 2,
						"3k": 3,
						"4k": 0,
						"5k": 1,
						"1v1": 1,
						"1v2": 3,
						"1v3": 2,
						"1v4": 0,
						"1v5": 1,
						"first_kills_t": 6,
						"first_kills_ct": 5,
						"first_deaths_t": 1,
						"first_deaths_ct": 1,
						"trade_kills": 3,
						"kast": 23,
						"score": 45,
						"mvp": 4
					  }
					}
				  ],
				  "side": "ct",
				  "starting_side": "ct"
				},
				"team2": {
				  "id": "2843",
				  "name": "Natus Vincere",
				  "series_score": 0,
				  "score": 14,
				  "score_ct": 10,
				  "score_t": 14,
				  "players": [
					{
					  "steamid": "76561198279375306",
					  "name": "s1mple",
					  "stats": {
						"kills": 34,
						"deaths": 8,
						"assists": 5,
						"flash_assists": 3,
						"team_kills": 0,
						"suicides": 0,
						"damage": 2948,
						"utility_damage": 173,
						"enemies_flashed": 6,
						"friendlies_flashed": 2,
						"knife_kills": 1,
						"headshot_kills": 19,
						"rounds_played": 27,
						"bomb_defuses": 4,
						"bomb_plants": 3,
						"1k": 3,
						"2k": 2,
						"3k": 3,
						"4k": 0,
						"5k": 1,
						"1v1": 1,
						"1v2": 3,
						"1v3": 2,
						"1v4": 0,
						"1v5": 1,
						"first_kills_t": 6,
						"first_kills_ct": 5,
						"first_deaths_t": 1,
						"first_deaths_ct": 1,
						"trade_kills": 3,
						"kast": 23,
						"score": 45,
						"mvp": 4
					  }
					}
				  ],
				  "side": "ct",
				  "starting_side": "ct"
				},
				"winner": {
				  "side": "ct",
				  "team": "team1"
				}
			  }`),
		},
		{
			title: "SeriesEnd",
			eventHandler: &mockEventHandler{expect: got5.OnSeriesResultPayload{
				Event:            got5.Event{Event: "series_end"},
				MatchID:          "14272",
				Team1SeriesScore: 2,
				Team2SeriesScore: 0,
				Winner:           got5.Winner{Side: "ct", Team: "team1"},
				TimeUntilRestore: 45,
			},
			},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "series_end",
				"matchid": "14272",
				"team1_series_score": 2,
				"team2_series_score": 0,
				"winner": {
				  "side": "ct",
				  "team": "team1"
				},
				"time_until_restore": 45
			  }`),
		},
		{
			title: "SidePicked",
			eventHandler: &mockEventHandler{expect: got5.OnSidePickedPayload{
				Event:     got5.Event{Event: "side_picked"},
				MatchID:   "14272",
				Team:      "team1",
				MapName:   "de_nuke",
				Side:      "ct",
				MapNumber: 0,
			},
			},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "side_picked",
				"matchid": "14272",
				"team": "team1",
				"map_name": "de_nuke",
				"side": "ct",
				"map_number": 0
			  }`),
		},
		{
			title: "MapPicked",
			eventHandler: &mockEventHandler{expect: got5.OnMapPickedPayload{
				Event:     got5.Event{Event: "map_picked"},
				MatchID:   "14272",
				Team:      "team1",
				MapName:   "de_nuke",
				MapNumber: 0,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "map_picked",
				"matchid": "14272",
				"team": "team1",
				"map_name": "de_nuke",
				"map_number": 0
			  }`),
		},
		{
			title: "MapVetoed",
			eventHandler: &mockEventHandler{expect: got5.OnMapVetoedPayload{
				Event:   got5.Event{Event: "map_vetoed"},
				MatchID: "14272",
				Team:    "team1",
				MapName: "de_nuke",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "map_vetoed",
				"matchid": "14272",
				"team": "team1",
				"map_name": "de_nuke"
			  }`),
		},
		{
			title: "BackupRestore",
			eventHandler: &mockEventHandler{expect: got5.OnBackupRestorePayload{
				Event:       got5.Event{Event: "backup_loaded"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				Filename:    "string",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "backup_loaded",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"filename": "string"
			  }`),
		},
		{
			title: "DemoFinished",
			eventHandler: &mockEventHandler{expect: got5.OnDemoFinishedPayload{
				Event:     got5.Event{Event: "demo_finished"},
				MatchID:   "14272",
				MapNumber: 0,
				Filename:  "1324_map_0_de_nuke.dem",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "demo_finished",
				"matchid": "14272",
				"map_number": 0,
				"filename": "1324_map_0_de_nuke.dem"
			  }`),
		},
		{
			title: "DemoUploadEnded",
			eventHandler: &mockEventHandler{expect: got5.OnDemoUploadEndedPayload{
				Event:     got5.Event{Event: "demo_upload_ended"},
				MatchID:   "14272",
				MapNumber: 0,
				Filename:  "1324_map_0_de_nuke.dem",
				Success:   true,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "demo_upload_ended",
				"matchid": "14272",
				"map_number": 0,
				"filename": "1324_map_0_de_nuke.dem",
				"success": true
			  }`),
		},
		{
			title: "GamePaused",
			eventHandler: &mockEventHandler{expect: got5.OnMatchPausedPayload{
				Event:     got5.Event{Event: "game_paused"},
				MatchID:   "14272",
				MapNumber: 0,
				Team:      "team1",
				PauseType: "tactical",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "game_paused",
				"matchid": "14272",
				"map_number": 0,
				"team": "team1",
				"pause_type": "tactical"
			  }`),
		},
		{
			title: "Get5_OnPauseBegan",
			eventHandler: &mockEventHandler{expect: got5.OnPauseBeganPayload{
				Event:     got5.Event{Event: "pause_began"},
				MatchID:   "14272",
				MapNumber: 0,
				Team:      "team1",
				PauseType: "tactical",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "pause_began",
				"matchid": "14272",
				"map_number": 0,
				"team": "team1",
				"pause_type": "tactical"
			  }`),
		},
		{
			title: "KnifeStart",
			eventHandler: &mockEventHandler{expect: got5.OnKnifeRoundStartedPayload{
				Event:     got5.Event{Event: "knife_start"},
				MatchID:   "14272",
				MapNumber: 0,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "knife_start",
				"matchid": "14272",
				"map_number": 0
			  }`),
		},
		{
			title: "KnifeRoundWon",
			eventHandler: &mockEventHandler{expect: got5.OnKnifeRoundWonPayload{
				Event:     got5.Event{Event: "knife_won"},
				MatchID:   "14272",
				MapNumber: 0,
				Team:      "team1",
				Side:      "ct",
				Swapped:   true,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "knife_won",
				"matchid": "14272",
				"map_number": 0,
				"team": "team1",
				"side": "ct",
				"swapped": true
			  }`),
		},
		{
			title: "TeamReadyStatusChanged",
			eventHandler: &mockEventHandler{expect: got5.OnTeamReadyStatusChangedPayload{
				Event:     got5.Event{Event: "team_ready_status_changed"},
				MatchID:   "14272",
				Team:      getStringPointer("team1"),
				Ready:     true,
				GameState: "none",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "team_ready_status_changed",
				"matchid": "14272",
				"team": "team1",
				"ready": true,
				"game_state": "none"
			  }`),
		},
		{
			title: "GoingLive",
			eventHandler: &mockEventHandler{expect: got5.OnGoingLivePayload{
				Event:     got5.Event{Event: "going_live"},
				MatchID:   "14272",
				MapNumber: 0,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "going_live",
				"matchid": "14272",
				"map_number": 0
			  }`),
		},
		{
			title: "RoundStart",
			eventHandler: &mockEventHandler{expect: got5.OnRoundStartPayload{
				Event:       got5.Event{Event: "round_start"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "round_start",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13
			  }`),
		},
		{
			title: "RoundEnd",
			eventHandler: &mockEventHandler{expect: got5.OnRoundEndPayload{
				Event:       got5.Event{Event: "round_end"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Reason:      0,
				Winner:      got5.Winner{Side: "ct", Team: "team1"},
				Team1: got5.Get5StatsTeam{
					ID:          "2843",
					Name:        "Natus Vincere",
					SeriesScore: 0,
					Score:       14,
					ScoreCt:     10,
					ScoreT:      14,
					Players: []got5.Get5StatsPlayer{
						{
							SteamID: "76561198279375306",
							Name:    "s1mple",
							Stats: struct {
								Kills             int "json:\"kills\""
								Deaths            int "json:\"deaths\""
								Assists           int "json:\"assists\""
								FlashAssists      int "json:\"flash_assists\""
								TeamKills         int "json:\"team_kills\""
								Suicides          int "json:\"suicides\""
								Damage            int "json:\"damage\""
								UtilityDamage     int "json:\"utility_damage\""
								EnemiesFlashed    int "json:\"enemies_flashed\""
								FriendliesFlashed int "json:\"friendlies_flashed\""
								KnifeKills        int "json:\"knife_kills\""
								HeadshotKills     int "json:\"headshot_kills\""
								RoundsPlayed      int "json:\"rounds_played\""
								BombDefuses       int "json:\"bomb_defuses\""
								BombPlants        int "json:\"bomb_plants\""
								OneK              int "json:\"1k\""
								TwoK              int "json:\"2k\""
								ThreeK            int "json:\"3k\""
								FourK             int "json:\"4k\""
								FiveK             int "json:\"5k\""
								OneV1             int "json:\"1v1\""
								OneV2             int "json:\"1v2\""
								OneV3             int "json:\"1v3\""
								OneV4             int "json:\"1v4\""
								OneV5             int "json:\"1v5\""
								FirstKillsT       int "json:\"first_kills_t\""
								FirstKillsCt      int "json:\"first_kills_ct\""
								FirstDeathsT      int "json:\"first_deaths_t\""
								FirstDeathsCt     int "json:\"first_deaths_ct\""
								TradeKills        int "json:\"trade_kills\""
								KAST              int "json:\"kast\""
								Score             int "json:\"score\""
								Mvp               int "json:\"mvp\""
							}{
								Kills:             34,
								Deaths:            8,
								Assists:           5,
								FlashAssists:      3,
								TeamKills:         0,
								Suicides:          0,
								Damage:            2948,
								UtilityDamage:     173,
								EnemiesFlashed:    6,
								FriendliesFlashed: 2,
								KnifeKills:        1,
								HeadshotKills:     19,
								RoundsPlayed:      27,
								BombDefuses:       4,
								BombPlants:        3,
								OneK:              3,
								TwoK:              2,
								ThreeK:            3,
								FourK:             0,
								FiveK:             1,
								OneV1:             1,
								OneV2:             3,
								OneV3:             2,
								OneV4:             0,
								OneV5:             1,
								FirstKillsT:       6,
								FirstKillsCt:      5,
								FirstDeathsT:      1,
								FirstDeathsCt:     1,
								TradeKills:        3,
								KAST:              23,
								Score:             45,
								Mvp:               4,
							},
						},
					},
					Side:         "ct",
					StartingSide: "ct",
				},
				Team2: got5.Get5StatsTeam{
					ID:          "2843",
					Name:        "Natus Vincere",
					SeriesScore: 0,
					Score:       14,
					ScoreCt:     10,
					ScoreT:      14,
					Players: []got5.Get5StatsPlayer{
						{
							SteamID: "76561198279375306",
							Name:    "s1mple",
							Stats: struct {
								Kills             int "json:\"kills\""
								Deaths            int "json:\"deaths\""
								Assists           int "json:\"assists\""
								FlashAssists      int "json:\"flash_assists\""
								TeamKills         int "json:\"team_kills\""
								Suicides          int "json:\"suicides\""
								Damage            int "json:\"damage\""
								UtilityDamage     int "json:\"utility_damage\""
								EnemiesFlashed    int "json:\"enemies_flashed\""
								FriendliesFlashed int "json:\"friendlies_flashed\""
								KnifeKills        int "json:\"knife_kills\""
								HeadshotKills     int "json:\"headshot_kills\""
								RoundsPlayed      int "json:\"rounds_played\""
								BombDefuses       int "json:\"bomb_defuses\""
								BombPlants        int "json:\"bomb_plants\""
								OneK              int "json:\"1k\""
								TwoK              int "json:\"2k\""
								ThreeK            int "json:\"3k\""
								FourK             int "json:\"4k\""
								FiveK             int "json:\"5k\""
								OneV1             int "json:\"1v1\""
								OneV2             int "json:\"1v2\""
								OneV3             int "json:\"1v3\""
								OneV4             int "json:\"1v4\""
								OneV5             int "json:\"1v5\""
								FirstKillsT       int "json:\"first_kills_t\""
								FirstKillsCt      int "json:\"first_kills_ct\""
								FirstDeathsT      int "json:\"first_deaths_t\""
								FirstDeathsCt     int "json:\"first_deaths_ct\""
								TradeKills        int "json:\"trade_kills\""
								KAST              int "json:\"kast\""
								Score             int "json:\"score\""
								Mvp               int "json:\"mvp\""
							}{
								Kills:             34,
								Deaths:            8,
								Assists:           5,
								FlashAssists:      3,
								TeamKills:         0,
								Suicides:          0,
								Damage:            2948,
								UtilityDamage:     173,
								EnemiesFlashed:    6,
								FriendliesFlashed: 2,
								KnifeKills:        1,
								HeadshotKills:     19,
								RoundsPlayed:      27,
								BombDefuses:       4,
								BombPlants:        3,
								OneK:              3,
								TwoK:              2,
								ThreeK:            3,
								FourK:             0,
								FiveK:             1,
								OneV1:             1,
								OneV2:             3,
								OneV3:             2,
								OneV4:             0,
								OneV5:             1,
								FirstKillsT:       6,
								FirstKillsCt:      5,
								FirstDeathsT:      1,
								FirstDeathsCt:     1,
								TradeKills:        3,
								KAST:              23,
								Score:             45,
								Mvp:               4,
							},
						},
					},
					Side:         "ct",
					StartingSide: "ct",
				},
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "round_end",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"reason": 0,
				"winner": {
				  "side": "ct",
				  "team": "team1"
				},
				"team1": {
				  "id": "2843",
				  "name": "Natus Vincere",
				  "series_score": 0,
				  "score": 14,
				  "score_ct": 10,
				  "score_t": 14,
				  "players": [
					{
					  "steamid": "76561198279375306",
					  "name": "s1mple",
					  "stats": {
						"kills": 34,
						"deaths": 8,
						"assists": 5,
						"flash_assists": 3,
						"team_kills": 0,
						"suicides": 0,
						"damage": 2948,
						"utility_damage": 173,
						"enemies_flashed": 6,
						"friendlies_flashed": 2,
						"knife_kills": 1,
						"headshot_kills": 19,
						"rounds_played": 27,
						"bomb_defuses": 4,
						"bomb_plants": 3,
						"1k": 3,
						"2k": 2,
						"3k": 3,
						"4k": 0,
						"5k": 1,
						"1v1": 1,
						"1v2": 3,
						"1v3": 2,
						"1v4": 0,
						"1v5": 1,
						"first_kills_t": 6,
						"first_kills_ct": 5,
						"first_deaths_t": 1,
						"first_deaths_ct": 1,
						"trade_kills": 3,
						"kast": 23,
						"score": 45,
						"mvp": 4
					  }
					}
				  ],
				  "side": "ct",
				  "starting_side": "ct"
				},
				"team2": {
				  "id": "2843",
				  "name": "Natus Vincere",
				  "series_score": 0,
				  "score": 14,
				  "score_ct": 10,
				  "score_t": 14,
				  "players": [
					{
					  "steamid": "76561198279375306",
					  "name": "s1mple",
					  "stats": {
						"kills": 34,
						"deaths": 8,
						"assists": 5,
						"flash_assists": 3,
						"team_kills": 0,
						"suicides": 0,
						"damage": 2948,
						"utility_damage": 173,
						"enemies_flashed": 6,
						"friendlies_flashed": 2,
						"knife_kills": 1,
						"headshot_kills": 19,
						"rounds_played": 27,
						"bomb_defuses": 4,
						"bomb_plants": 3,
						"1k": 3,
						"2k": 2,
						"3k": 3,
						"4k": 0,
						"5k": 1,
						"1v1": 1,
						"1v2": 3,
						"1v3": 2,
						"1v4": 0,
						"1v5": 1,
						"first_kills_t": 6,
						"first_kills_ct": 5,
						"first_deaths_t": 1,
						"first_deaths_ct": 1,
						"trade_kills": 3,
						"kast": 23,
						"score": 45,
						"mvp": 4
					  }
					}
				  ],
				  "side": "ct",
				  "starting_side": "ct"
				}
			  }`),
		},
		{
			title: "RoundStatsUpdated",
			eventHandler: &mockEventHandler{expect: got5.OnRoundStatsUpdatedPayload{
				Event:       got5.Event{Event: "stats_updated"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "stats_updated",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13
			  }`),
		},
		{
			title: "PlayerBecomeMVP",
			eventHandler: &mockEventHandler{expect: got5.OnPlayerBecameMVPPayload{
				Event:       got5.Event{Event: "round_mvp"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false,
				},
				Reason: 0,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "round_mvp",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"reason": 0
			  }`),
		},
		{
			title: "GreanadeThrown",
			eventHandler: &mockEventHandler{expect: got5.OnGrenadeThrownPayload{
				Event:       got5.Event{Event: "grenade_thrown"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false},
				Weapon: got5.Weapon{
					Name: "ak47",
					ID:   27,
				},
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "grenade_thrown",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"weapon": {
				  "name": "ak47",
				  "id": 27
				}
			  }`),
		},

		// Penetrated should be float64, but example in https://splewis.github.io/get5/latest/events_and_forwards/#events shows true
		/*
			{
				title:        "PlayerDeath",
				eventHandler: &mockEventHandler{expect: got5.OnPlayerDeathPayload{Event: got5.Event{Event: "player_death"}, Matchid: "14272", MapNumber: 0, RoundNumber: 13, RoundTime: 51434, Player: got5.Player{UserID: 4, Steamid: "76561198279375306", Side: "ct", Name: "s1mple", IsBot: false}, Weapon: got5.Weapon{Name: "ak47", ID: 27}, Bomb: true, Headshot: true, ThruSmoke: true, Penetrated: 1 , AttackerBlind: true, NoScope: true, Suicide: true, FriendlyFire: true, Attacker: got5.Attacker{UserID: 4, Steamid: "76561198279375306", Side: "ct", Name: "s1mple", IsBot: false}, Assist: got5.Assist{Player: got5.Player{UserID: 4, Steamid: "76561198279375306", Side: "ct", Name: "s1mple", IsBot: false}, FriendlyFire: true, FlashAssist: true}}},
				auth:         &mockAuth{},
				statusCode:   http.StatusOK,
				input:        []byte(`{"event":"player_death","matchid":"14272","map_number":0,"round_number":13,"round_time":51434,"player":{"user_id":4,"steamid":"76561198279375306","side":"ct","name":"s1mple","is_bot":false},"weapon":{"name":"ak47","id":27},"bomb":true,"headshot":true,"thru_smoke":true,"penetrated":true,"attacker_blind":true,"no_scope":true,"suicide":true,"friendly_fire":true,"attacker":{"user_id":4,"steamid":"76561198279375306","side":"ct","name":"s1mple","is_bot":false},"assist":{"player":{"user_id":4,"steamid":"76561198279375306","side":"ct","name":"s1mple","is_bot":false},"friendly_fire":true,"flash_assist":true}}`),
			},
		*/

		{
			title: "HEGrenadeDetonated",
			eventHandler: &mockEventHandler{expect: got5.OnHEGrenadeDetonatedPayload{
				Event:       got5.Event{Event: "hegrenade_detonated"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					Side:    "ct",
					UserID:  4,
					IsBot:   false,
				},
				Weapon: got5.Weapon{
					Name: "ak47",
					ID:   27,
				},
				Victims: []got5.Victim{
					{
						Player: got5.Player{
							SteamID: "76561198279375306",
							Name:    "s1mple",
							UserID:  4,
							Side:    "ct",
							IsBot:   false,
						},
						FriendlyFire:  true,
						BlindDuration: 0.5,
					},
				},
				DamageEnemies:    0,
				DamageFriendlies: 0,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "hegrenade_detonated",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"weapon": {
				  "name": "ak47",
				  "id": 27
				},
				"victims": [
				  {
					"player": {
					  "steamid": "76561198279375306",
					  "name": "s1mple",
					  "user_id": 4,
					  "side": "ct",
					  "is_bot": false
					},
					"friendly_fire": true,
					"blind_duration": 0.5
				  }
				],
				"damage_enemies": 0,
				"damage_friendlies": 0
			  }`),
		},
		{
			title: "MolotovDetonated",
			eventHandler: &mockEventHandler{expect: got5.OnMolotovDetonatedPayload{
				Event:       got5.Event{Event: "molotov_detonated"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false,
				},
				Weapon: got5.Weapon{
					Name: "ak47",
					ID:   27,
				},
				Victims: []got5.Victim{
					{
						Player: got5.Player{
							SteamID: "76561198279375306",
							Name:    "s1mple",
							UserID:  4,
							Side:    "ct",
							IsBot:   false,
						},
						FriendlyFire:  true,
						BlindDuration: 0.5,
					},
				},
				DamageEnemies:    0,
				DamageFriendlies: 0,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "molotov_detonated",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"weapon": {
				  "name": "ak47",
				  "id": 27
				},
				"victims": [
				  {
					"player": {
					  "steamid": "76561198279375306",
					  "name": "s1mple",
					  "user_id": 4,
					  "side": "ct",
					  "is_bot": false
					},
					"friendly_fire": true,
					"blind_duration": 0.5
				  }
				],
				"damage_enemies": 0,
				"damage_friendlies": 0
			  }`),
		},
		{
			title: "FlashbangDetonated",
			eventHandler: &mockEventHandler{expect: got5.OnFlashbangDetonatedPayload{
				Event:       got5.Event{Event: "flashbang_detonated"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false,
				},
				Weapon: got5.Weapon{
					Name: "ak47",
					ID:   27,
				},
				Victims: []got5.Victim{
					{
						Player: got5.Player{SteamID: "76561198279375306",
							Name:   "s1mple",
							UserID: 4,
							Side:   "ct",
							IsBot:  false,
						},
						FriendlyFire:  true,
						BlindDuration: 0.5,
					},
				},
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "flashbang_detonated",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"weapon": {
				  "name": "ak47",
				  "id": 27
				},
				"victims": [
				  {
					"player": {
					  "steamid": "76561198279375306",
					  "name": "s1mple",
					  "user_id": 4,
					  "side": "ct",
					  "is_bot": false
					},
					"friendly_fire": true,
					"blind_duration": 0.5
				  }
				]
			  }`),
		},
		{
			title: "SmokegrenadeDetonated",
			eventHandler: &mockEventHandler{expect: got5.OnSmokeGrenadeDetonatedPayload{
				Event:       got5.Event{Event: "smokegrenade_detonated"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false,
				},
				Weapon: got5.Weapon{
					Name: "ak47",
					ID:   27,
				},
				ExtinguishedMolotov: true,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "smokegrenade_detonated",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"weapon": {
				  "name": "ak47",
				  "id": 27
				},
				"extinguished_molotov": true
			  }`),
		},
		{
			title: "DecoyStarted",
			eventHandler: &mockEventHandler{expect: got5.OnDecoyStartedPayload{
				Event:       got5.Event{Event: "decoygrenade_started"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false,
				},
				Weapon: got5.Weapon{
					Name: "ak47",
					ID:   27,
				},
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "decoygrenade_started",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"weapon": {
				  "name": "ak47",
				  "id": 27
				}
			  }`),
		},
		{
			title: "BombPlanted",
			eventHandler: &mockEventHandler{expect: got5.OnBombPlantedPayload{
				Event:       got5.Event{Event: "bomb_planted"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false},
				Site: "a",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "bomb_planted",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"site": "a"
			  }`),
		},
		{
			title: "BombDefused",
			eventHandler: &mockEventHandler{expect: got5.OnBombDefusedPayload{
				Event:       got5.Event{Event: "bomb_defused"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false},
				Site:              "a",
				BombTimeRemaining: 12438,
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "bomb_defused",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"site": "a",
				"bomb_time_remaining": 12438
			  }`),
		},
		{
			title: "BombExploded",
			eventHandler: &mockEventHandler{expect: got5.OnBombExplodedPayload{
				Event:       got5.Event{Event: "bomb_exploded"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Site:        "a",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "bomb_exploded",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"site": "a"
			  }`),
		},
		{
			title: "PlayerConnected",
			eventHandler: &mockEventHandler{expect: got5.OnPlayerConnectedPayload{
				Event:   got5.Event{Event: "player_connect"},
				MatchID: "14272",
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false,
				},
				IPAddress: "34.132.182.66",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "player_connect",
				"matchid": "14272",
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"ip_address": "34.132.182.66"
			  }`),
		},
		{
			title: "PlayerDisconnected",
			eventHandler: &mockEventHandler{expect: got5.OnPlayerDisconnectedPayload{
				Event:   got5.Event{Event: "player_disconnect"},
				MatchID: "14272",
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false,
				},
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "player_disconnect",
				"matchid": "14272",
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				}
			  }`),
		},
		{
			title: "PlayerSay",
			eventHandler: &mockEventHandler{expect: got5.OnPlayerSayPayload{
				Event:       got5.Event{Event: "player_say"},
				MatchID:     "14272",
				MapNumber:   0,
				RoundNumber: 13,
				RoundTime:   51434,
				Player: got5.Player{
					SteamID: "76561198279375306",
					Name:    "s1mple",
					UserID:  4,
					Side:    "ct",
					IsBot:   false,
				},
				Command: "say",
				Message: "gg have fun",
			}},
			auth:       &mockAuth{},
			statusCode: http.StatusOK,
			input: []byte(`{
				"event": "player_say",
				"matchid": "14272",
				"map_number": 0,
				"round_number": 13,
				"round_time": 51434,
				"player": {
				  "steamid": "76561198279375306",
				  "name": "s1mple",
				  "user_id": 4,
				  "side": "ct",
				  "is_bot": false
				},
				"command": "say",
				"message": "gg have fun"
			  }`),
		},
	}

	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			// Setup fiber
			app := gin.New()
			g5test := app.Group("/get5testevent") // /test
			ginroute.SetupEventHandlers(tt.eventHandler, tt.auth, g5test)

			r := httptest.NewRequest("POST", "/get5testevent/event", bytes.NewBuffer(tt.input))
			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			app.ServeHTTP(w, r)

			asserts.Equal(tt.statusCode, w.Code)
			asserts.Equal(tt.eventHandler.expect, tt.eventHandler.parsed)
		})
	}
}
