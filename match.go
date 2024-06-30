package got5

// get5: https://splewis.github.io/get5/latest/match_schema/#schema
// MatchZy: https://shobhit-pathak.github.io/MatchZy/match_setup/#example
// MatchZy(source): https://github.com/shobhit-pathak/MatchZy/blob/ae597e50756f8ed8380e4cbacec4a4c1eb9013da/MatchManagement.cs#L254

type Match struct {
	MatchID              *int              `json:"matchid,omitempty"` // originally string in get5 but MatchZy uses int
	ClinchSeries         *bool             `json:"clinch_series,omitempty"`
	NumMaps              int               `json:"num_maps,omitempty"`
	Wingman              bool              `json:"wingman,omitempty"`
	PlayersPerTeam       *int              `json:"players_per_team,omitempty"`
	MinPlayersToReady    *int              `json:"min_players_to_ready,omitempty"`
	MinSpectatorsToReady *int              `json:"min_spectators_to_ready,omitempty"`
	SkipVeto             *bool             `json:"skip_veto,omitempty"`
	VetoMode             []string          `json:"veto_mode,omitempty"`
	MapSides             []string          `json:"map_sides,omitempty"`
	Spectators           *Spectators       `json:"spectators,omitempty"`
	Maplist              []string          `json:"maplist,omitempty"`
	Team1                *Team             `json:"team1,omitempty"`
	Team2                *Team             `json:"team2,omitempty"`
	Cvars                map[string]string `json:"cvars,omitempty"`
}

func GetDefaultMatchBO1() Match {
	return Match{
		MatchID:              nil,
		ClinchSeries:         nil,
		NumMaps:              1,
		PlayersPerTeam:       nil,
		MinPlayersToReady:    nil,
		MinSpectatorsToReady: nil,
		SkipVeto:             nil,
		Spectators:           nil,
		Maplist:              []string{"de_anubis", "de_nuke", "de_inferno", "de_mirage", "de_vertigo", "de_ancient", "de_overpass"},
		MapSides:             []string{"knife"},
		Team1: &Team{
			Name:    "",
			Players: map[string]string{},
		},
		Team2: &Team{
			Name:    "",
			Players: map[string]string{},
		},
		Cvars: map[string]string{},
	}
}

type Spectators struct {
	Players map[string]string `json:"players,omitempty"`
}
type Team struct {
	ID      string            `json:"id,omitempty"`
	Players map[string]string `json:"players,omitempty"`
	Name    string            `json:"name,omitempty"`
}
