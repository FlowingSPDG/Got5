package webhook

import (
	"context"
	"fmt"
	"strconv"
	"time"

	// AWS package converts value to pointer in function, so this is not something about Amazon Web Service
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

var _ controller.EventHandler = (*webhookEventHandler)(nil)

// webhookEventHandler Discord webhook
type webhookEventHandler struct {
	eventURL string
	auth     string
	c        webhook.Client
}

// Hostname implements controller.EventHandler
func (wh *webhookEventHandler) Hostname() string {
	return wh.eventURL
}

func (wh *webhookEventHandler) Close() error {
	return nil
}

// HandleOnBackupRestore implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "FILENAME", Value: p.Filename, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnBombDefused implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				{Name: "SITE", Value: p.Site, Inline: aws.Bool(false)},
				{Name: "TIME REMAINING", Value: strconv.Itoa(p.BombTimeRemaining), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnBombExploded implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
				{Name: "SITE", Value: p.Site, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnBombPlanted implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				{Name: "SITE", Value: p.Site, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnDecoyStarted implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				{Name: "WEAPON", Value: p.Weapon.Name, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnDemoFinished implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "FILE NAME", Value: p.Filename, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnDemoUploadEnded implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "FILE NAME", Value: p.Filename, Inline: aws.Bool(false)},
				{Name: "SUCCESS", Value: strconv.FormatBool(p.Success), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnEvent implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnEvent(ctx context.Context, p models.OnEventPayload) error {
	// any event
	return nil
}

// HandleOnFlashbangDetonated implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	fields := []discord.EmbedField{
		{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
		{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
		{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
		{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
		{Name: "WEAPON", Value: p.Weapon.Name, Inline: aws.Bool(false)},
	}
	emb := []discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields:      fields,
		},
	}
	for i := 0; i < len(p.Victims); i++ {
		fields = append(fields, discord.EmbedField{
			Name: fmt.Sprintf("VICTIM [%d]", i), Value: p.Victims[i].Player.Name, Inline: aws.Bool(false),
		})
	}
	_, err := wh.c.CreateEmbeds(emb)
	return err
}

// HandleOnGameStateChanged implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:     p.Event.Event,
			Timestamp: aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "OLD STATE", Value: p.OldState, Inline: aws.Bool(false)},
				{Name: "NEW STATE", Value: p.NewState, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnGoingLive implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnGrenadeThrown implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				{Name: "WEAPON", Value: p.Weapon.Name, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnHEGrenadeDetonated implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	fields := []discord.EmbedField{
		{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
		{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
		{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
		{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
		{Name: "WEAPON", Value: p.Weapon.Name, Inline: aws.Bool(false)},
		{Name: "DAMAGE ENEMIES", Value: strconv.Itoa(p.DamageEnemies), Inline: aws.Bool(false)},
		{Name: "DAMAGE FRIENDLIES", Value: strconv.Itoa(p.DamageFriendlies), Inline: aws.Bool(false)},
	}
	emb := []discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields:      fields,
		},
	}
	for i := 0; i < len(p.Victims); i++ {
		fields = append(fields, discord.EmbedField{
			Name: fmt.Sprintf("VICTIM [%d]", i), Value: p.Victims[i].Player.Name, Inline: aws.Bool(false),
		})
	}
	_, err := wh.c.CreateEmbeds(emb)
	return err
}

// HandleOnKnifeRoundStarted implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnKnifeRoundWon implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "TEAM", Value: p.Team, Inline: aws.Bool(false)},
				{Name: "SIDE", Value: p.Side, Inline: aws.Bool(false)},
				{Name: "SWAPPED", Value: strconv.FormatBool(p.Swapped), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnLoadMatchConfigFailed implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:     p.Event.Event,
			Timestamp: aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "REASON", Value: p.Reason, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnMapPicked implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "TEAM", Value: p.Team, Inline: aws.Bool(false)},
				{Name: "MAP NAME", Value: p.MapName, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnMapResult implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "TEAM1 Score", Value: strconv.Itoa(p.Team1.Score), Inline: aws.Bool(false)},
				{Name: "TEAM2Score", Value: strconv.Itoa(p.Team1.Score), Inline: aws.Bool(false)},
				{Name: "WINNER", Value: p.Winner.Team, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnMapVetoed implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "TEAM", Value: p.Team, Inline: aws.Bool(false)},
				{Name: "MAP NAME", Value: p.MapName, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnMatchPaused implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "TEAM", Value: p.Team, Inline: aws.Bool(false)},
				{Name: "PAUSE TYPE", Value: p.PauseType, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnMatchUnpaused implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "TEAM", Value: p.Team, Inline: aws.Bool(false)},
				{Name: "PAUSE TYPE", Value: p.PauseType, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnPauseBegan implements controller.EventHandler.
func (wh *webhookEventHandler) HandleOnPauseBegan(ctx context.Context, p models.OnPauseBeganPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "TEAM", Value: p.Team, Inline: aws.Bool(false)},
				{Name: "PAUSE TYPE", Value: p.PauseType, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnMolotovDetonated implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	fields := []discord.EmbedField{
		{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
		{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
		{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
		{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
		{Name: "WEAPON", Value: p.Weapon.Name, Inline: aws.Bool(false)},
		{Name: "DAMAGE ENEMIES", Value: strconv.Itoa(p.DamageEnemies), Inline: aws.Bool(false)},
		{Name: "DAMAGE FRIENDLIES", Value: strconv.Itoa(p.DamageFriendlies), Inline: aws.Bool(false)},
	}
	emb := []discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields:      fields,
		},
	}
	for i := 0; i < len(p.Victims); i++ {
		fields = append(fields, discord.EmbedField{
			Name: fmt.Sprintf("VICTIM [%d]", i), Value: p.Victims[i].Player.Name, Inline: aws.Bool(false),
		})
	}
	_, err := wh.c.CreateEmbeds(emb)
	return err
}

// HandleOnPlayerBecameMVP implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				{Name: "REASON", Value: strconv.Itoa(p.Reason), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnPlayerConnected implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:     p.Event.Event,
			Timestamp: aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				// {Name: "IP Address", Value: p.IPAddress, Inline: aws.Bool(false)}, // You may want to hide this
			},
		},
	})
	return err
}

// HandleOnPlayerDeath implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				{Name: "WEAPON", Value: p.Weapon.Name, Inline: aws.Bool(false)},
				{Name: "BOMB", Value: strconv.FormatBool(p.Bomb), Inline: aws.Bool(false)},
				{Name: "HEADSHOT", Value: strconv.FormatBool(p.Headshot), Inline: aws.Bool(false)},
				{Name: "THRU SMOKE", Value: strconv.FormatBool(p.ThruSmoke), Inline: aws.Bool(false)},
				{Name: "PENETRATED", Value: strconv.FormatFloat(p.Penetrated, 'f', -1, 64), Inline: aws.Bool(false)},
				{Name: "ATTACKER BLIND", Value: strconv.FormatBool(p.AttackerBlind), Inline: aws.Bool(false)},
				{Name: "NO SCOPE", Value: strconv.FormatBool(p.NoScope), Inline: aws.Bool(false)},
				{Name: "SUICIDE", Value: strconv.FormatBool(p.Suicide), Inline: aws.Bool(false)},
				{Name: "FRIENDLY FIRE", Value: strconv.FormatBool(p.FriendlyFire), Inline: aws.Bool(false)},
				{Name: "ATTACKER", Value: p.Attacker.Name, Inline: aws.Bool(false)},
				{Name: "ASSIST", Value: p.Assist.Player.Name, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnPlayerDisconnected implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:     p.Event.Event,
			Timestamp: aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnPlayerSay implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	// You can GET their profile image/name by Steam API, and pretend on webhook message
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
				{Name: "COMMAND", Value: p.Command, Inline: aws.Bool(false)},
				{Name: "MESSAGE", Value: p.Message, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnPreLoadMatchConfig implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	// You can GET their profile image/name by Steam API, and pretend on webhook message
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:     p.Event.Event,
			Timestamp: aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "FILE NAME", Value: p.Filename, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnRoundEnd implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
				{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
				{Name: "REASON", Value: strconv.Itoa(p.Reason), Inline: aws.Bool(false)},
				{Name: "WINNER", Value: p.Winner.Team, Inline: aws.Bool(false)},
				{Name: "TEAM1 Score", Value: strconv.Itoa(p.Team1.Score), Inline: aws.Bool(false)},
				{Name: "TEAM2Score", Value: strconv.Itoa(p.Team2.Score), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnRoundStart implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnRoundStatsUpdated implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnSeriesInit implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "TEAM1 NAME", Value: p.Team1.Name, Inline: aws.Bool(false)},
				{Name: "TEAM2 NAME", Value: p.Team2.Name, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnSeriesResult implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "TEAM1 SERIES SCORE", Value: strconv.Itoa(p.Team1SeriesScore), Inline: aws.Bool(false)},
				{Name: "TEAM2 SERIES SCORE", Value: strconv.Itoa(p.Team2SeriesScore), Inline: aws.Bool(false)},
				{Name: "WINNER TEAM", Value: p.Winner.Team, Inline: aws.Bool(false)},
				{Name: "WINNER SIDE", Value: p.Winner.Side, Inline: aws.Bool(false)},
				{Name: "TIME UNTIL RESTORE", Value: strconv.Itoa(p.TimeUntilRestore), Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnSidePicked implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "TEAM", Value: p.Team, Inline: aws.Bool(false)},
				{Name: "MAP NAME", Value: p.MapName, Inline: aws.Bool(false)},
				{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
				{Name: "SIDE", Value: p.Side, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// HandleOnSmokeGrenadeDetonated implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	fields := []discord.EmbedField{
		{Name: "MAP", Value: strconv.Itoa(p.MapNumber), Inline: aws.Bool(false)},
		{Name: "ROUND", Value: strconv.Itoa(p.RoundNumber), Inline: aws.Bool(false)},
		{Name: "ROUND TIME", Value: strconv.Itoa(p.RoundTime), Inline: aws.Bool(false)},
		{Name: "PLAYER", Value: p.Player.Name, Inline: aws.Bool(false)},
		{Name: "WEAPON", Value: p.Weapon.Name, Inline: aws.Bool(false)},
		{Name: "EXTINGUISHED MOLOTOV", Value: strconv.FormatBool(p.ExtinguishedMolotov), Inline: aws.Bool(false)},
	}
	emb := []discord.Embed{
		{
			Title:       p.Event.Event,
			Description: p.MatchID,
			Timestamp:   aws.Time(time.Now()),
			Fields:      fields,
		},
	}
	_, err := wh.c.CreateEmbeds(emb)
	return err
}

// HandleOnTeamReadyStatusChanged implements controller.EventHandler
func (wh *webhookEventHandler) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	_, err := wh.c.CreateEmbeds([]discord.Embed{
		{
			Title:     p.Event.Event,
			Timestamp: aws.Time(time.Now()),
			Fields: []discord.EmbedField{
				{Name: "TEAM", Value: *p.Team, Inline: aws.Bool(false)},
				{Name: "READY", Value: strconv.FormatBool(p.Ready), Inline: aws.Bool(false)},
				{Name: "GAME STATE", Value: p.GameState, Inline: aws.Bool(false)},
			},
		},
	})
	return err
}

// NewEventHandler Get Logger pointer
func NewEventHandler(webhookID uint64, webhookToken string, url string, auth string) controller.EventHandler {
	return &webhookEventHandler{
		eventURL: url, // can be empty
		auth:     auth,
		c:        webhook.New(snowflake.ID(webhookID), webhookToken),
	}
}
