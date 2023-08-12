package got5

// https://splewis.github.io/get5/latest/match_schema/#schema

type Match struct {
	MatchTitle             string            `json:"match_title,omitempty"`
	MatchID                string            `json:"matchid,omitempty"`
	ClinchSeries           bool              `json:"clinch_series,omitempty"`
	NumMaps                int               `json:"num_maps,omitempty"`
	Scrim                  bool              `json:"scrim,omitempty"`
	Wingman                bool              `json:"wingman,omitempty"`
	PlayersPerTeam         int               `json:"players_per_team,omitempty"`
	CoachesPerTeam         int               `json:"coaches_per_team,omitempty"`
	CoachesMustReady       bool              `json:"coaches_must_ready,omitempty"`
	MinPlayersToReady      int               `json:"min_players_to_ready,omitempty"`
	MinSpectatorsToReady   int               `json:"min_spectators_to_ready,omitempty"`
	SkipVeto               bool              `json:"skip_veto,omitempty"`
	VetoFirst              string            `json:"veto_first,omitempty"`
	VetoMode               string            `json:"veto_mode,omitempty"`
	SideType               string            `json:"side_type,omitempty"`
	MapSides               []string          `json:"map_sides,omitempty"`
	Spectators             Spectators        `json:"spectators,omitempty"`
	Maplist                []string          `json:"maplist,omitempty"`
	FavoredPercentageTeam1 int               `json:"favored_percentage_team1,omitempty"`
	FavoredPercentageText  string            `json:"favored_percentage_text,omitempty"`
	Team1                  Team              `json:"team1,omitempty"`
	Team2                  Team              `json:"team2,omitempty"`
	Cvars                  map[string]string `json:"cvars,omitempty"`
}

func (m Match) ToG5Format() Match {
	return m
}

func GetDefaultMatchBO1() Match {
	return Match{
		MatchTitle:           "",
		MatchID:              "",
		ClinchSeries:         true,
		NumMaps:              1,
		PlayersPerTeam:       5,
		CoachesPerTeam:       0,
		CoachesMustReady:     false,
		MinPlayersToReady:    1,
		MinSpectatorsToReady: 0,
		SkipVeto:             false,
		VetoFirst:            "team1",
		SideType:             "standard",
		Spectators: Spectators{
			Name:    "",
			Players: map[string]string{},
		},
		Maplist:  []string{"de_anubis", "de_nuke", "de_inferno", "de_mirage", "de_vertigo", "de_ancient", "de_overpass"},
		MapSides: []string{"knife"},
		Team1: Team{
			Name:    "",
			Tag:     "",
			Flag:    "",
			Logo:    "",
			Players: map[string]string{},
			Coaches: map[string]string{},
		},
		Team2: Team{
			Name:    "",
			Tag:     "",
			Flag:    "",
			Logo:    "",
			Players: map[string]string{},
			Coaches: map[string]string{},
		},
		Cvars: map[string]string{},
	}
}

type Spectators struct {
	Name    string            `json:"name,omitempty"`
	Players map[string]string `json:"players,omitempty"`
}
type Team struct {
	ID          string            `json:"id,omitempty"`
	Players     map[string]string `json:"players,omitempty"`
	Coaches     map[string]string `json:"coaches,omitempty"`
	Name        string            `json:"name,omitempty"`
	Tag         string            `json:"tag,omitempty"`
	Flag        string            `json:"flag,omitempty"`
	Logo        string            `json:"logo,omitempty"`
	SeriesScore int               `json:"series_score,omitempty"`
	MatchText   string            `json:"matchtext,omitempty"`
	FromFile    string            `json:"fromfile,omitempty"`
}

// G5Match 別の構造体にG5Matchインターフェースを実装すれば型が違っても変換してGet5に渡してくれる
type G5Match interface {
	ToG5Format() Match
}
