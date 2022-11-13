package discord

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

// Discord Interaction(POST) の実装をしてもいいかもしれない

// discord is Automated CS:GO/get5 BOT.
type discord struct {
	s       *discordgo.Session // Session本体
	mu      sync.RWMutex
	matches map[string]struct {
		member *discordgo.Member // マッチ作成を実行したユーザー
		match  models.Match      // GET5自体のマッチ情報
	} // GET5のマッチID(=ユーザーID))に対応したマッチ情報
}

func (d *discord) Close() error {
	return d.s.Close()
}

type modalForm struct {
	MatchTitle    string
	Team1Name     string
	Team1SteamIDs []string
	Team2Name     string
	Team2SteamIDs []string
}

func getModalFormBySubmitted(d discordgo.ModalSubmitInteractionData) modalForm {
	t1steamids := strings.Split(d.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value, "\n")
	t2steamids := strings.Split(d.Components[3].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value, "\n")
	return modalForm{
		MatchTitle:    d.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		Team1Name:     d.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		Team1SteamIDs: t1steamids,
		Team2Name:     d.Components[3].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		Team2SteamIDs: t2steamids,
	}
}

// NewDiscordController Get discord pointer
func NewDiscordController(ctx context.Context, token string) (controller.Controller, error) {
	d, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	c := &discord{
		s: d,
		matches: make(map[string]struct {
			member *discordgo.Member
			match  models.Match
		}),
	}

	if err := c.s.Open(); err != nil {
		return nil, err
	}

	// コマンドを登録する
	cmd, err := c.s.ApplicationCommandCreate(c.s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "get5_create", // TODO: Localization
		Description: "GET5のマッチを作成します",
	})
	if err != nil {
		return nil, err
	}
	_ = cmd // 特に使わないので捨てる

	// モーダルからの情報送信時に実行されるイベント
	c.s.AddHandler(func(s *discordgo.Session, m *discordgo.InteractionCreate) {
		// 受け取ったModalからのデータを元にマッチを作成する
		// SteamIDを取得して登録、get5_loadmatch_url を発行して返す
		// マッチデータが無限に増えてしまうので、有効期限を設定して過ぎたらmapから削除でも良いかもしれない

		// Modalの発生コマンドか、Modalから送られてきたイベントかで分岐
		switch m.Type {
		case discordgo.InteractionApplicationCommand:
			if err := s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseModal,
				Data: &discordgo.InteractionResponseData{
					CustomID: "get5_create_" + m.Interaction.Member.User.ID,
					Title:    "GET5 Match Create",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:    "title",
									Label:       "MATCH TITLE",
									Style:       discordgo.TextInputShort,
									Placeholder: m.Member.Nick + "のMatch",
									Required:    true,
									MaxLength:   300,
									MinLength:   5,
								},
							},
						},

						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:  "team1name",
									Label:     "Team1 Name",
									Style:     discordgo.TextInputShort,
									Required:  true,
									MaxLength: 30,
								},
							},
						},

						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:  "team1steamids",
									Label:     "Team1 SteamIDs",
									Style:     discordgo.TextInputParagraph,
									Required:  true,
									MaxLength: 500,
								},
							},
						},

						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:  "team2name",
									Label:     "Team2 Name",
									Style:     discordgo.TextInputShort,
									Required:  true,
									MaxLength: 30,
								},
							},
						},

						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.TextInput{
									CustomID:  "team2steamids",
									Label:     "Team2 SteamIDs",
									Style:     discordgo.TextInputParagraph,
									Required:  true,
									MaxLength: 500,
								},
							},
						},
					},
				}}); err != nil {
				fmt.Println("err:", err)
			}

			// Modalの送信実行時
		case discordgo.InteractionModalSubmit:
			// データを取得してMatchを作成する
			data := m.ModalSubmitData()
			if !strings.HasPrefix(data.CustomID, "get5_create_") {
				return
			}

			mf := getModalFormBySubmitted(data)
			match := models.GetDefaultMatchBO1()
			match.MatchTitle = mf.MatchTitle
			match.MatchID = m.Interaction.Member.User.ID
			match.Team1.Name = mf.Team1Name
			// プレイヤーネームは仮で一旦SteamIDのみにする
			match.Team1.Players = map[string]string{}
			for _, v := range mf.Team1SteamIDs {
				match.Team1.Players[v] = v
			}
			match.Team2.Name = mf.Team2Name
			for _, v := range mf.Team2SteamIDs {
				match.Team2.Players[v] = v
			}
			c.mu.Lock()
			c.matches[m.Member.User.ID] = struct {
				member *discordgo.Member
				match  models.Match
			}{
				member: m.Member,
				match:  match,
			}
			c.mu.Unlock()

			if err := c.RegisterMatch(ctx, match); err != nil {
				fmt.Println("err:", err) // 作成に失敗した旨を発言する
			}

			// TODO: 127.0.0.1... の部分を差し替える
			content := fmt.Sprintf("マッチを作成しました\n ``get5_loadmatch_url \"http://127.0.0.1:3000/get5/match/%s\"``", match.MatchID)
			err := s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				fmt.Println("err:", err)
			}

		}
	})

	// In this example, we only care about receiving message events.
	c.s.Identify.Intents = discordgo.IntentsGuildMessages

	return c, nil
}

// chunkSlice スライスを分割するユーティリティ関数
func chunkSlice[T comparable](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
