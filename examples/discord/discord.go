package discord

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"

	got5 "github.com/FlowingSPDG/Got5"
)

var _ got5.EventHandler = (*Discord)(nil)
var _ got5.MatchLoader = (*Discord)(nil)
var _ got5.DemoUploader = (*Discord)(nil)

// ControllerSetting Settings
type ControllerSetting struct {
	Hostname string
	Port     int
	password string
}

// Discord is Automated CS:GO/get5 BOT.
// 外部のDBなどを使わずに動くため、永続化したいのであればFirebaseなどのクライアントへ保存することを推奨
// discord implements all of EventHandler, Database, MatchLoader, DemoUploader.
type Discord struct {
	setting ControllerSetting
	s       *discordgo.Session // Session本体
	mu      sync.RWMutex
	matches map[string]struct {
		interaction *discordgo.Interaction // Interaction
		member      *discordgo.Member      // マッチ作成を実行したユーザー
		match       got5.Match             // GET5自体のマッチ情報
	} // GET5のマッチID(=interactionID))に対応したマッチ情報
}

// Hostname implements got5.EventHandler
func (d *Discord) Hostname() string {
	return d.setting.Hostname
}

// Close implements got5.EventHandler
func (d *Discord) Close() error {
	return d.s.Close()
}

type modalCreateMatchForm struct {
	MatchTitle string
	Team1Name  string
	Team2Name  string
}

func getCreateMatchModalFormBySubmitted(d discordgo.ModalSubmitInteractionData) modalCreateMatchForm {
	return modalCreateMatchForm{
		MatchTitle: d.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		Team1Name:  d.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
	}
}

type modalAddPlayerForm struct {
	MatchID   string
	Team      string
	Name      string
	SteamID64 string
}

func getAddPlayerModalFormBySubmitted(d discordgo.ModalSubmitInteractionData) modalAddPlayerForm {
	return modalAddPlayerForm{
		MatchID:   d.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		Team:      d.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		Name:      d.Components[2].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
		SteamID64: d.Components[3].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
	}
}

func getInteractionIDByAddModal(d discordgo.ModalSubmitInteractionData) string {
	return d.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
}

// NewDiscord Get discord pointer
func NewDiscord(ctx context.Context, token string, setting ControllerSetting) (*Discord, error) {
	d, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	c := &Discord{
		setting: setting,
		s:       d,
		mu:      sync.RWMutex{},
		matches: make(map[string]struct {
			interaction *discordgo.Interaction
			member      *discordgo.Member
			match       got5.Match
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

	cmd, err = c.s.ApplicationCommandCreate(c.s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "get5_addplayer", // TODO: Localization
		Description: "GET5のマッチにプレイヤーを追加します",
	})
	_ = cmd // 特に使わないので捨てる

	// TODO: 過去に登録したマッチをリストで表示できるようにする

	// モーダルからの情報送信時に実行されるイベント
	c.s.AddHandler(func(s *discordgo.Session, m *discordgo.InteractionCreate) {
		// 受け取ったModalからのデータを元にマッチを作成する
		// SteamIDを取得して登録、get5_loadmatch_url を発行して返す
		// マッチデータが無限に増えてしまうので、有効期限を設定して過ぎたらmapから削除でも良いかもしれない

		// Modalの発生コマンドか、Modalから送られてきたイベントかで分岐

		switch m.Type {
		case discordgo.InteractionApplicationCommand:
			switch m.ApplicationCommandData().Name {
			case "get5_create":
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
										CustomID:  "team2name",
										Label:     "Team2 Name",
										Style:     discordgo.TextInputShort,
										Required:  true,
										MaxLength: 30,
									},
								},
							},
						},
					}}); err != nil {
					fmt.Println("err:", err)
				}
			case "get5_addplayer":
				if err := s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseModal,
					Data: &discordgo.InteractionResponseData{
						CustomID: "get5_addplayer_" + m.Interaction.Member.User.ID,
						Title:    "GET5 Add Player",
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.TextInput{
										CustomID:    "match_id",
										Label:       "Match ID",
										Style:       discordgo.TextInputShort,
										Placeholder: "",
										Required:    true,
										MaxLength:   300,
										MinLength:   1,
									},
								},
							},

							// 手動入力させるのは美しくないので別の方法を考えたい
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.TextInput{
										CustomID:    "team",
										Label:       "Team to add (team1, team2 or spectator)",
										Style:       discordgo.TextInputShort,
										Placeholder: "team1",
										Required:    true,
										MaxLength:   9,
										MinLength:   5,
									},
								},
							},

							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.TextInput{
										CustomID:    "name",
										Label:       "Player Name",
										Style:       discordgo.TextInputShort,
										Placeholder: "John",
										Required:    true,
										MaxLength:   30,
										MinLength:   2,
									},
								},
							},

							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									discordgo.TextInput{
										CustomID:    "steamid",
										Label:       "Player's SteamID(SteamID64 recommended)",
										Style:       discordgo.TextInputShort,
										Placeholder: "",
										Required:    true,
										MaxLength:   17,
										MinLength:   17,
									},
								},
							},
						},
					},
				}); err != nil {
					fmt.Println("err:", err)
				}
			}

			// Modalの送信実行時
		case discordgo.InteractionModalSubmit:
			// データを取得してMatchを作成する
			data := m.ModalSubmitData()
			// get5_create の場合
			if strings.HasPrefix(data.CustomID, "get5_create_") {
				// m.User, m.Message など、nilになる値があるので注意
				mf := getCreateMatchModalFormBySubmitted(data)
				match := got5.GetDefaultMatchBO1()
				match.MatchTitle = fmt.Sprintf("%s : Created by %s", mf.MatchTitle, m.Member.Nick)
				match.MatchID = m.Interaction.ID
				match.Team1.Name = mf.Team1Name
				c.mu.Lock()
				c.matches[m.Interaction.ID] = struct {
					interaction *discordgo.Interaction
					member      *discordgo.Member
					match       got5.Match
				}{
					interaction: m.Interaction,
					member:      m.Member,
					match:       match,
				}
				c.mu.Unlock()

				if _, err := c.RegisterMatch(ctx, match); err != nil {
					fmt.Println("err:", err) // 作成に失敗した旨を発言する
					return
				}

				content := fmt.Sprintf("マッチを作成しました\n ``get5_loadmatch_url \"http://%s:%d/get5/match/%s\"``", c.setting.Hostname, c.setting.Port, match.MatchID)
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
			} else if strings.HasPrefix(data.CustomID, "get5_addplayer_") {
				f := getAddPlayerModalFormBySubmitted(data)
				c.mu.Lock()
				defer c.mu.Unlock()
				match, ok := c.matches[f.MatchID]
				if !ok {
					fmt.Println("No match found") // 作成に失敗した旨を発言する
					return
				}
				switch f.Team {
				case "team1":
					match.match.Team1.Players[f.SteamID64] = f.Name
				case "team2":
					match.match.Team2.Players[f.SteamID64] = f.Name
				case "spectator":
					match.match.Spectators.Players[f.SteamID64] = f.Name
				default:
					fmt.Println("Please specify team to add:") // 作成に失敗した旨を発言する
					return
				}

				content := fmt.Sprintf("プレイヤー``%s``を追加しました", f.Name)
				err := s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: content,
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				if err != nil {
					fmt.Println("err:", err)
					return
				}
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
