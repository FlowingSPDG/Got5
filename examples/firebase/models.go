package fb

import "github.com/FlowingSPDG/Got5/models"

// https://splewis.github.io/get5/latest/match_schema/#schema

// Match for Get5 match
type Match struct {
	MatchTitle           string            `firestore:"match_title"`
	MatchID              string            `firestore:"matchid"`
	ClinchSeries         bool              `firestore:"clinch_series"`
	NumMaps              int               `firestore:"num_maps"`
	PlayersPerTeam       int               `firestore:"players_per_team"`
	CoachesPerTeam       int               `firestore:"coaches_per_team"`
	CoachesMustReady     bool              `firestore:"coaches_must_ready"`
	MinPlayersToReady    int               `firestore:"min_players_to_ready"`
	MinSpectatorsToReady int               `firestore:"min_spectators_to_ready"`
	SkipVeto             bool              `firestore:"skip_veto"`
	VetoFirst            string            `firestore:"veto_first"`
	SideType             string            `firestore:"side_type"`
	Spectators           Spectators        `firestore:"spectators"`
	Maplist              []string          `firestore:"maplist"`
	MapSides             []string          `firestore:"map_sides"`
	Team1                Team              `firestore:"team1"`
	Team2                Team              `firestore:"team2"`
	Cvars                map[string]string `firestore:"cvars"`

	// Firetore向けに追加したいフィールドがあればここで追記する
	// Additional fields for firestore
	AuthValue string `firestore:"auth_value"`
}

// ToG5Format Convert into models.Match
func (m Match) ToG5Format() models.Match {
	return models.Match{
		MatchTitle:           m.MatchTitle,
		MatchID:              m.MatchID,
		ClinchSeries:         m.ClinchSeries,
		NumMaps:              m.NumMaps,
		PlayersPerTeam:       m.PlayersPerTeam,
		CoachesPerTeam:       m.CoachesPerTeam,
		CoachesMustReady:     m.CoachesMustReady,
		MinPlayersToReady:    m.MinPlayersToReady,
		MinSpectatorsToReady: m.MinSpectatorsToReady,
		SkipVeto:             m.SkipVeto,
		VetoFirst:            m.VetoFirst,
		SideType:             m.SideType,
		Spectators:           models.Spectators(m.Spectators),
		Maplist:              m.Maplist,
		MapSides:             m.MapSides,
		Team1:                models.Team(m.Team1),
		Team2:                models.Team(m.Team2),
		Cvars:                m.Cvars,
	}
}

// Spectators Spec team
type Spectators struct {
	Name    string            `firestore:"name"`
	Players map[string]string `firestore:"players"`
}

// Team for team1,team2
type Team struct {
	ID          string            `firestore:"id"`
	Players     map[string]string `firestore:"players"`
	Coaches     map[string]string `firestore:"coaches"`
	Name        string            `firestore:"name"`
	Tag         string            `firestore:"tag"`
	Flag        string            `firestore:"flag"`
	Logo        string            `firestore:"logo"`
	SeriesScore int               `firestore:"series_score"`
	MatchText   string            `firestore:"matchtext"`
	FromFile    string            `firestore:"fromfile"`
}

func toFirestoreMatch(m models.Match) Match {
	return Match{
		MatchTitle:           m.MatchTitle,
		MatchID:              m.MatchID,
		ClinchSeries:         m.ClinchSeries,
		NumMaps:              m.NumMaps,
		PlayersPerTeam:       m.PlayersPerTeam,
		CoachesPerTeam:       m.CoachesPerTeam,
		CoachesMustReady:     m.CoachesMustReady,
		MinPlayersToReady:    m.MinPlayersToReady,
		MinSpectatorsToReady: m.MinSpectatorsToReady,
		SkipVeto:             m.SkipVeto,
		VetoFirst:            m.VetoFirst,
		SideType:             m.SideType,
		Spectators:           Spectators(m.Spectators),
		Maplist:              m.Maplist,
		MapSides:             m.MapSides,
		Team1:                Team(m.Team1),
		Team2:                Team(m.Team2),
		Cvars:                m.Cvars,
	}
}
