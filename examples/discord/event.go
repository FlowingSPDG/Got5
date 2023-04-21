package discord

import (
	"context"
	"fmt"

	"github.com/FlowingSPDG/Got5/models"
)

// CheckEventAuth implements controller.EventHandler
func (*Discord) CheckEventAuth(ctx context.Context, mid string, reqAuth string) error {
	return nil
}

// UpdateMatch implements controller.EventHandler
func (*Discord) UpdateMatch(ctx context.Context, mid string, m models.Match) error {
	return nil
}

// RegisterMatch implements controller.EventHandler
func (d *Discord) RegisterMatch(ctx context.Context, m models.Match) (models.Match, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if ms, ok := d.matches[m.MatchID]; ok {
		ms.match = m.ToG5Format()
		return m, nil
	}
	return models.Match{}, fmt.Errorf("Match not found")
}

// Load implements controller.MatchLoader
func (d *Discord) Load(ctx context.Context, mid string) (models.G5Match, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m, ok := d.matches[mid]
	if !ok {
		return nil, fmt.Errorf("Not found")
	}
	return m.match, nil
}

// HandleOnBackupRestore implements controller.EventHandler
func (d *Discord) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	return nil
}

// HandleOnBombDefused implements controller.EventHandler
func (d *Discord) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	return nil
}

// HandleOnBombExploded implements controller.EventHandler
func (d *Discord) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	return nil
}

// HandleOnBombPlanted implements controller.EventHandler
func (d *Discord) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	return nil
}

// HandleOnDecoyStarted implements controller.EventHandler
func (d *Discord) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	return nil
}

// HandleOnDemoFinished implements controller.EventHandler
func (d *Discord) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	return nil
}

// HandleOnDemoUploadEnded implements controller.EventHandler
func (d *Discord) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	return nil
}

// HandleOnEvent implements controller.EventHandler
func (d *Discord) HandleOnEvent(ctx context.Context, p models.OnEventPayload) error {
	return nil
}

// HandleOnFlashbangDetonated implements controller.EventHandler
func (d *Discord) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	return nil
}

// HandleOnGameStateChanged implements controller.EventHandler
func (d *Discord) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	return nil
}

// HandleOnGoingLive implements controller.EventHandler
func (d *Discord) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	// 一度Interactionに返事をするとそれ以上の返信はできず、メッセージを編集するかFollow Up Meesageのみとなる
	// しかしInteraction token自体が15分しか持続しないため、それ以降は個別にメッセージを送信する必要がある
	d.mu.RLock()
	defer d.mu.RUnlock()
	m := d.matches[p.MatchID]
	msg := fmt.Sprintf("🔫ゲーム %s がまもなく開始します！\n[G]ood [L]uck [H]ave [F]un!", m.match.MatchTitle)
	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	return nil
}

// HandleOnGrenadeThrown implements controller.EventHandler
func (d *Discord) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	return nil
}

// HandleOnHEGrenadeDetonated implements controller.EventHandler
func (d *Discord) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnKnifeRoundStarted implements controller.EventHandler
func (d *Discord) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m := d.matches[p.MatchID]
	msg := fmt.Sprintf("🔪ゲーム ``%s``\nナイフラウンドがまもなく開始します！🔪🔪", m.match.MatchTitle)

	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	return nil
}

// HandleOnKnifeRoundWon implements controller.EventHandler
func (d *Discord) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m := d.matches[p.MatchID]
	msg := fmt.Sprintf("🔪ゲーム ``%s``\nチーム%sがナイフラウンドに勝利しました！🔪🔪", m.match.MatchTitle, p.Team)
	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	return nil
}

// HandleOnLoadMatchConfigFailed implements controller.EventHandler
func (d *Discord) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	return nil
}

// HandleOnMapPicked implements controller.EventHandler
func (d *Discord) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	return nil
}

// HandleOnMapResult implements controller.EventHandler
// どうせBO1だしマップ結果も出さなくていいんじゃないか...?
func (d *Discord) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	return nil
}

// HandleOnMapVetoed implements controller.EventHandler
func (d *Discord) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	return nil
}

// HandleOnMatchPaused implements controller.EventHandler
func (d *Discord) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	return nil
}

// HandleOnMatchUnpaused implements controller.EventHandler
func (d *Discord) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	return nil
}

// HandleOnMolotovDetonated implements controller.EventHandler
func (d *Discord) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	return nil
}

// HandleOnPlayerBecameMVP implements controller.EventHandler
func (d *Discord) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	return nil
}

// HandleOnPlayerConnected implements controller.EventHandler
func (d *Discord) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	return nil
}

// HandleOnPlayerDeath implements controller.EventHandler
func (d *Discord) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	return nil
}

// HandleOnPlayerDisconnected implements controller.EventHandler
func (d *Discord) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	return nil
}

// HandleOnPlayerSay implements controller.EventHandler
func (d *Discord) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	return nil
}

// HandleOnPreLoadMatchConfig implements controller.EventHandler
func (d *Discord) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	return nil
}

// HandleOnRoundEnd implements controller.EventHandler
func (d *Discord) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	return nil
}

// HandleOnRoundStart implements controller.EventHandler
func (d *Discord) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	return nil
}

// HandleOnRoundStatsUpdated implements controller.EventHandler
func (d *Discord) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	return nil
}

// HandleOnSeriesInit implements controller.EventHandler
func (d *Discord) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	return nil
}

// HandleOnSeriesResult implements controller.EventHandler
func (d *Discord) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m, ok := d.matches[p.MatchID]
	if !ok {
		return fmt.Errorf("Invalid Match ID")
	}
	msg := fmt.Sprintf("🎉チーム %s がゲームに勝利しました！GGWP！\n最終スコア: %s[%d]:[%d]%s", p.Winner.Team, m.match.Team1.Name, p.Team1SeriesScore, p.Team2SeriesScore, m.match.Team2.Name)
	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	// mapの容量が無限に増えちゃうのでdeleteをかける
	delete(d.matches, p.MatchID)
	return nil
}

// HandleOnSidePicked implements controller.EventHandler
func (d *Discord) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	return nil
}

// HandleOnSmokeGrenadeDetonated implements controller.EventHandler
func (d *Discord) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnTeamReadyStatusChanged implements controller.EventHandler
func (d *Discord) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	return nil
}
