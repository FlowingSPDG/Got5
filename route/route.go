package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

// SetupRouter Setup get5 handlers to specified fiber.Router
func SetupRouter(ctrl controller.Controller, r fiber.Router) error {
	r.Post("/Get5_OnEvent", OnEventHandler(ctrl))
	r.Post("/Get5_OnGameStateChanged", OnGameStateChangedHandler(ctrl))
	r.Post("/Get5_OnPlayerConnected", OnPlayerConnectedHandler(ctrl))
	r.Post("/Get5_OnPlayerDisconnected", OnPlayerDisconnectedHandler(ctrl))
	r.Post("/Get5_OnPreLoadMatchConfig", OnPreLoadMatchConfigHandler(ctrl))
	r.Post("/Get5_OnLoadMatchConfigFailed", OnLoadMatchConfigFailedHandler(ctrl))
	r.Post("/Get5_OnSeriesInit", OnSeriesInitHandler(ctrl))
	r.Post("/Get5_OnMatchPaused", OnMatchPausedHandler(ctrl))
	r.Post("/Get5_OnMatchUnpaused", OnMatchUnpausedHandler(ctrl))
	r.Post("/Get5_OnMapResult", OnMapResultHandler(ctrl))
	r.Post("/Get5_OnSeriesResult", OnSeriesResultHandler(ctrl))
	r.Post("/Get5_OnKnifeRoundStarted", OnKnifeRoundStartedHandler(ctrl))
	r.Post("/Get5_OnKnifeRoundWon", OnKnifeRoundWonHandler(ctrl))
	r.Post("/Get5_OnSidePicked", OnSidePickedHandler(ctrl))
	r.Post("/Get5_OnMapPicked", OnMapPickedHandler(ctrl))
	r.Post("/Get5_OnMapVetoed", OnMapVetoedHandler(ctrl))
	r.Post("/Get5_OnTeamReadyStatusChanged", OnTeamReadyStatusChanngeHandler(ctrl))
	r.Post("/Get5_OnGoingLive", OnGoingLiveHandler(ctrl))
	r.Post("/Get5_OnRoundStart", OnRoundStartHandler(ctrl))
	r.Post("/Get5_OnRoundEnd", OnRoundEndHandler(ctrl))
	r.Post("/Get5_OnRoundStatsUpdated", OnRoundStatusUpdatedHandler(ctrl))
	r.Post("/Get5_OnBackupRestore", OnBackupRestoreHandler(ctrl))
	r.Post("/Get5_OnDemoFinished", OnDemoFinishedHandler(ctrl))
	r.Post("/Get5_OnDemoUploadEnded", OnDemoUploadEndedHandler(ctrl))
	r.Post("/Get5_OnPlayerBecameMVP", OnPlayerBecameMVPHandler(ctrl))
	r.Post("/Get5_OnGrenadeThrown", OnGrenadeThrownHandler(ctrl))
	r.Post("/Get5_OnPlayerDeath", OnPlayerDeathHandler(ctrl))
	r.Post("/Get5_OnPlayerSay", OnPlayerSayHandler(ctrl))
	r.Post("/Get5_OnHEGrenadeDetonated", OnHEGrenadeDetonatedHandler(ctrl))
	r.Post("/Get5_OnMolotovDetonated", OnMolotovDetonatedHandler(ctrl))
	r.Post("/Get5_OnFlashbangDetonated", OnFlashbangDetonatedHandler(ctrl))
	r.Post("/Get5_OnSmokeGrenadeDetonated", OnSmokeGrenadeDetonatedHandler(ctrl))
	r.Post("/Get5_OnDecoyStarted", OnDecoyStartedHandler(ctrl))
	r.Post("/Get5_OnBombPlanted", OnBombPlantedHandler(ctrl))
	r.Post("/Get5_OnBombDefused", OnBombDefusedHandler(ctrl))
	r.Post("/Get5_OnBombExploded", OnBombExpodedHandler(ctrl))
	return nil
}

// OnEventHandler POST on /Get5_OnEvent
func OnEventHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnEventPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnEvent(p)
	})
}

// OnGameStateChangedHandler
func OnGameStateChangedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnGameStateChangedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnGameStateChanged(p)
	})
}

func OnPreLoadMatchConfigHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnPreLoadMatchConfigPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnPreLoadMatchConfig(p)
	})
}

func OnLoadMatchConfigFailedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnLoadMatchConfigFailedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnLoadMatchConfigFailed(p)
	})
}

func OnSeriesInitHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnSeriesInitPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnSeriesInit(p)
	})
}

func OnMapResultHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnMapResultPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnMapResult(p)
	})
}

func OnSeriesResultHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnSeriesResultPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnSeriesResult(p)
	})
}

func OnSidePickedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnSidePickedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnSidePicked(p)
	})
}

func OnMapPickedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnMapPickedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnMapPicked(p)
	})
}

func OnMapVetoedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnMapVetoedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnMapVetoed(p)
	})
}

func OnBackupRestoreHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnBackupRestorePayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnBackupRestore(p)
	})
}

func OnDemoFinishedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnDemoFinishedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnDemoFinished(p)
	})
}

func OnDemoUploadEndedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnDemoUploadEndedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnDemoUploadEnded(p)
	})
}

func OnMatchPausedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnMatchPausedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnMatchPaused(p)
	})
}

func OnMatchUnpausedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnMatchUnpausedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnMatchUnpaused(p)
	})
}

func OnKnifeRoundStartedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnKnifeRoundStartedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnKnifeRoundStarted(p)
	})
}

func OnKnifeRoundWonHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnKnifeRoundWonPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnKnifeRoundWon(p)
	})
}

func OnTeamReadyStatusChanngeHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnTeamReadyStatusChangedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnTeamReadyStatusChanged(p)
	})
}

func OnGoingLiveHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnGoingLivePayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnGoingLive(p)
	})
}

func OnRoundStartHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnRoundStartPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnRoundStart(p)
	})
}

func OnRoundEndHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnRoundEndPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnRoundEnd(p)
	})
}

func OnRoundStatusUpdatedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnRoundStatsUpdatedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnRoundStatsUpdated(p)
	})
}

func OnPlayerBecameMVPHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnPlayerBecameMVPPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnPlayerBecameMVP(p)
	})
}

func OnGrenadeThrownHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnGrenadeThrownPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnGrenadeThrown(p)
	})
}

func OnPlayerDeathHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnPlayerDeathPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnPlayerDeath(p)
	})
}

func OnHEGrenadeDetonatedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnHEGrenadeDetonatedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnHEGrenadeDetonated(p)
	})
}

func OnMolotovDetonatedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnMolotovDetonatedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnMolotovDetonated(p)
	})
}

func OnFlashbangDetonatedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnFlashbangDetonatedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnFlashbangDetonated(p)
	})
}

func OnSmokeGrenadeDetonatedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnSmokeGrenadeDetonatedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnSmokeGrenadeDetonated(p)
	})
}

func OnDecoyStartedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnDecoyStartedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnDecoyStarted(p)
	})
}

func OnBombPlantedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnBombPlantedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnBombPlanted(p)
	})
}

func OnBombDefusedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnBombDefusedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnBombDefused(p)
	})
}

func OnBombExpodedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnBombExplodedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnBombExploded(p)
	})
}

func OnPlayerConnectedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnPlayerConnectedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnPlayerConnected(p)
	})
}

func OnPlayerDisconnectedHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnPlayerDisconnectedPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnPlayerDisconnected(p)
	})
}

// OnPlayerSayHandler POST on /Get5_OnPlayerSay
func OnPlayerSayHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := models.OnPlayerSayPayload{}
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		return ctrl.HandleOnPlayerSay(p)
	})
}
