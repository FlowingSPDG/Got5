package discord

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

// discord is Automated CS:GO/get5 BOT.
type discord struct {
	s       *discordgo.Session // Session本体
	mu      sync.RWMutex
	matches map[string]struct {
		msg   *discordgo.Message // マッチ作成に使ったメッセージ
		match models.Match       // GET5自体のマッチ情報
	} // GET5のマッチID(=メッセージID))に対応したマッチ情報
}

func (d *discord) Close() error {
	return d.s.Close()
}

// NewDiscordController Get discord pointer
func NewDiscordController(token string) (controller.Controller, error) {
	d, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	c := &discord{
		s: d,
		matches: make(map[string]struct {
			msg   *discordgo.Message
			match models.Match
		}),
	}

	if err := c.s.Open(); err != nil {
		return nil, err
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	c.s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// 自分が発信したメッセージを全て無視する
		if m.Author.ID == s.State.User.ID {
			return
		}
		// .get5 と打ち込まれた場合にマッチを作成する
		// 引数でSteamIDを取得して登録、get5_loadmatch_url を発行する
		if strings.HasPrefix(m.Content, ".get5") {
			// パース処理を実行
			ids := strings.Split(m.Content, " ")
			// IDをシャッフルする
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })

			// チームを分配
			// steamids := chunkSlice[string](ids, 2)
			// t1 := steamids[0]
			// t2 := steamids[1]

			c.mu.Lock()
			defer c.mu.Unlock()

			match := models.GetDefaultMatchBO1()
			match.MatchTitle = m.Author.Username + "の作成した紅白"
			match.MatchID = m.ID

			// SteamIDなどを使って情報を入れる
			match.Team1 = models.Team{
				Name: "",
				Tag:  "",
				Flag: "",
				Logo: "",
				Players: map[string]string{
					"": "",
				},
				Coaches: map[string]string{
					"": "",
				},
			}
			match.Team2 = models.Team{}

			c.matches[m.ID] = struct {
				msg   *discordgo.Message
				match models.Match
			}{
				msg:   m.Message,
				match: match,
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
