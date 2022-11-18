package logger

import (
	"context"
	"log"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

var _ controller.Database = (*loggerDatabase)(nil)

type loggerDatabase struct{}

// UpdateMatch implements controller.EventHandler
func (*loggerDatabase) UpdateMatch(ctx context.Context, mid string, m models.Match) error {
	return nil
}

// RegisterDemoFile implements controller.EventHandler
func (*loggerDatabase) RegisterDemoFile(ctx context.Context, mid string, filename string, b []byte) error {
	return nil
}

// RegisterMatch implements controller.EventHandler
func (*loggerDatabase) RegisterMatch(ctx context.Context, m models.Match) (models.Match, error) {
	log.Println(m)
	return models.Match{}, nil
}

// NewLoggerEventHandler Get Logger pointer
func NewDatabase() controller.Database {
	return &loggerDatabase{}
}
