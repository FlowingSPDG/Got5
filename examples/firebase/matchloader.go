package fb

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

var _ controller.MatchLoader = (*firebaseMatchLoader)(nil)

type firebaseMatchLoader struct {
	fs *firestore.Client
}

// GetMatch implements controller.EventHandler
func (f *firebaseMatchLoader) Load(ctx context.Context, mid string) (models.G5Match, error) {
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

// NewMatchLoader Get Firebase Match Loader
func NewMatchLoader(ctx context.Context, c *firebase.App) (controller.MatchLoader, error) {
	// Firestore
	fs, err := c.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseMatchLoader{
		fs: fs,
	}, nil
}
