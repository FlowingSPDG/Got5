package fiberroute

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	got5 "github.com/FlowingSPDG/Got5"
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
func OnEventHandler(ctrl got5.EventHandler) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		p := make(map[string]any)
		if err := c.BodyParser(&p); err != nil {
			c.SendString(err.Error()) // TODO: Wrap error code
			return c.SendStatus(fiber.StatusBadRequest)
		}
		ev, ok := (p["event"]).(string)
		if !ok {
			c.SendString("Invalid JSON")
			return c.SendStatus(fiber.StatusBadRequest)
		}
		switch ev {
		case "game_state_changed":
			ret := got5.OnGameStateChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnGameStateChanged(c.Context(), ret)
		case "preload_match_config":
			ret := got5.OnPreLoadMatchConfigPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPreLoadMatchConfig(c.Context(), ret)
		case "match_config_load_fail":
			ret := got5.OnLoadMatchConfigFailedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnLoadMatchConfigFailed(c.Context(), ret)
		case "series_start":
			ret := got5.OnSeriesInitPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnSeriesInit(c.Context(), ret)
		case "map_result":
			ret := got5.OnMapResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMapResult(c.Context(), ret)
		case "series_end":
			ret := got5.OnSeriesResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnSeriesResult(c.Context(), ret)
		case "side_picked":
			ret := got5.OnSidePickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnSidePicked(c.Context(), ret)
		case "map_picked":
			ret := got5.OnMapPickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMapPicked(c.Context(), ret)
		case "map_vetoed":
			ret := got5.OnMapVetoedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMapVetoed(c.Context(), ret)
		case "backup_loaded":
			ret := got5.OnBackupRestorePayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnBackupRestore(c.Context(), ret)
		case "demo_finished":
			ret := got5.OnDemoFinishedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnDemoFinished(c.Context(), ret)
		case "demo_upload_ended":
			ret := got5.OnDemoUploadEndedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnDemoUploadEnded(c.Context(), ret)
		case "game_paused":
			ret := got5.OnMatchPausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMatchPaused(c.Context(), ret)
		case "game_unpaused":
			ret := got5.OnMatchUnpausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMatchUnpaused(c.Context(), ret)
		case "pause_began":
			ret := got5.OnPauseBeganPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error())) // TODO: Wrap error code
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPauseBegan(c.Context(), ret)
		case "knife_start":
			ret := got5.OnKnifeRoundStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnKnifeRoundStarted(c.Context(), ret)
		case "knife_won":
			ret := got5.OnKnifeRoundWonPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnKnifeRoundWon(c.Context(), ret)
		case "team_ready_status_changed":
			ret := got5.OnTeamReadyStatusChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnTeamReadyStatusChanged(c.Context(), ret)
		case "going_live":
			ret := got5.OnGoingLivePayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnGoingLive(c.Context(), ret)
		case "round_start":
			ret := got5.OnRoundStartPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnRoundStart(c.Context(), ret)
		case "round_end":
			ret := got5.OnRoundEndPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnRoundEnd(c.Context(), ret)
		case "stats_updated":
			ret := got5.OnRoundStatsUpdatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnRoundStatsUpdated(c.Context(), ret)
		case "round_mvp":
			ret := got5.OnPlayerBecameMVPPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerBecameMVP(c.Context(), ret)
		case "grenade_thrown":
			ret := got5.OnGrenadeThrownPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnGrenadeThrown(c.Context(), ret)
		case "player_death":
			ret := got5.OnPlayerDeathPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerDeath(c.Context(), ret)
		case "hegrenade_detonated":
			ret := got5.OnHEGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnHEGrenadeDetonated(c.Context(), ret)
		case "molotov_detonated":
			ret := got5.OnMolotovDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMolotovDetonated(c.Context(), ret)
		case "flashbang_detonated":
			ret := got5.OnFlashbangDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnFlashbangDetonated(c.Context(), ret)
		case "smokegrenade_detonated":
			ret := got5.OnSmokeGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnSmokeGrenadeDetonated(c.Context(), ret)
		case "decoygrenade_started":
			ret := got5.OnDecoyStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnDecoyStarted(c.Context(), ret)
		case "bomb_planted":
			ret := got5.OnBombPlantedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnBombPlanted(c.Context(), ret)
		case "bomb_defused":
			ret := got5.OnBombDefusedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnBombDefused(c.Context(), ret)
		case "bomb_exploded":
			ret := got5.OnBombExplodedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnBombExploded(c.Context(), ret)
		case "player_connect":
			ret := got5.OnPlayerConnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerConnected(c.Context(), ret)
		case "player_disconnect":
			ret := got5.OnPlayerDisconnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerDisconnected(c.Context(), ret)
		case "player_say":
			ret := got5.OnPlayerSayPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerSay(c.Context(), ret)
		default:
			c.SendString(fmt.Sprintf("Invalid Event"))
			return c.SendStatus(fiber.StatusBadRequest)
		}
	})
}

// CheckEventHandlerAuth Check Auth
func CheckEventHandlerAuth(auth got5.Auth) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// NOTICE: get5_remote_log_header_value only supports max 128 characters
		reqAuth := c.Get("Authorization")
		serverID := c.Get("Get5-ServerId")
		if err := auth.EventAuth(c.Context(), serverID, reqAuth); err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error()) // カスタムエラーを返したい
		}
		return c.Next()
	})
}
