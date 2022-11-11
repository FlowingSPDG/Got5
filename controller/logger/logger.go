package logger

import (
	"log"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

// logger is simply prints what happend.
type logger struct{}

// HandleOnEvent implements controller.Controller
func (*logger) HandleOnEvent(p models.OnEventPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnBackupRestore implements controller.Controller
func (*logger) HandleOnBackupRestore(p models.OnBackupRestorePayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnBombDefused implements controller.Controller
func (*logger) HandleOnBombDefused(p models.OnBombDefusedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnBombExploded implements controller.Controller
func (*logger) HandleOnBombExploded(p models.OnBombExplodedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnBombPlanted implements controller.Controller
func (*logger) HandleOnBombPlanted(p models.OnBombPlantedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnDecoyStarted implements controller.Controller
func (*logger) HandleOnDecoyStarted(p models.OnDecoyStartedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnDemoFinished implements controller.Controller
func (*logger) HandleOnDemoFinished(p models.OnDemoFinishedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnDemoUploadEnded implements controller.Controller
func (*logger) HandleOnDemoUploadEnded(p models.OnDemoUploadEndedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnFlashbangDetonated implements controller.Controller
func (*logger) HandleOnFlashbangDetonated(p models.OnFlashbangDetonatedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnGameStateChanged implements controller.Controller
func (*logger) HandleOnGameStateChanged(p models.OnGameStateChangedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnGoingLive implements controller.Controller
func (*logger) HandleOnGoingLive(p models.OnGoingLivePayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnGrenadeThrown implements controller.Controller
func (*logger) HandleOnGrenadeThrown(p models.OnGrenadeThrownPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnHEGrenadeDetonated implements controller.Controller
func (*logger) HandleOnHEGrenadeDetonated(p models.OnHEGrenadeDetonatedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnKnifeRoundStarted implements controller.Controller
func (*logger) HandleOnKnifeRoundStarted(p models.OnKnifeRoundStartedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnKnifeRoundWon implements controller.Controller
func (*logger) HandleOnKnifeRoundWon(p models.OnKnifeRoundWonPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnLoadMatchConfigFailed implements controller.Controller
func (*logger) HandleOnLoadMatchConfigFailed(p models.OnLoadMatchConfigFailedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnMapPicked implements controller.Controller
func (*logger) HandleOnMapPicked(p models.OnMapPickedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnMapResult implements controller.Controller
func (*logger) HandleOnMapResult(p models.OnMapResultPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnMapVetoed implements controller.Controller
func (*logger) HandleOnMapVetoed(p models.OnMapVetoedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnMatchPaused implements controller.Controller
func (*logger) HandleOnMatchPaused(p models.OnMatchPausedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnMatchUnpaused implements controller.Controller
func (*logger) HandleOnMatchUnpaused(p models.OnMatchUnpausedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnMolotovDetonated implements controller.Controller
func (*logger) HandleOnMolotovDetonated(p models.OnMolotovDetonatedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnPlayerBecameMVP implements controller.Controller
func (*logger) HandleOnPlayerBecameMVP(p models.OnPlayerBecameMVPPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnPlayerConnected implements controller.Controller
func (*logger) HandleOnPlayerConnected(p models.OnPlayerConnectedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnPlayerDeath implements controller.Controller
func (*logger) HandleOnPlayerDeath(p models.OnPlayerDeathPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnPlayerDisconnected implements controller.Controller
func (*logger) HandleOnPlayerDisconnected(p models.OnPlayerDisconnectedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnPlayerSay implements controller.Controller
func (*logger) HandleOnPlayerSay(p models.OnPlayerSayPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnPreLoadMatchConfig implements controller.Controller
func (*logger) HandleOnPreLoadMatchConfig(p models.OnPreLoadMatchConfigPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnRoundEnd implements controller.Controller
func (*logger) HandleOnRoundEnd(p models.OnRoundEndPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnRoundStart implements controller.Controller
func (*logger) HandleOnRoundStart(p models.OnRoundStartPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnRoundStatsUpdated implements controller.Controller
func (*logger) HandleOnRoundStatsUpdated(p models.OnRoundStatsUpdatedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnSeriesInit implements controller.Controller
func (*logger) HandleOnSeriesInit(p models.OnSeriesInitPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnSeriesResult implements controller.Controller
func (*logger) HandleOnSeriesResult(p models.OnSeriesResultPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnSidePicked implements controller.Controller
func (*logger) HandleOnSidePicked(p models.OnSidePickedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnSmokeGrenadeDetonated implements controller.Controller
func (*logger) HandleOnSmokeGrenadeDetonated(p models.OnSmokeGrenadeDetonatedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// HandleOnTeamReadyStatusChanged implements controller.Controller
func (*logger) HandleOnTeamReadyStatusChanged(p models.OnTeamReadyStatusChangedPayload) error {
	log.Println("Payload:", p)
	return nil
}

// NewLoggerController Get Logger pointer
func NewLoggerController() controller.Controller {
	return &logger{}
}
