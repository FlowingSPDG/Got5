package fb

import (
	"context"
	"io"

	firebase "firebase.google.com/go"
	storage "firebase.google.com/go/storage"
	"github.com/FlowingSPDG/Got5/controller"
)

var _ controller.MatchLoader = (*firebaseMatchLoader)(nil)

type firebaseDemoUploader struct {
	s *storage.Client
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
	// Firebase Storage
	fbs, err := c.Storage(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseDemoUploader{
		s: fbs,
	}, nil
}
