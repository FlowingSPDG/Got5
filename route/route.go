package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
)

func SetupAllGet5Handlers(ctrl controller.Controller, r fiber.Router, bucket string) error {
	if err := SetupMatchLoadHandler(ctrl, r); err != nil {
		return err
	}
	if err := SetupEventHandlers(ctrl, r); err != nil {
		return err
	}
	if err := SetupDemoUploadHandler(ctrl, r, bucket); err != nil {
		return err
	}
	return nil
}

// SetupDemoUploadHandler Setup get5 upload demo handler
func SetupDemoUploadHandler(ctrl controller.Controller, r fiber.Router, bucket string) error {
	r.Post("/match/:matchID/demo", DemoUploadHandler(ctrl, bucket))
	return nil
}

// SetupMatchLoadHandlers Setup get5 loadmatch handler
func SetupMatchLoadHandler(ctrl controller.Controller, r fiber.Router) error {
	r.Get("/match/:matchID", LoadMatchHandler(ctrl))
	return nil
}

// SetupRouter Setup get5 handlers to specified fiber.Router
func SetupEventHandlers(ctrl controller.Controller, r fiber.Router) error {
	ev := r.Group("/event") // /event
	ev.Post("/Get5_OnEvent", OnEventHandler(ctrl))
	ev.Post("/Get5_OnGameStateChanged", OnGameStateChangedHandler(ctrl))
	ev.Post("/Get5_OnPlayerConnected", OnPlayerConnectedHandler(ctrl))
	ev.Post("/Get5_OnPlayerDisconnected", OnPlayerDisconnectedHandler(ctrl))
	ev.Post("/Get5_OnPreLoadMatchConfig", OnPreLoadMatchConfigHandler(ctrl))
	ev.Post("/Get5_OnLoadMatchConfigFailed", OnLoadMatchConfigFailedHandler(ctrl))
	ev.Post("/Get5_OnSeriesInit", OnSeriesInitHandler(ctrl))
	ev.Post("/Get5_OnMatchPaused", OnMatchPausedHandler(ctrl))
	ev.Post("/Get5_OnMatchUnpaused", OnMatchUnpausedHandler(ctrl))
	ev.Post("/Get5_OnMapResult", OnMapResultHandler(ctrl))
	ev.Post("/Get5_OnSeriesResult", OnSeriesResultHandler(ctrl))
	ev.Post("/Get5_OnKnifeRoundStarted", OnKnifeRoundStartedHandler(ctrl))
	ev.Post("/Get5_OnKnifeRoundWon", OnKnifeRoundWonHandler(ctrl))
	ev.Post("/Get5_OnSidePicked", OnSidePickedHandler(ctrl))
	ev.Post("/Get5_OnMapPicked", OnMapPickedHandler(ctrl))
	ev.Post("/Get5_OnMapVetoed", OnMapVetoedHandler(ctrl))
	ev.Post("/Get5_OnTeamReadyStatusChanged", OnTeamReadyStatusChanngeHandler(ctrl))
	ev.Post("/Get5_OnGoingLive", OnGoingLiveHandler(ctrl))
	ev.Post("/Get5_OnRoundStart", OnRoundStartHandler(ctrl))
	ev.Post("/Get5_OnRoundEnd", OnRoundEndHandler(ctrl))
	ev.Post("/Get5_OnRoundStatsUpdated", OnRoundStatusUpdatedHandler(ctrl))
	ev.Post("/Get5_OnBackupRestore", OnBackupRestoreHandler(ctrl))
	ev.Post("/Get5_OnDemoFinished", OnDemoFinishedHandler(ctrl))
	ev.Post("/Get5_OnDemoUploadEnded", OnDemoUploadEndedHandler(ctrl))
	ev.Post("/Get5_OnPlayerBecameMVP", OnPlayerBecameMVPHandler(ctrl))
	ev.Post("/Get5_OnGrenadeThrown", OnGrenadeThrownHandler(ctrl))
	ev.Post("/Get5_OnPlayerDeath", OnPlayerDeathHandler(ctrl))
	ev.Post("/Get5_OnPlayerSay", OnPlayerSayHandler(ctrl))
	ev.Post("/Get5_OnHEGrenadeDetonated", OnHEGrenadeDetonatedHandler(ctrl))
	ev.Post("/Get5_OnMolotovDetonated", OnMolotovDetonatedHandler(ctrl))
	ev.Post("/Get5_OnFlashbangDetonated", OnFlashbangDetonatedHandler(ctrl))
	ev.Post("/Get5_OnSmokeGrenadeDetonated", OnSmokeGrenadeDetonatedHandler(ctrl))
	ev.Post("/Get5_OnDecoyStarted", OnDecoyStartedHandler(ctrl))
	ev.Post("/Get5_OnBombPlanted", OnBombPlantedHandler(ctrl))
	ev.Post("/Get5_OnBombDefused", OnBombDefusedHandler(ctrl))
	ev.Post("/Get5_OnBombExploded", OnBombExpodedHandler(ctrl))
	return nil
}
