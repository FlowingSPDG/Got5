package discord

import (
	"context"
	"fmt"

	"github.com/FlowingSPDG/Got5/models"
)

// UpdateMatch implements controller.Controller
func (*discord) UpdateMatch(ctx context.Context, mid string, m models.Match) error {
	return nil
}

// RegisterDemoFile implements controller.Controller
func (d *discord) RegisterDemoFile(ctx context.Context, mid string, filename string, b []byte) error {
	return nil
}

// RegisterMatch implements controller.Controller
func (d *discord) RegisterMatch(ctx context.Context, m models.Match) (models.Match, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if ms, ok := d.matches[m.MatchID]; ok {
		ms.match = m
		return m, nil
	}
	return models.Match{}, fmt.Errorf("Match not found")
}

// GetMatch implements controller.Controller
func (d *discord) GetMatch(ctx context.Context, mid string) (models.G5Match, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m, ok := d.matches[mid]
	if !ok {
		return nil, fmt.Errorf("Not found")
	}
	return m.match, nil
}

// HandleOnBackupRestore implements controller.Controller
func (d *discord) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	return nil
}

// HandleOnBombDefused implements controller.Controller
func (d *discord) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	return nil
}

// HandleOnBombExploded implements controller.Controller
func (d *discord) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	return nil
}

// HandleOnBombPlanted implements controller.Controller
func (d *discord) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	return nil
}

// HandleOnDecoyStarted implements controller.Controller
func (d *discord) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	return nil
}

// HandleOnDemoFinished implements controller.Controller
func (d *discord) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	return nil
}

// HandleOnDemoUploadEnded implements controller.Controller
func (d *discord) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	return nil
}

// HandleOnEvent implements controller.Controller
func (d *discord) HandleOnEvent(ctx context.Context, p models.OnEventPayload) error {
	return nil
}

// HandleOnFlashbangDetonated implements controller.Controller
func (d *discord) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	return nil
}

// HandleOnGameStateChanged implements controller.Controller
func (d *discord) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	return nil
}

// HandleOnGoingLive implements controller.Controller
func (d *discord) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	// 一度Interactionに返事をするとそれ以上の返信はできず、メッセージを編集するかFollow Up Meesageのみとなる
	// しかしInteraction token自体が15分しか持続しないため、それ以降は個別にメッセージを送信する必要がある
	d.mu.RLock()
	defer d.mu.RUnlock()
	m := d.matches[p.Matchid]
	msg := fmt.Sprintf("🔫ゲーム %s がまもなく開始します！\n[G]ood [L]uck [H]ave [F]un!", m.match.MatchTitle)
	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	return nil
}

// HandleOnGrenadeThrown implements controller.Controller
func (d *discord) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	return nil
}

// HandleOnHEGrenadeDetonated implements controller.Controller
func (d *discord) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnKnifeRoundStarted implements controller.Controller
func (d *discord) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m := d.matches[p.Matchid]
	msg := fmt.Sprintf("🔪ゲーム ``%s``\nナイフラウンドがまもなく開始します！🔪🔪", m.match.MatchTitle)

	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	return nil
}

// HandleOnKnifeRoundWon implements controller.Controller
func (d *discord) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m := d.matches[p.Matchid]
	msg := fmt.Sprintf("🔪ゲーム ``%s``\nチーム%sがナイフラウンドに勝利しました！🔪🔪", m.match.MatchTitle, p.Team)
	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	return nil
}

// HandleOnLoadMatchConfigFailed implements controller.Controller
func (d *discord) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	return nil
}

// HandleOnMapPicked implements controller.Controller
func (d *discord) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	return nil
}

// HandleOnMapResult implements controller.Controller
// どうせBO1だしマップ結果も出さなくていいんじゃないか...?
func (d *discord) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	return nil
}

// HandleOnMapVetoed implements controller.Controller
func (d *discord) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	return nil
}

// HandleOnMatchPaused implements controller.Controller
func (d *discord) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	return nil
}

// HandleOnMatchUnpaused implements controller.Controller
func (d *discord) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	return nil
}

// HandleOnMolotovDetonated implements controller.Controller
func (d *discord) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	return nil
}

// HandleOnPlayerBecameMVP implements controller.Controller
func (d *discord) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	return nil
}

// HandleOnPlayerConnected implements controller.Controller
func (d *discord) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	return nil
}

// HandleOnPlayerDeath implements controller.Controller
func (d *discord) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	return nil
}

// HandleOnPlayerDisconnected implements controller.Controller
func (d *discord) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	return nil
}

// HandleOnPlayerSay implements controller.Controller
func (d *discord) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	return nil
}

// HandleOnPreLoadMatchConfig implements controller.Controller
func (d *discord) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	return nil
}

// HandleOnRoundEnd implements controller.Controller
func (d *discord) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	return nil
}

// HandleOnRoundStart implements controller.Controller
func (d *discord) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	return nil
}

// HandleOnRoundStatsUpdated implements controller.Controller
func (d *discord) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	return nil
}

// HandleOnSeriesInit implements controller.Controller
func (d *discord) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	return nil
}

// HandleOnSeriesResult implements controller.Controller
func (d *discord) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	m, ok := d.matches[p.Matchid]
	if !ok {
		return fmt.Errorf("Invalid Match ID")
	}
	msg := fmt.Sprintf("🎉チーム %s がゲームに勝利しました！GGWP！\n最終スコア: %s[%d]:[%d]%s", p.Winner.Team, m.match.Team1.Name, p.Team1SeriesScore, p.Team2SeriesScore, m.match.Team2.Name)
	if _, err := d.s.ChannelMessageSend(m.interaction.ChannelID, msg); err != nil {
		return err
	}
	// mapの容量が無限に増えちゃうのでdeleteをかける
	delete(d.matches, p.Matchid)
	return nil
}

// HandleOnSidePicked implements controller.Controller
func (d *discord) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	return nil
}

// HandleOnSmokeGrenadeDetonated implements controller.Controller
func (d *discord) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnTeamReadyStatusChanged implements controller.Controller
func (d *discord) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	return nil
}
