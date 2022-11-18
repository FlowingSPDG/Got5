package fb

import (
	"bytes"
	"context"
	"io"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	storage "firebase.google.com/go/storage"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

var _ controller.Database = (*firebaseDatabase)(nil)

// fs is Get5 API Database on firestore
type firebaseDatabase struct {
	fs *firestore.Client
	s  *storage.Client
}

// UpdateMatch implements controller.EventHandler
func (f *firebaseDatabase) UpdateMatch(ctx context.Context, mid string, m models.Match) error {
	if _, err := f.fs.Collection(CollectionMatch).Doc(mid).Set(ctx, m, firestore.MergeAll); err != nil {
		return err
	}
	return nil
}

// RegisterDemoFile implements controller.EventHandler
func (f *firebaseDatabase) RegisterDemoFile(ctx context.Context, mid string, filename string, b []byte) error {
	// MatchIDを取得しているので、Firestore上のMatchにdemoのURLを記載しても良い(TODO)
	bh, err := f.s.DefaultBucket()
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

// RegisterMatch implements controller.EventHandler
func (f *firebaseDatabase) RegisterMatch(ctx context.Context, m models.Match) (models.Match, error) {
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

// NewDatabase Get Firestore Database
func NewDatabase(ctx context.Context, c *firebase.App) (controller.Database, error) {
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

	return &firebaseDatabase{
		fs: fs,
		s:  s,
	}, nil
}
