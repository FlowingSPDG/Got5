package logger

import (
	"context"
	"log"

	"github.com/FlowingSPDG/Got5/models"
)

type Database struct{}

// UpdateMatch implements controller.EventHandler
func (*Database) UpdateMatch(ctx context.Context, mid string, m models.Match) error {
	return nil
}

// RegisterMatch implements controller.EventHandler
func (*Database) RegisterMatch(ctx context.Context, m models.Match) (models.G5Match, error) {
	log.Println(m)
	return models.Match{}, nil
}

// NewDatabase Get Logger pointer
func NewDatabase() *Database {
	return &Database{}
}
