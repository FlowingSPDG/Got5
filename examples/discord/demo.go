package discord

import (
	"context"
	"io"
)

// Upload implements controller.DemoUploader
func (d *Discord) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	io.Copy(io.Discard, r)
	return nil
}
