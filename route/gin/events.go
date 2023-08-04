package ginroute

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"

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
func OnEventHandler(ctrl controller.EventHandler) func(c *gin.Context) {
	return (func(c *gin.Context) {
		p := make(map[string]any)
		if err := c.ShouldBindJSON(&p); err != nil {
			c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			return
		}
		ev, ok := (p["event"]).(string)
		if !ok {
			c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			return
		}
		switch ev {
		case "game_state_changed":
			ret := models.OnGameStateChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnGameStateChanged(c, ret)
		case "preload_match_config":
			ret := models.OnPreLoadMatchConfigPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnPreLoadMatchConfig(c, ret)
		case "match_config_load_fail":
			ret := models.OnLoadMatchConfigFailedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnLoadMatchConfigFailed(c, ret)
		case "series_start":
			ret := models.OnSeriesInitPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnSeriesInit(c, ret)
		case "map_result":
			ret := models.OnMapResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnMapResult(c, ret)
		case "series_end":
			ret := models.OnSeriesResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnSeriesResult(c, ret)
		case "side_picked":
			ret := models.OnSidePickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnSidePicked(c, ret)
		case "map_picked":
			ret := models.OnMapPickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMapPicked(c, ret)
		case "map_vetoed":
			ret := models.OnMapVetoedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMapVetoed(c, ret)
		case "backup_loaded":
			ret := models.OnBackupRestorePayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnBackupRestore(c, ret)
		case "demo_finished":
			ret := models.OnDemoFinishedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnDemoFinished(c, ret)
		case "demo_upload_ended":
			ret := models.OnDemoUploadEndedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnDemoUploadEnded(c, ret)
		case "game_paused":
			ret := models.OnMatchPausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMatchPaused(c, ret)
		case "game_unpaused":
			ret := models.OnMatchUnpausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMatchUnpaused(c, ret)
		case "pause_began":
			ret := models.OnPauseBeganPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPauseBegan(c, ret)
		case "knife_start":
			ret := models.OnKnifeRoundStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnKnifeRoundStarted(c, ret)
		case "knife_won":
			ret := models.OnKnifeRoundWonPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnKnifeRoundWon(c, ret)
		case "team_ready_status_changed":
			ret := models.OnTeamReadyStatusChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnTeamReadyStatusChanged(c, ret)
		case "going_live":
			ret := models.OnGoingLivePayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnGoingLive(c, ret)
		case "round_start":
			ret := models.OnRoundStartPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnRoundStart(c, ret)
		case "round_end":
			ret := models.OnRoundEndPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnRoundEnd(c, ret)
		case "stats_updated":
			ret := models.OnRoundStatsUpdatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnRoundStatsUpdated(c, ret)
		case "round_mvp":
			ret := models.OnPlayerBecameMVPPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerBecameMVP(c, ret)
		case "grenade_thrown":
			ret := models.OnGrenadeThrownPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnGrenadeThrown(c, ret)
		case "player_death":
			ret := models.OnPlayerDeathPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerDeath(c, ret)
		case "hegrenade_detonated":
			ret := models.OnHEGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnHEGrenadeDetonated(c, ret)
		case "molotov_detonated":
			ret := models.OnMolotovDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMolotovDetonated(c, ret)
		case "flashbang_detonated":
			ret := models.OnFlashbangDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnFlashbangDetonated(c, ret)
		case "smokegrenade_detonated":
			ret := models.OnSmokeGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnSmokeGrenadeDetonated(c, ret)
		case "decoygrenade_started":
			ret := models.OnDecoyStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnDecoyStarted(c, ret)
		case "bomb_planted":
			ret := models.OnBombPlantedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnBombPlanted(c, ret)
		case "bomb_defused":
			ret := models.OnBombDefusedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnBombDefused(c, ret)
		case "bomb_exploded":
			ret := models.OnBombExplodedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnBombExploded(c, ret)
		case "player_connect":
			ret := models.OnPlayerConnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerConnected(c, ret)
		case "player_disconnect":
			ret := models.OnPlayerDisconnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerDisconnected(c, ret)
		case "player_say":
			ret := models.OnPlayerSayPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerSay(c, ret)
		default:
			c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid Event"))
		}
	})
}

// CheckEventHandlerAuth Check Auth
func CheckEventHandlerAuth(auth controller.Auth) func(c *gin.Context) {
	return (func(c *gin.Context) {
		// NOTICE: get5_remote_log_header_value only supports max 128 characters
		reqAuth := c.GetHeader("Authorization")
		serverID := c.GetHeader("Get5-ServerId")
		if err := auth.EventAuth(c, serverID, reqAuth); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err) // カスタムエラーを返したい
			return
		}
		c.Next()
	})
}
