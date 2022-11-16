package fb

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	storage "firebase.google.com/go/storage"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

// ControllerSetting Settings
type ControllerSetting struct {
	Hostname string
	Port     int
}

// fs is Get5 API Database on firestore
type f struct {
	setting ControllerSetting
	fs      *firestore.Client
	s       *storage.Client
}

// UpdateMatch implements controller.Controller
func (f *f) UpdateMatch(ctx context.Context, mid string, m models.Match) error {
	if _, err := f.fs.Collection(CollectionMatch).Doc(mid).Set(ctx, m, firestore.MergeAll); err != nil {
		return err
	}
	return nil
}

// Hostname implements controller.Controller
func (f *f) Hostname() string {
	return f.setting.Hostname
}

func (f *f) Close() error {
	if err := f.fs.Close(); err != nil {
		return err
	}
	return nil
}

// RegisterDemoFile implements controller.Controller
func (f *f) RegisterDemoFile(ctx context.Context, bucket string, mid string, filename string, b []byte) error {
	// MatchIDを取得しているので、Firestore上のMatchにdemoのURLを記載しても良い(TODO)
	bh, err := f.s.Bucket(bucket)
	if err != nil {
		return err
	}

	writer := bh.Object(filename).NewWriter(ctx)
	defer writer.Close()

	writer.ObjectAttrs.ContentType = "application/octet-stream"
	writer.ObjectAttrs.CacheControl = "public,max-age=86400"
	// 必要であれば権限設定を実施

	if _, err := io.Copy(writer, bytes.NewReader(b)); err != nil {
		return err
	}
	return nil
}

// RegisterMatch implements controller.Controller
func (f *f) RegisterMatch(ctx context.Context, m models.Match) (models.Match, error) {
	ret := models.Match{}
	// models.G5Match を使った方が汎用性が高いか？
	fm := toFirestoreMatch(m)
	ref := f.fs.Collection(CollectionMatch).NewDoc()
	fm.MatchID = ref.ID // MatchIDを上書きする
	_, err := ref.Set(ctx, fm)
	if err != nil {
		return ret, err
	}
	snap, err := ref.Get(ctx)
	if err != nil {
		return ret, err
	}
	if err := snap.DataTo(&ret); err != nil {
		return ret, err
	}
	return ret, nil
}

// GetMatch implements controller.Controller
func (f *f) GetMatch(ctx context.Context, mid string) (models.G5Match, error) {
	ret := Match{}
	dsnap, err := f.fs.Collection(CollectionMatch).Doc(mid).Get(ctx)
	if err != nil {
		return ret, err
	}
	if err := dsnap.DataTo(&ret); err != nil {
		return ret, err
	}
	return ret, nil
}

// HandleOnBackupRestore implements controller.Controller
func (f *f) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	panic("unimplemented")
}

// HandleOnBombDefused implements controller.Controller
func (f *f) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	panic("unimplemented")
}

// HandleOnBombExploded implements controller.Controller
func (f *f) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	panic("unimplemented")
}

// HandleOnBombPlanted implements controller.Controller
func (f *f) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	panic("unimplemented")
}

// HandleOnDecoyStarted implements controller.Controller
func (f *f) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	panic("unimplemented")
}

// HandleOnDemoFinished implements controller.Controller
func (f *f) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	panic("unimplemented")
}

// HandleOnDemoUploadEnded implements controller.Controller
func (f *f) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	panic("unimplemented")
}

// HandleOnEvent implements controller.Controller
func (f *f) HandleOnEvent(ctx context.Context, p models.OnEventPayload) error {
	panic("unimplemented")
}

// HandleOnFlashbangDetonated implements controller.Controller
func (f *f) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnGameStateChanged implements controller.Controller
func (f *f) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	panic("unimplemented")
}

// HandleOnGoingLive implements controller.Controller
func (f *f) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	panic("unimplemented")
}

// HandleOnGrenadeThrown implements controller.Controller
func (f *f) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	panic("unimplemented")
}

// HandleOnHEGrenadeDetonated implements controller.Controller
func (f *f) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnKnifeRoundStarted implements controller.Controller
func (f *f) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	panic("unimplemented")
}

// HandleOnKnifeRoundWon implements controller.Controller
func (f *f) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	panic("unimplemented")
}

// HandleOnLoadMatchConfigFailed implements controller.Controller
func (f *f) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	panic("unimplemented")
}

// HandleOnMapPicked implements controller.Controller
func (f *f) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	panic("unimplemented")
}

// HandleOnMapResult implements controller.Controller
func (f *f) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	panic("unimplemented")
}

// HandleOnMapVetoed implements controller.Controller
func (f *f) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	panic("unimplemented")
}

// HandleOnMatchPaused implements controller.Controller
func (f *f) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	panic("unimplemented")
}

// HandleOnMatchUnpaused implements controller.Controller
func (f *f) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	panic("unimplemented")
}

// HandleOnMolotovDetonated implements controller.Controller
func (f *f) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerBecameMVP implements controller.Controller
func (f *f) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerConnected implements controller.Controller
func (f *f) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerDeath implements controller.Controller
func (f *f) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerDisconnected implements controller.Controller
func (f *f) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerSay implements controller.Controller
func (f *f) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
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

// HandleOnPreLoadMatchConfig implements controller.Controller
func (f *f) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	panic("unimplemented")
}

// HandleOnRoundEnd implements controller.Controller
func (f *f) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	panic("unimplemented")
}

// HandleOnRoundStart implements controller.Controller
func (f *f) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	panic("unimplemented")
}

// HandleOnRoundStatsUpdated implements controller.Controller
func (f *f) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	panic("unimplemented")
}

// HandleOnSeriesInit implements controller.Controller
func (f *f) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	panic("unimplemented")
}

// HandleOnSeriesResult implements controller.Controller
func (f *f) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	panic("unimplemented")
}

// HandleOnSidePicked implements controller.Controller
func (f *f) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	panic("unimplemented")
}

// HandleOnSmokeGrenadeDetonated implements controller.Controller
func (f *f) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnTeamReadyStatusChanged implements controller.Controller
func (f *f) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	panic("unimplemented")
}

// NewFirebaseController Get Firestore controller
func NewFirebaseController(ctx context.Context, c *firebase.App, setting ControllerSetting) (controller.Controller, error) {
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

	return &f{
		fs:      fs,
		s:       s,
		setting: setting,
	}, nil
}
