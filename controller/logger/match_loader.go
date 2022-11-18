package logger

import (
	"context"
	"log"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

var _ controller.MatchLoader = (*loggerMatchLoader)(nil)

type loggerMatchLoader struct{}

// GetMatch implements controller.EventHandler
func (*loggerMatchLoader) Load(ctx context.Context, mid string) (models.G5Match, error) {
	log.Println(mid)
	return models.GetDefaultMatchBO1(), nil
}

// NewMatchLoader Get Logger pointer
func NewMatchLoader() controller.MatchLoader {
	return &loggerMatchLoader{}
}
