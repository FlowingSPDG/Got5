package got5_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	got5 "github.com/FlowingSPDG/Got5"
)

func TestG5Match(t *testing.T) {
	asserts := assert.New(t)
	cases := []struct {
		title string
		input got5.Match[int]
	}{
		{
			title: "マッチのG5インターフェース変換テスト",
			input: got5.Match[int]{
				MatchID:              0,
				ClinchSeries:         false,
				NumMaps:              3,
				PlayersPerTeam:       5,
				MinPlayersToReady:    10,
				MinSpectatorsToReady: 0,
				SkipVeto:             false,
				Spectators:           got5.Spectators{},
				Maplist:              []string{"de_anubis", "de_nuke", "de_inferno", "de_mirage", "de_vertigo", "de_ancient", "de_overpass"},
				MapSides:             []string{"knife"},
				Team1:                got5.Team{},
				Team2:                got5.Team{},
				Cvars:                map[string]string{},
			},
		},
		{
			title: "複雑なマッチのパース",
			input: got5.Match[int]{
				MatchID:              1234567890,
				ClinchSeries:         true,
				NumMaps:              5,
				PlayersPerTeam:       3,
				MinPlayersToReady:    8,
				MinSpectatorsToReady: 2,
				SkipVeto:             true,
				Spectators: got5.Spectators{
					Players: map[string]string{
						"rjzefl": "flowingspdg",
					},
				},
				Maplist:  []string{"de_ancient", "de_overpass"},
				MapSides: []string{"knife"},
				Team1: got5.Team{
					Name: "Team1 Test",
					Players: map[string]string{
						"fl1": "flowingspdg",
					},
				},
				Team2: got5.Team{
					Name: "",
					Players: map[string]string{
						"fl2": "flowingspdg",
					},
				},
				Cvars: map[string]string{},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			asserts.Equal(tt.input.MatchID, tt.input.ToG5Format().MatchID)
			asserts.Equal(tt.input.ClinchSeries, tt.input.ToG5Format().ClinchSeries)
			asserts.Equal(tt.input.NumMaps, tt.input.ToG5Format().NumMaps)
			asserts.Equal(tt.input.PlayersPerTeam, tt.input.ToG5Format().PlayersPerTeam)
			asserts.Equal(tt.input.MinPlayersToReady, tt.input.ToG5Format().MinPlayersToReady)
			asserts.Equal(tt.input.MinSpectatorsToReady, tt.input.ToG5Format().MinSpectatorsToReady)
			asserts.Equal(tt.input.SkipVeto, tt.input.ToG5Format().SkipVeto)
			asserts.Equal(tt.input.Spectators, tt.input.ToG5Format().Spectators)
			asserts.Equal(tt.input.Maplist, tt.input.ToG5Format().Maplist)
			asserts.Equal(tt.input.MapSides, tt.input.ToG5Format().MapSides)
			asserts.Equal(tt.input.Team1, tt.input.ToG5Format().Team1)
			asserts.Equal(tt.input.Team2, tt.input.ToG5Format().Team2)
			asserts.Equal(tt.input.Cvars, tt.input.ToG5Format().Cvars)
		})
	}
}

// DefaultBO1
