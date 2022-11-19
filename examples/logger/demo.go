package logger

import (
	"context"
	"io"
	"log"

	"github.com/FlowingSPDG/Got5/controller"
)

var _ controller.DemoUploader = (*loggerDemoUploader)(nil)

type loggerDemoUploader struct{}

// Verify implements controller.DemoUploader
func (*loggerDemoUploader) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	log.Println("mid", mid)
	log.Println("filename", filename)
	log.Println("mapNumber", mapNumber)
	log.Println("serverID", serverID)
	return nil
}

// GetMatch implements controller.EventHandler
func (*loggerDemoUploader) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	log.Println("mid:", mid)
	log.Println("filename:", filename)
	io.Copy(io.Discard, r)
	return nil
}

// NewDemoUploader Get Logger pointer
func NewDemoUploader() controller.DemoUploader {
	return &loggerDemoUploader{}
}
