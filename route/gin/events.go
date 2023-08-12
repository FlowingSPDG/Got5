package ginroute

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"

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
func OnEventHandler(ctrl got5.EventHandler) func(c *gin.Context) {
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
			ret := got5.OnGameStateChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnGameStateChanged(c, ret)
		case "preload_match_config":
			ret := got5.OnPreLoadMatchConfigPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnPreLoadMatchConfig(c, ret)
		case "match_config_load_fail":
			ret := got5.OnLoadMatchConfigFailedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnLoadMatchConfigFailed(c, ret)
		case "series_start":
			ret := got5.OnSeriesInitPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnSeriesInit(c, ret)
		case "map_result":
			ret := got5.OnMapResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
			}
			ctrl.HandleOnMapResult(c, ret)
		case "series_end":
			ret := got5.OnSeriesResultPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnSeriesResult(c, ret)
		case "side_picked":
			ret := got5.OnSidePickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnSidePicked(c, ret)
		case "map_picked":
			ret := got5.OnMapPickedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMapPicked(c, ret)
		case "map_vetoed":
			ret := got5.OnMapVetoedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMapVetoed(c, ret)
		case "backup_loaded":
			ret := got5.OnBackupRestorePayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnBackupRestore(c, ret)
		case "demo_finished":
			ret := got5.OnDemoFinishedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnDemoFinished(c, ret)
		case "demo_upload_ended":
			ret := got5.OnDemoUploadEndedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnDemoUploadEnded(c, ret)
		case "game_paused":
			ret := got5.OnMatchPausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMatchPaused(c, ret)
		case "game_unpaused":
			ret := got5.OnMatchUnpausedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMatchUnpaused(c, ret)
		case "pause_began":
			ret := got5.OnPauseBeganPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPauseBegan(c, ret)
		case "knife_start":
			ret := got5.OnKnifeRoundStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnKnifeRoundStarted(c, ret)
		case "knife_won":
			ret := got5.OnKnifeRoundWonPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnKnifeRoundWon(c, ret)
		case "team_ready_status_changed":
			ret := got5.OnTeamReadyStatusChangedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnTeamReadyStatusChanged(c, ret)
		case "going_live":
			ret := got5.OnGoingLivePayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnGoingLive(c, ret)
		case "round_start":
			ret := got5.OnRoundStartPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnRoundStart(c, ret)
		case "round_end":
			ret := got5.OnRoundEndPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnRoundEnd(c, ret)
		case "stats_updated":
			ret := got5.OnRoundStatsUpdatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnRoundStatsUpdated(c, ret)
		case "round_mvp":
			ret := got5.OnPlayerBecameMVPPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerBecameMVP(c, ret)
		case "grenade_thrown":
			ret := got5.OnGrenadeThrownPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnGrenadeThrown(c, ret)
		case "player_death":
			ret := got5.OnPlayerDeathPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerDeath(c, ret)
		case "hegrenade_detonated":
			ret := got5.OnHEGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnHEGrenadeDetonated(c, ret)
		case "molotov_detonated":
			ret := got5.OnMolotovDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnMolotovDetonated(c, ret)
		case "flashbang_detonated":
			ret := got5.OnFlashbangDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnFlashbangDetonated(c, ret)
		case "smokegrenade_detonated":
			ret := got5.OnSmokeGrenadeDetonatedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnSmokeGrenadeDetonated(c, ret)
		case "decoygrenade_started":
			ret := got5.OnDecoyStartedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnDecoyStarted(c, ret)
		case "bomb_planted":
			ret := got5.OnBombPlantedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnBombPlanted(c, ret)
		case "bomb_defused":
			ret := got5.OnBombDefusedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnBombDefused(c, ret)
		case "bomb_exploded":
			ret := got5.OnBombExplodedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnBombExploded(c, ret)
		case "player_connect":
			ret := got5.OnPlayerConnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerConnected(c, ret)
		case "player_disconnect":
			ret := got5.OnPlayerDisconnectedPayload{}
			if err := reMarshal(p, &ret); err != nil {
				c.AbortWithError(http.StatusBadRequest, xerrors.New("Invalid JSON")) // TODO: Wrap error code
				return
			}
			ctrl.HandleOnPlayerDisconnected(c, ret)
		case "player_say":
			ret := got5.OnPlayerSayPayload{}
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
func CheckEventHandlerAuth(auth got5.Auth) func(c *gin.Context) {
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
