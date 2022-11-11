package controller

import "github.com/FlowingSPDG/Got5/models"

// Controller Controller interface operates write operation for e.g. Databases.
type Controller interface {
	HandleOnEvent(p models.OnEventPayload) error
	HandleOnGameStateChanged(p models.OnGameStateChangedPayload) error
	HandleOnPreLoadMatchConfig(p models.OnPreLoadMatchConfigPayload) error
	HandleOnLoadMatchConfigFailed(p models.OnLoadMatchConfigFailedPayload) error
	HandleOnSeriesInit(p models.OnSeriesInitPayload) error
	HandleOnMapResult(p models.OnMapResultPayload) error
	HandleOnSeriesResult(p models.OnSeriesResultPayload) error
	HandleOnSidePicked(p models.OnSidePickedPayload) error
	HandleOnMapPicked(p models.OnMapPickedPayload) error
	HandleOnMapVetoed(p models.OnMapVetoedPayload) error
	HandleOnBackupRestore(p models.OnBackupRestorePayload) error
	HandleOnDemoFinished(p models.OnDemoFinishedPayload) error
	HandleOnDemoUploadEnded(p models.OnDemoUploadEndedPayload) error
	HandleOnMatchPaused(p models.OnMatchPausedPayload) error
	HandleOnMatchUnpaused(p models.OnMatchUnpausedPayload) error
	HandleOnKnifeRoundStarted(p models.OnKnifeRoundStartedPayload) error
	HandleOnKnifeRoundWon(p models.OnKnifeRoundWonPayload) error
	HandleOnTeamReadyStatusChanged(p models.OnTeamReadyStatusChangedPayload) error
	HandleOnGoingLive(p models.OnGoingLivePayload) error
	HandleOnRoundStart(p models.OnRoundStartPayload) error
	HandleOnRoundEnd(p models.OnRoundEndPayload) error
	HandleOnRoundStatsUpdated(p models.OnRoundStatsUpdatedPayload) error
	HandleOnPlayerBecameMVP(p models.OnPlayerBecameMVPPayload) error
	HandleOnGrenadeThrown(p models.OnGrenadeThrownPayload) error
	HandleOnPlayerDeath(p models.OnPlayerDeathPayload) error
	HandleOnHEGrenadeDetonated(p models.OnHEGrenadeDetonatedPayload) error
	HandleOnMolotovDetonated(p models.OnMolotovDetonatedPayload) error
	HandleOnFlashbangDetonated(p models.OnFlashbangDetonatedPayload) error
	HandleOnSmokeGrenadeDetonated(p models.OnSmokeGrenadeDetonatedPayload) error
	HandleOnDecoyStarted(p models.OnDecoyStartedPayload) error
	HandleOnBombPlanted(p models.OnBombPlantedPayload) error
	HandleOnBombDefused(p models.OnBombDefusedPayload) error
	HandleOnBombExploded(p models.OnBombExplodedPayload) error
	HandleOnPlayerConnected(p models.OnPlayerConnectedPayload) error
	HandleOnPlayerDisconnected(p models.OnPlayerDisconnectedPayload) error
	HandleOnPlayerSay(p models.OnPlayerSayPayload) error
}
