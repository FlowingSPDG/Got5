package models

// https://splewis.github.io/get5/latest/match_schema/#schema

type Match struct {
	MatchTitle             string            `json:"match_title"`
	MatchID                string            `json:"matchid"`
	ClinchSeries           bool              `json:"clinch_series"`
	NumMaps                int               `json:"num_maps"`
	Scrim                  bool              `json:"scrim"`
	Wingman                bool              `json:"wingman"`
	PlayersPerTeam         int               `json:"players_per_team"`
	CoachesPerTeam         int               `json:"coaches_per_team"`
	CoachesMustReady       bool              `json:"coaches_must_ready"`
	MinPlayersToReady      int               `json:"min_players_to_ready"`
	MinSpectatorsToReady   int               `json:"min_spectators_to_ready"`
	SkipVeto               bool              `json:"skip_veto"`
	VetoFirst              string            `json:"veto_first"`
	VetoMode               string            `json:"veto_mode"`
	SideType               string            `json:"side_type"`
	MapSides               []string          `json:"map_sides"`
	Spectators             Spectators        `json:"spectators"`
	Maplist                []string          `json:"maplist"`
	FavoredPercentageTeam1 int               `json:"favored_percentage_team1"`
	FavoredPercentageText  string            `json:"favored_percentage_text"`
	Team1                  Team              `json:"team1"`
	Team2                  Team              `json:"team2"`
	Cvars                  map[string]string `json:"cvars"`
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
	Name    string            `json:"name"`
	Players map[string]string `json:"players"`
}
type Team struct {
	ID          string            `json:"id"`
	Players     map[string]string `json:"players"`
	Coaches     map[string]string `json:"coaches"`
	Name        string            `json:"name"`
	Tag         string            `json:"tag"`
	Flag        string            `json:"flag"`
	Logo        string            `json:"logo"`
	SeriesScore int               `json:"series_score"`
	MatchText   string            `json:"matchtext"`
	FromFile    string            `json:"fromfile"`
}

// G5Match 別の構造体にG5Matchインターフェースを実装すれば型が違っても変換してGet5に渡してくれる
type G5Match interface {
	ToG5Format() Match
}
