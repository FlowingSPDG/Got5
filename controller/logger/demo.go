package logger

import (
	"context"
	"log"

	"github.com/FlowingSPDG/Got5/controller"
)

var _ controller.DemoUploader = (*loggerDemoUploader)(nil)

type loggerDemoUploader struct{}

// GetMatch implements controller.EventHandler
func (*loggerDemoUploader) Upload(ctx context.Context, mid string, filename string, b []byte) error {
	log.Println("mid:", mid)
	log.Println("filename:", filename)
	log.Println("b len:", len(b))
	return nil
}

// NewDemoUploader Get Logger pointer
func NewDemoUploader() controller.DemoUploader {
	return &loggerDemoUploader{}
}
