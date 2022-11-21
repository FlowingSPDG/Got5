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
			ret := models.OnGameStateChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnGameStateChanged(c.Context(), ret)
		case "preload_match_config":
			ret := models.OnPreLoadMatchConfigPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPreLoadMatchConfig(c.Context(), ret)
		case "match_config_load_fail":
			ret := models.OnLoadMatchConfigFailedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnLoadMatchConfigFailed(c.Context(), ret)
		case "series_start":
			ret := models.OnSeriesInitPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnSeriesInit(c.Context(), ret)
		case "map_result":
			ret := models.OnMapResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMapResult(c.Context(), ret)
		case "series_end":
			ret := models.OnSeriesResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnSeriesResult(c.Context(), ret)
		case "side_picked":
			ret := models.OnSidePickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnSidePicked(c.Context(), ret)
		case "map_picked":
			ret := models.OnMapPickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMapPicked(c.Context(), ret)
		case "map_vetoed":
			ret := models.OnMapVetoedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMapVetoed(c.Context(), ret)
		case "backup_loaded":
			ret := models.OnBackupRestorePayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnBackupRestore(c.Context(), ret)
		case "demo_finished":
			ret := models.OnDemoFinishedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnDemoFinished(c.Context(), ret)
		case "demo_upload_ended":
			ret := models.OnDemoUploadEndedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnDemoUploadEnded(c.Context(), ret)
		case "game_paused":
			ret := models.OnMatchPausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMatchPaused(c.Context(), ret)
		case "game_unpaused":
			ret := models.OnMatchUnpausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMatchUnpaused(c.Context(), ret)
		case "knife_start":
			ret := models.OnKnifeRoundStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnKnifeRoundStarted(c.Context(), ret)
		case "knife_won":
			ret := models.OnKnifeRoundWonPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnKnifeRoundWon(c.Context(), ret)
		case "team_ready_status_changed":
			ret := models.OnTeamReadyStatusChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnTeamReadyStatusChanged(c.Context(), ret)
		case "going_live":
			ret := models.OnGoingLivePayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnGoingLive(c.Context(), ret)
		case "round_start":
			ret := models.OnRoundStartPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnRoundStart(c.Context(), ret)
		case "round_end":
			ret := models.OnRoundEndPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnRoundEnd(c.Context(), ret)
		case "stats_updated":
			ret := models.OnRoundStatsUpdatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnRoundStatsUpdated(c.Context(), ret)
		case "round_mvp":
			ret := models.OnPlayerBecameMVPPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerBecameMVP(c.Context(), ret)
		case "grenade_thrown":
			ret := models.OnGrenadeThrownPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnGrenadeThrown(c.Context(), ret)
		case "player_death":
			ret := models.OnPlayerDeathPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerDeath(c.Context(), ret)
		case "hegrenade_detonated":
			ret := models.OnHEGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnHEGrenadeDetonated(c.Context(), ret)
		case "molotov_detonated":
			ret := models.OnMolotovDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnMolotovDetonated(c.Context(), ret)
		case "flashbang_detonated":
			ret := models.OnFlashbangDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnFlashbangDetonated(c.Context(), ret)
		case "smokegrenade_detonated":
			ret := models.OnSmokeGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnSmokeGrenadeDetonated(c.Context(), ret)
		case "decoygrenade_started":
			ret := models.OnDecoyStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnDecoyStarted(c.Context(), ret)
		case "bomb_planted":
			ret := models.OnBombPlantedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnBombPlanted(c.Context(), ret)
		case "bomb_defused":
			ret := models.OnBombDefusedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnBombDefused(c.Context(), ret)
		case "bomb_exploded":
			ret := models.OnBombExplodedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnBombExploded(c.Context(), ret)
		case "player_connect":
			ret := models.OnPlayerConnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerConnected(c.Context(), ret)
		case "player_disconnect":
			ret := models.OnPlayerDisconnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.SendString(fmt.Sprintf("Invalid JSON:%s", err.Error()))
				return c.SendStatus(fiber.StatusBadRequest)
			}
			return ctrl.HandleOnPlayerDisconnected(c.Context(), ret)
		case "player_say":
			ret := models.OnPlayerSayPayload{}
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

// CheckEventHandlerAuth 認証用ハンドラ Under construction, this middleware is not working now.
func CheckEventHandlerAuth(ctrl controller.EventHandler) func(c *fiber.Ctx) error {
	return (func(c *fiber.Ctx) error {
		// TODO: マッチIDはどう取得する？
		// https://github.com/FlowingSPDG/Got5/issues/25

		// ヘッダーのKey/Valueは一つだけ設定出来るが、それだと認証用キーとマッチIDが同時に設定出来ない
		// だが、マッチがロードされていない際のイベントにはMatchIDが付与されないので、送ることが難しい...
		//
		// [IDEA1] get5_server_id を設定(数字のみだが)し、これをベースに認証を行う方法もあるが実装が煩雑になりそう
		// [IDEA2] ValueにマッチIDと認証用の鍵を詰め込み、両方扱える形にする
		//
		//		e.g.1 アンダースコアで分割する
		//			get5_remote_log_header_value "{MatchID}_{AuthKey}""
		//		e.g.2 JSON形式にする
		//			get5_remote_log_header_value "{"matchid":"{MatchID}"","auth":"{AuthKey}"}"
		//		e.g.3 JWTを使用する
		//			get5_remote_log_header_value "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

		// (in English for any suggestions/help)
		// TODO: How to get MatchID of event?
		// There is a cvar that sets Key/Value on HTTP header but it is not possible to send Auth key and MatchID in same time
		// But server will NOT send MatchID if no match loaded
		//
		// [IDEA1] set get5_server_id cvar(integer only) and get match ID from this, but this could be more complicated Database structure
		// [IDEA2] Put both of MatchID and Auth string into one HTTP header value
		//
		//		e.g.1 split by underscore
		//			get5_remote_log_header_value "{MatchID}_{AuthKey}""
		//		e.g.2 Send via JSON
		//			get5_remote_log_header_value "{"matchid":"{MatchID}"","auth":"{AuthKey}"}"
		//		e.g.3 Use JWT
		//			get5_remote_log_header_value "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

		// reqAuthVal := c.Get("Authorization")
		/*
			if err := ctrl.CheckEventAuth(c.Context(), mid, reqAuthVal); err != nil {
				c.Status(fiber.StatusUnauthorized)
				return nil
			}
		*/
		return c.Next()
	})
}
