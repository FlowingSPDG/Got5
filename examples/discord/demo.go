package discord

import (
	"context"
	"io"
)

// CheckDemoAuth implements controller.DemoUploader
func (*Discord) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	// TODO:auth
	return nil
}

// Upload implements controller.DemoUploader
func (d *Discord) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	io.Copy(io.Discard, r)
	return nil
}
