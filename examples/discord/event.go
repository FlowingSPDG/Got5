package discord

import (
	"context"
	"fmt"

	got5 "github.com/FlowingSPDG/Got5"
)

// CheckEventAuth implements got5.EventHandler
func (*Discord) CheckEventAuth(ctx context.Context, mid string, reqAuth string) error {
	return nil
}

// UpdateMatch implements got5.EventHandler
func (*Discord) UpdateMatch(ctx context.Context, mid string, m got5.Match) error {
	return nil
}

// RegisterMatch implements got5.EventHandler
func (d *Discord) RegisterMatch(ctx context.Context, m got5.Match) (got5.Match, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if ms, ok := d.matches[m.MatchID]; ok {
		ms.match = m.ToG5Format()
		return m, nil
	}
	return got5.Match{}, fmt.Errorf("Match not found")
}

// Load implements got5.MatchLoader
func (d *Discord) Load(ctx context.Context, mid string) (got5.G5Match, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m, ok := d.matches[mid]
	if !ok {
		return nil, fmt.Errorf("Not found")
	}
	return m.match, nil
}

// HandleOnBackupRestore implements got5.EventHandler
func (d *Discord) HandleOnBackupRestore(ctx context.Context, p got5.OnBackupRestorePayload) error {
	return nil
}

// HandleOnBombDefused implements got5.EventHandler
func (d *Discord) HandleOnBombDefused(ctx context.Context, p got5.OnBombDefusedPayload) error {
	return nil
}

// HandleOnBombExploded implements got5.EventHandler
func (d *Discord) HandleOnBombExploded(ctx context.Context, p got5.OnBombExplodedPayload) error {
	return nil
}

// HandleOnBombPlanted implements got5.EventHandler
func (d *Discord) HandleOnBombPlanted(ctx context.Context, p got5.OnBombPlantedPayload) error {
	return nil
}

// HandleOnDecoyStarted implements got5.EventHandler
func (d *Discord) HandleOnDecoyStarted(ctx context.Context, p got5.OnDecoyStartedPayload) error {
	return nil
}

// HandleOnDemoFinished implements got5.EventHandler
func (d *Discord) HandleOnDemoFinished(ctx context.Context, p got5.OnDemoFinishedPayload) error {
	return nil
}

// HandleOnDemoUploadEnded implements got5.EventHandler
func (d *Discord) HandleOnDemoUploadEnded(ctx context.Context, p got5.OnDemoUploadEndedPayload) error {
	return nil
}

// HandleOnEvent implements got5.EventHandler
func (d *Discord) HandleOnEvent(ctx context.Context, p got5.OnEventPayload) error {
	return nil
}

// HandleOnFlashbangDetonated implements got5.EventHandler
func (d *Discord) HandleOnFlashbangDetonated(ctx context.Context, p got5.OnFlashbangDetonatedPayload) error {
	return nil
}

// HandleOnGameStateChanged implements got5.EventHandler
func (d *Discord) HandleOnGameStateChanged(ctx context.Context, p got5.OnGameStateChangedPayload) error {
	return nil
}

// HandleOnGoingLive implements got5.EventHandler
func (d *Discord) HandleOnGoingLive(ctx context.Context, p got5.OnGoingLivePayload) error {
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

// HandleOnGrenadeThrown implements got5.EventHandler
func (d *Discord) HandleOnGrenadeThrown(ctx context.Context, p got5.OnGrenadeThrownPayload) error {
	return nil
}

// HandleOnHEGrenadeDetonated implements got5.EventHandler
func (d *Discord) HandleOnHEGrenadeDetonated(ctx context.Context, p got5.OnHEGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnKnifeRoundStarted implements got5.EventHandler
func (d *Discord) HandleOnKnifeRoundStarted(ctx context.Context, p got5.OnKnifeRoundStartedPayload) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m := d.matches[p.MatchID]
	msg := fmt.Sprintf("🔪ゲーム ``%s``\nナイフラウンドがまもなく開始します！🔪🔪", m.match.MatchTitle)

	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	return nil
}

// HandleOnKnifeRoundWon implements got5.EventHandler
func (d *Discord) HandleOnKnifeRoundWon(ctx context.Context, p got5.OnKnifeRoundWonPayload) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m := d.matches[p.MatchID]
	msg := fmt.Sprintf("🔪ゲーム ``%s``\nチーム%sがナイフラウンドに勝利しました！🔪🔪", m.match.MatchTitle, p.Team)
	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	return nil
}

// HandleOnLoadMatchConfigFailed implements got5.EventHandler
func (d *Discord) HandleOnLoadMatchConfigFailed(ctx context.Context, p got5.OnLoadMatchConfigFailedPayload) error {
	return nil
}

// HandleOnMapPicked implements got5.EventHandler
func (d *Discord) HandleOnMapPicked(ctx context.Context, p got5.OnMapPickedPayload) error {
	return nil
}

// HandleOnMapResult implements got5.EventHandler
// どうせBO1だしマップ結果も出さなくていいんじゃないか...?
func (d *Discord) HandleOnMapResult(ctx context.Context, p got5.OnMapResultPayload) error {
	return nil
}

// HandleOnMapVetoed implements got5.EventHandler
func (d *Discord) HandleOnMapVetoed(ctx context.Context, p got5.OnMapVetoedPayload) error {
	return nil
}

// HandleOnMatchPaused implements got5.EventHandler
func (d *Discord) HandleOnMatchPaused(ctx context.Context, p got5.OnMatchPausedPayload) error {
	return nil
}

// HandleOnMatchUnpaused implements got5.EventHandler
func (d *Discord) HandleOnMatchUnpaused(ctx context.Context, p got5.OnMatchUnpausedPayload) error {
	return nil
}

// HandleOnPauseBegan implements got5.EventHandler.
func (d *Discord) HandleOnPauseBegan(ctx context.Context, p got5.OnPauseBeganPayload) error {
	return nil
}

// HandleOnMolotovDetonated implements got5.EventHandler
func (d *Discord) HandleOnMolotovDetonated(ctx context.Context, p got5.OnMolotovDetonatedPayload) error {
	return nil
}

// HandleOnPlayerBecameMVP implements got5.EventHandler
func (d *Discord) HandleOnPlayerBecameMVP(ctx context.Context, p got5.OnPlayerBecameMVPPayload) error {
	return nil
}

// HandleOnPlayerConnected implements got5.EventHandler
func (d *Discord) HandleOnPlayerConnected(ctx context.Context, p got5.OnPlayerConnectedPayload) error {
	return nil
}

// HandleOnPlayerDeath implements got5.EventHandler
func (d *Discord) HandleOnPlayerDeath(ctx context.Context, p got5.OnPlayerDeathPayload) error {
	return nil
}

// HandleOnPlayerDisconnected implements got5.EventHandler
func (d *Discord) HandleOnPlayerDisconnected(ctx context.Context, p got5.OnPlayerDisconnectedPayload) error {
	return nil
}

// HandleOnPlayerSay implements got5.EventHandler
func (d *Discord) HandleOnPlayerSay(ctx context.Context, p got5.OnPlayerSayPayload) error {
	return nil
}

// HandleOnPreLoadMatchConfig implements got5.EventHandler
func (d *Discord) HandleOnPreLoadMatchConfig(ctx context.Context, p got5.OnPreLoadMatchConfigPayload) error {
	return nil
}

// HandleOnRoundEnd implements got5.EventHandler
func (d *Discord) HandleOnRoundEnd(ctx context.Context, p got5.OnRoundEndPayload) error {
	return nil
}

// HandleOnRoundStart implements got5.EventHandler
func (d *Discord) HandleOnRoundStart(ctx context.Context, p got5.OnRoundStartPayload) error {
	return nil
}

// HandleOnRoundStatsUpdated implements got5.EventHandler
func (d *Discord) HandleOnRoundStatsUpdated(ctx context.Context, p got5.OnRoundStatsUpdatedPayload) error {
	return nil
}

// HandleOnSeriesInit implements got5.EventHandler
func (d *Discord) HandleOnSeriesInit(ctx context.Context, p got5.OnSeriesInitPayload) error {
	return nil
}

// HandleOnSeriesResult implements got5.EventHandler
func (d *Discord) HandleOnSeriesResult(ctx context.Context, p got5.OnSeriesResultPayload) error {
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

// HandleOnSidePicked implements got5.EventHandler
func (d *Discord) HandleOnSidePicked(ctx context.Context, p got5.OnSidePickedPayload) error {
	return nil
}

// HandleOnSmokeGrenadeDetonated implements got5.EventHandler
func (d *Discord) HandleOnSmokeGrenadeDetonated(ctx context.Context, p got5.OnSmokeGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnTeamReadyStatusChanged implements got5.EventHandler
func (d *Discord) HandleOnTeamReadyStatusChanged(ctx context.Context, p got5.OnTeamReadyStatusChangedPayload) error {
	return nil
}
