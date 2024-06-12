package got5

// get5: https://splewis.github.io/get5/latest/match_schema/#schema
// MatchZy: https://shobhit-pathak.github.io/MatchZy/match_setup/#example
// MatchZy(source): https://github.com/shobhit-pathak/MatchZy/blob/ae597e50756f8ed8380e4cbacec4a4c1eb9013da/MatchManagement.cs#L254

// MatchID currently int only
type MatchID interface {
	int
}

type Match[TmatchID MatchID] struct {
	MatchID              TmatchID          `json:"matchid,omitempty"` // originally string in get5 but MatchZy uses int
	ClinchSeries         bool              `json:"clinch_series,omitempty"`
	NumMaps              int               `json:"num_maps,omitempty"`
	Wingman              bool              `json:"wingman,omitempty"`
	PlayersPerTeam       int               `json:"players_per_team,omitempty"`
	MinPlayersToReady    int               `json:"min_players_to_ready,omitempty"`
	MinSpectatorsToReady int               `json:"min_spectators_to_ready,omitempty"`
	SkipVeto             bool              `json:"skip_veto,omitempty"`
	VetoMode             []string          `json:"veto_mode,omitempty"`
	MapSides             []string          `json:"map_sides,omitempty"`
	Spectators           Spectators        `json:"spectators,omitempty"`
	Maplist              []string          `json:"maplist,omitempty"`
	Team1                Team              `json:"team1,omitempty"`
	Team2                Team              `json:"team2,omitempty"`
	Cvars                map[string]string `json:"cvars,omitempty"`
}

func (m Match[TMatchID]) ToG5Format() Match[TMatchID] {
	return m
}

func GetDefaultMatchBO1[TMatchID MatchID]() Match[TMatchID] {
	return Match[TMatchID]{
		MatchID:              *new(TMatchID),
		ClinchSeries:         true,
		NumMaps:              1,
		PlayersPerTeam:       5,
		MinPlayersToReady:    1,
		MinSpectatorsToReady: 0,
		SkipVeto:             false,
		Spectators: Spectators{
			Players: map[string]string{},
		},
		Maplist:  []string{"de_anubis", "de_nuke", "de_inferno", "de_mirage", "de_vertigo", "de_ancient", "de_overpass"},
		MapSides: []string{"knife"},
		Team1: Team{
			Name:    "",
			Players: map[string]string{},
		},
		Team2: Team{
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

// G5Match 別の構造体にG5Matchインターフェースを実装すれば型が違っても変換してGet5に渡してくれる
type G5Match[TMatchID MatchID] interface {
	ToG5Format() Match[TMatchID]
}