package route

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

func reMarshal(m map[string]any, p any) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, p); err != nil {
		return err
	}
	return nil
}

// OnEventHandler POST on /Get5_OnEvent
func OnEventHandler(ctrl controller.EventHandler) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := make(map[string]any)
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		ev, ok := (p["event"]).(string)
		if !ok {
			return fmt.Errorf("Unsupported JSON Format")
		}
		switch ev {
		case "game_state_changed":
			ret := models.OnGameStateChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnGameStateChanged(c.Context(), ret)
		case "preload_match_config":
			ret := models.OnPreLoadMatchConfigPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnPreLoadMatchConfig(c.Context(), ret)
		case "match_config_load_fail":
			ret := models.OnLoadMatchConfigFailedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnLoadMatchConfigFailed(c.Context(), ret)
		case "series_start":
			ret := models.OnSeriesInitPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnSeriesInit(c.Context(), ret)
		case "map_result":
			ret := models.OnMapResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnMapResult(c.Context(), ret)
		case "series_end":
			ret := models.OnSeriesResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnSeriesResult(c.Context(), ret)
		case "side_picked":
			ret := models.OnSidePickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnSidePicked(c.Context(), ret)
		case "map_picked":
			ret := models.OnMapPickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnMapPicked(c.Context(), ret)
		case "map_vetoed":
			ret := models.OnMapVetoedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnMapVetoed(c.Context(), ret)
		case "backup_loaded":
			ret := models.OnBackupRestorePayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnBackupRestore(c.Context(), ret)
		case "demo_finished":
			ret := models.OnDemoFinishedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnDemoFinished(c.Context(), ret)
		case "demo_upload_ended":
			ret := models.OnDemoUploadEndedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnDemoUploadEnded(c.Context(), ret)
		case "game_paused":
			ret := models.OnMatchPausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnMatchPaused(c.Context(), ret)
		case "game_unpaused":
			ret := models.OnMatchUnpausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnMatchUnpaused(c.Context(), ret)
		case "knife_start":
			ret := models.OnKnifeRoundStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnKnifeRoundStarted(c.Context(), ret)
		case "knife_won":
			ret := models.OnKnifeRoundWonPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnKnifeRoundWon(c.Context(), ret)
		case "team_ready_status_changed":
			ret := models.OnTeamReadyStatusChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnTeamReadyStatusChanged(c.Context(), ret)
		case "going_live":
			ret := models.OnGoingLivePayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnGoingLive(c.Context(), ret)
		case "round_start":
			ret := models.OnRoundStartPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnRoundStart(c.Context(), ret)
		case "round_end":
			ret := models.OnRoundEndPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnRoundEnd(c.Context(), ret)
		case "stats_updated":
			ret := models.OnRoundStatsUpdatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnRoundStatsUpdated(c.Context(), ret)
		case "round_mvp":
			ret := models.OnPlayerBecameMVPPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnPlayerBecameMVP(c.Context(), ret)
		case "grenade_thrown":
			ret := models.OnGrenadeThrownPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnGrenadeThrown(c.Context(), ret)
		case "player_death":
			ret := models.OnPlayerDeathPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnPlayerDeath(c.Context(), ret)
		case "hegrenade_detonated":
			ret := models.OnHEGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnHEGrenadeDetonated(c.Context(), ret)
		case "molotov_detonated":
			ret := models.OnMolotovDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnMolotovDetonated(c.Context(), ret)
		case "flashbang_detonated":
			ret := models.OnFlashbangDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnFlashbangDetonated(c.Context(), ret)
		case "smokegrenade_detonated":
			ret := models.OnSmokeGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnSmokeGrenadeDetonated(c.Context(), ret)
		case "decoygrenade_started":
			ret := models.OnDecoyStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnDecoyStarted(c.Context(), ret)
		case "bomb_planted":
			ret := models.OnBombPlantedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnBombPlanted(c.Context(), ret)
		case "bomb_defused":
			ret := models.OnBombDefusedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnBombDefused(c.Context(), ret)
		case "bomb_exploded":
			ret := models.OnBombExplodedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnBombExploded(c.Context(), ret)
		case "player_connect":
			ret := models.OnPlayerConnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnPlayerConnected(c.Context(), ret)
		case "player_disconnect":
			ret := models.OnPlayerDisconnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnPlayerDisconnected(c.Context(), ret)
		case "player_say":
			ret := models.OnPlayerSayPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			return ctrl.HandleOnPlayerSay(c.Context(), ret)
		default:
			return fmt.Errorf("Event Not found%s", ev)
		}
	})
}

// CheckEventHandlerAuth 認証用ハンドラ
func CheckEventHandlerAuth(ctrl controller.EventHandler) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := make(map[string]any)
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		mid, ok := (p["matchid"]).(string)
		if !ok {
			return fmt.Errorf("Unsupported JSON Format")
		}

		reqAuthVal := c.Get("Authorization")

		if err := ctrl.CheckEventAuth(c.Context(), mid, reqAuthVal); err != nil {
			c.Status(fiber.StatusUnauthorized)
			return nil
		}
		return c.Next()
	})
}
