package fb

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

var _ controller.MatchLoader = (*firebaseMatchLoader)(nil)

type firebaseMatchLoader struct {
	*Database // Databaseを埋め込む
	fs        *firestore.Client
}

// CheckAuth implements controller.MatchLoader
func (f *firebaseMatchLoader) CheckMatchAuth(ctx context.Context, mid string, reqAuth string) error {
	m, err := f.Database.GetMatch(ctx, mid)
	if err != nil {
		return err
	}
	if reqAuth != m.AuthValue {
		return fmt.Errorf("Auth mismatch")
	}
	return nil
}

// GetMatch implements controller.EventHandler
func (f *firebaseMatchLoader) Load(ctx context.Context, mid string) (models.G5Match, error) {
	return f.Database.GetMatch(ctx, mid)
}

// NewMatchLoader Get Firebase Match Loader
func NewMatchLoader(ctx context.Context, c *firebase.App) (controller.MatchLoader, error) {
	// Firestore
	fs, err := c.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	db, err := NewDatabase(ctx, c)
	if err != nil {
		return nil, err
	}

	return &firebaseMatchLoader{
		Database: db,
		fs:       fs,
	}, nil
}
