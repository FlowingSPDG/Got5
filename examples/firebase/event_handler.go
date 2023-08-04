package fb

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	storage "firebase.google.com/go/storage"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

var _ controller.EventHandler = (*firebaseEventHandler)(nil)

type firebaseEventHandler struct {
	*Database // Databaseを埋め込んでおく
	fs        *firestore.Client
	s         *storage.Client
	setting   setting
}

// ControllerSetting Settings
type setting struct {
	Hostname string
	Port     int
}

// Hostname implements controller.EventHandler
func (f *firebaseEventHandler) Hostname() string {
	return f.setting.Hostname
}

func (f *firebaseEventHandler) Close() error {
	if err := f.fs.Close(); err != nil {
		return err
	}
	return nil
}

// HandleOnBackupRestore implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	panic("unimplemented")
}

// HandleOnBombDefused implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	panic("unimplemented")
}

// HandleOnBombExploded implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	panic("unimplemented")
}

// HandleOnBombPlanted implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	panic("unimplemented")
}

// HandleOnDecoyStarted implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	panic("unimplemented")
}

// HandleOnDemoFinished implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	panic("unimplemented")
}

// HandleOnDemoUploadEnded implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	panic("unimplemented")
}

// HandleOnEvent implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnEvent(ctx context.Context, p models.OnEventPayload) error {
	panic("unimplemented")
}

// HandleOnFlashbangDetonated implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnGameStateChanged implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	panic("unimplemented")
}

// HandleOnGoingLive implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	panic("unimplemented")
}

// HandleOnGrenadeThrown implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	panic("unimplemented")
}

// HandleOnHEGrenadeDetonated implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnKnifeRoundStarted implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	panic("unimplemented")
}

// HandleOnKnifeRoundWon implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	panic("unimplemented")
}

// HandleOnLoadMatchConfigFailed implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	panic("unimplemented")
}

// HandleOnMapPicked implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	panic("unimplemented")
}

// HandleOnMapResult implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	panic("unimplemented")
}

// HandleOnMapVetoed implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	panic("unimplemented")
}

// HandleOnMatchPaused implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	panic("unimplemented")
}

// HandleOnMatchUnpaused implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	panic("unimplemented")
}

// HandleOnPauseBegan implements controller.EventHandler.
func (f *firebaseEventHandler) HandleOnPauseBegan(ctx context.Context, p models.OnPauseBeganPayload) error {
	panic("unimplemented")
}

// HandleOnMolotovDetonated implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerBecameMVP implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerConnected implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerDeath implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerDisconnected implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerSay implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	// Firestoreに書き込むデータのフィールド名を変更したい場合は
	// 独自のstructを定義して詰め替えることを推奨
	/*
		type payload struct {
		    Event       string `firestore:"event"`
		    Matchid     string `firestore:"matchid"`
		    MapNumber   int    `firestore:"map_number"`
		    RoundNumber int    `firestore:"round_number"`
		    RoundTime   int    `firestore:"round_time"`
		    Player      Player `firestore:"player"`
		    Command     string `firestore:"command"`
		    Message     string `firestore:"message"`
		}
		p を定義したpayloadへ詰め替える
		詰め替えたpayload構造体を書き込む
	*/

	// またはjson.Marshal & Unmarshal でmap[string]any型にキャストし、それを書き込む形でも可能
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	data := make(map[string]any)
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	_, _, err = f.fs.Collection(CollectionChat).Add(ctx, data)
	if err != nil {
		return err
	}
	return nil

	// 面倒な場合、そのまま書き込んでも良いがフィールド名はFirestoreの自動生成となる
	// _, _, err := f.fs.Collection(CollectionChat).Add(ctx, p)
	// return err
}

// HandleOnPreLoadMatchConfig implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	panic("unimplemented")
}

// HandleOnRoundEnd implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	panic("unimplemented")
}

// HandleOnRoundStart implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	panic("unimplemented")
}

// HandleOnRoundStatsUpdated implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	panic("unimplemented")
}

// HandleOnSeriesInit implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	panic("unimplemented")
}

// HandleOnSeriesResult implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	panic("unimplemented")
}

// HandleOnSidePicked implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	panic("unimplemented")
}

// HandleOnSmokeGrenadeDetonated implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnTeamReadyStatusChanged implements controller.EventHandler
func (f *firebaseEventHandler) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	panic("unimplemented")
}

// NewEventHandler Get Firebase Event Handler
func NewEventHandler(ctx context.Context, c *firebase.App) (controller.EventHandler, error) {
	// Firestore
	fs, err := c.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	// Firebase Storage
	s, err := c.Storage(ctx)
	if err != nil {
		return nil, err
	}

	// Database for writing
	db, err := NewDatabase(ctx, c)
	if err != nil {
		return nil, err
	}

	return &firebaseEventHandler{
		Database: db,
		fs:       fs,
		s:        s,
		setting:  setting{},
	}, nil
}
