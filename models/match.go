package models

// https://splewis.github.io/get5/latest/match_schema/#schema

type Match struct {
	MatchTitle           string            `json:"match_title"`
	MatchID              string            `json:"matchid"`
	ClinchSeries         bool              `json:"clinch_series"`
	NumMaps              int               `json:"num_maps"`
	PlayersPerTeam       int               `json:"players_per_team"`
	CoachesPerTeam       int               `json:"coaches_per_team"`
	CoachesMustReady     bool              `json:"coaches_must_ready"`
	MinPlayersToReady    int               `json:"min_players_to_ready"`
	MinSpectatorsToReady int               `json:"min_spectators_to_ready"`
	SkipVeto             bool              `json:"skip_veto"`
	VetoFirst            string            `json:"veto_first"`
	SideType             string            `json:"side_type"`
	Spectators           Spectators        `json:"spectators"`
	Maplist              []string          `json:"maplist"`
	MapSides             []string          `json:"map_sides"`
	Team1                Team              `json:"team1"`
	Team2                Team              `json:"team2"`
	Cvars                map[string]string `json:"cvars"`
}

type Spectators struct {
	Name    string            `json:"name"`
	Players map[string]string `json:"players"`
}
type Team struct {
	Name    string            `json:"name"`
	Tag     string            `json:"tag"`
	Flag    string            `json:"flag"`
	Logo    string            `json:"logo"`
	Players map[string]string `json:"players"`
	Coaches map[string]string `json:"coaches"`
}

// G5Match 別の構造体にG5Matchインターフェースを実装すれば型が違っても変換してGet5に渡してくれレウ
type G5Match interface {
	ToG5Format() Match
}
