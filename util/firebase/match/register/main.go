package main

import (
	"context"
	"flag"
	"log"

	firebase "firebase.google.com/go"

	fsc "github.com/FlowingSPDG/Got5/controller/firebase"
	"github.com/FlowingSPDG/Got5/models"
)

var (
	projectID = ""
)

func main() {
	// Parse -projectID flag
	flag.StringVar(&projectID, "projectID", "", "Firestore project ID")
	flag.Parse()

	// Get Firebase service
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: projectID,
	}
	fb, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	// Get Controller connected to firestore
	ctrl, err := fsc.NewDatabase(ctx, fb)
	if err != nil {
		log.Fatalln(err)
	}

	m := models.Match{
		MatchTitle: "[TEST] Astralis vs. NaVi",
		// MatchID:              "", // 自動生成に任せるので設定しない
		ClinchSeries:         true,
		NumMaps:              3,
		PlayersPerTeam:       5,
		CoachesPerTeam:       2,
		CoachesMustReady:     true,
		MinPlayersToReady:    10,
		MinSpectatorsToReady: 0,
		SkipVeto:             false,
		VetoFirst:            "team1",
		SideType:             "standard",
		Spectators: models.Spectators{
			Name: "Blast PRO 2021",
			Players: map[string]string{
				"76561197987511774": "Anders Blume",
			},
		},
		Maplist:  []string{"de_dust2", "de_nuke", "de_inferno", "de_mirage", "de_vertigo", "de_ancient", "de_overpass"},
		MapSides: []string{"team1_ct", "team2_ct", "knife"},
		Team1: models.Team{
			Name: "Natus Vincere",
			Tag:  "NaVi",
			Flag: "UA",
			Logo: "navi",
			Players: map[string]string{
				"76561198034202275": "s1mple",
				"76561198044045107": "electronic",
				"76561198246607476": "b1t",
				"76561198121220486": "Perfecto",
				"76561198040577200": "sdy",
			},
			Coaches: map[string]string{
				"76561198013523865": "B1ad3",
			},
		},
		Team2: models.Team{
			Name: "Astralis",
			Tag:  "Astralis",
			Flag: "DK",
			Logo: "astr",
			Players: map[string]string{
				"76561197990682262": "Xyp9x",
				"76561198010511021": "gla1ve",
				"76561197979669175": "K0nfig",
				"76561198028458803": "BlameF",
				"76561198024248129": "farlig",
			},
			Coaches: map[string]string{
				"76561197987144812": "Trace",
			},
		},
		Cvars: map[string]string{
			"hostname":                       "Get5 Match #3123",
			"mp_friendly_fire":               "0",
			"get5_stop_command_enabled":      "0",
			"sm_practicemode_can_be_started": "0",
		},
	}

	m, err = ctrl.RegisterMatch(ctx, m)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Match Registered:", m)
}
