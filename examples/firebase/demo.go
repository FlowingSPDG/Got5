package fb

import (
	"context"
	"fmt"
	"io"

	firebase "firebase.google.com/go"
	storage "firebase.google.com/go/storage"

	"github.com/FlowingSPDG/Got5/controller"
)

var _ controller.MatchLoader = (*firebaseMatchLoader)(nil)

type firebaseDemoUploader struct {
	*Database
	s *storage.Client
}

func (f *firebaseDemoUploader) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	m, err := f.Database.GetMatch(ctx, mid)
	if err != nil {
		return err
	}
	if auth != m.AuthValue {
		return fmt.Errorf("Auth mismatch")
	}
	return nil
}

// GetMatch implements controller.EventHandler
func (f *firebaseDemoUploader) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	bh, err := f.s.DefaultBucket()
	if err != nil {
		return err
	}

	writer := bh.Object(filename).NewWriter(ctx)
	defer writer.Close()

	writer.ObjectAttrs.ContentType = "application/octet-stream"
	writer.ObjectAttrs.CacheControl = "public,max-age=86400"

	if _, err := io.Copy(writer, r); err != nil {
		return err
	}
	return nil
}

// NewDemoUploader Get Firebase Demo Uploader
func NewDemoUploader(ctx context.Context, c *firebase.App) (controller.DemoUploader, error) {
	db, err := NewDatabase(ctx, c)
	if err != nil {
		return nil, err
	}

	// Firebase Storage
	fbs, err := c.Storage(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseDemoUploader{
		s:        fbs,
		Database: db,
	}, nil
}
