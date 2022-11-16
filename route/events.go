package route

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

func reMarshal(m map[string]interface{}, p any) error {
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
func OnEventHandler(ctrl controller.Controller) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// p := models.OnEventPayload{}
		p := make(map[string]interface{})
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
			ctrl.HandleOnGameStateChanged(c.Context(), ret)
		case "preload_match_config":
			ret := models.OnPreLoadMatchConfigPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnPreLoadMatchConfig(c.Context(), ret)
		case "match_config_load_fail":
			ret := models.OnLoadMatchConfigFailedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnLoadMatchConfigFailed(c.Context(), ret)
		case "series_start":
			ret := models.OnSeriesInitPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnSeriesInit(c.Context(), ret)
		case "map_result":
			ret := models.OnMapResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnMapResult(c.Context(), ret)
		case "series_end":
			ret := models.OnSeriesResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnSeriesResult(c.Context(), ret)
		case "side_picked":
			ret := models.OnSidePickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnSidePicked(c.Context(), ret)
		case "map_picked":
			ret := models.OnMapPickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnMapPicked(c.Context(), ret)
		case "map_vetoed":
			ret := models.OnMapVetoedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnMapVetoed(c.Context(), ret)
		case "backup_loaded":
			ret := models.OnBackupRestorePayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnBackupRestore(c.Context(), ret)
		case "demo_finished":
			ret := models.OnDemoFinishedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnDemoFinished(c.Context(), ret)
		case "demo_upload_ended":
			ret := models.OnDemoUploadEndedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnDemoUploadEnded(c.Context(), ret)
		case "game_paused":
			ret := models.OnMatchPausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnMatchPaused(c.Context(), ret)
		case "game_unpaused":
			ret := models.OnMatchUnpausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnMatchUnpaused(c.Context(), ret)
		case "knife_start":
			ret := models.OnKnifeRoundStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnKnifeRoundStarted(c.Context(), ret)
		case "knife_won":
			ret := models.OnKnifeRoundWonPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnKnifeRoundWon(c.Context(), ret)
		case "team_ready_status_changed":
			ret := models.OnTeamReadyStatusChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnTeamReadyStatusChanged(c.Context(), ret)
		case "going_live":
			ret := models.OnGoingLivePayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnGoingLive(c.Context(), ret)
		case "round_start":
			ret := models.OnRoundStartPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnRoundStart(c.Context(), ret)
		case "round_end":
			ret := models.OnRoundEndPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnRoundEnd(c.Context(), ret)
		case "stats_updated":
			ret := models.OnRoundStatsUpdatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnRoundStatsUpdated(c.Context(), ret)
		case "round_mvp":
			ret := models.OnPlayerBecameMVPPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnPlayerBecameMVP(c.Context(), ret)
		case "grenade_thrown":
			ret := models.OnGrenadeThrownPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnGrenadeThrown(c.Context(), ret)
		case "player_death":
			ret := models.OnPlayerDeathPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnPlayerDeath(c.Context(), ret)
		case "hegrenade_detonated":
			ret := models.OnHEGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnHEGrenadeDetonated(c.Context(), ret)
		case "molotov_detonated":
			ret := models.OnMolotovDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnMolotovDetonated(c.Context(), ret)
		case "flashbang_detonated":
			ret := models.OnFlashbangDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnFlashbangDetonated(c.Context(), ret)
		case "smokegrenade_detonated":
			ret := models.OnSmokeGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnSmokeGrenadeDetonated(c.Context(), ret)
		case "decoygrenade_started":
			ret := models.OnDecoyStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnDecoyStarted(c.Context(), ret)
		case "bomb_planted":
			ret := models.OnBombPlantedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnBombPlanted(c.Context(), ret)
		case "bomb_defused":
			ret := models.OnBombDefusedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnBombDefused(c.Context(), ret)
		case "bomb_exploded":
			ret := models.OnBombExplodedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnBombExploded(c.Context(), ret)
		case "player_connect":
			ret := models.OnPlayerConnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnPlayerConnected(c.Context(), ret)
		case "player_disconnect":
			ret := models.OnPlayerDisconnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnPlayerDisconnected(c.Context(), ret)
		case "player_say":
			ret := models.OnPlayerSayPayload{}
			if err := reMarshal(p, &ret); err != nil {
				return err
			}
			ctrl.HandleOnPlayerSay(c.Context(), ret)
		default:
			return fmt.Errorf("Event Not found%s", ev)
		}
		return nil
	})
}
