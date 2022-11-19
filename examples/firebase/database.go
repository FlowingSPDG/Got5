package fb

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

// Database is Get5 API Database on firestore
type Database struct {
	fs *firestore.Client
}

// UpdateMatch implements controller.EventHandler
func (f *Database) UpdateMatch(ctx context.Context, mid string, m Match) error {
	if _, err := f.fs.Collection(CollectionMatch).Doc(mid).Set(ctx, m, firestore.MergeAll); err != nil {
		return err
	}
	return nil
}

// RegisterMatch implements controller.EventHandler
func (f *Database) RegisterMatch(ctx context.Context, m Match) (Match, error) {
	ret := Match{}
	ref := f.fs.Collection(CollectionMatch).NewDoc()
	ret.MatchID = ref.ID // MatchIDを上書きする
	_, err := ref.Set(ctx, ret)
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
func NewDatabase(ctx context.Context, c *firebase.App) (*Database, error) {
	// Firestore
	fs, err := c.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &Database{
		fs: fs,
	}, nil
}
