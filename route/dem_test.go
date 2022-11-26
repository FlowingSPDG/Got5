package route_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/route"
)

var _ controller.DemoUploader = (*mockDemoUploaderGrant)(nil)
var _ controller.DemoUploader = (*mockDemoUploaderDenyAll)(nil)
var _ controller.DemoUploader = (*mockDemoUploaderDenyUpload)(nil)

// mockDemoUploaderGrant Grand all access
type mockDemoUploaderGrant struct{}

// Upload implements controller.DemoUploader
func (*mockDemoUploaderGrant) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	io.Copy(io.Discard, r)
	return nil
}

// mockDemoUploaderDeny Deny all access
type mockDemoUploaderDenyAll struct{}

// Upload implements controller.DemoUploader
func (*mockDemoUploaderDenyAll) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	io.Copy(io.Discard, r)
	return fmt.Errorf("Deny all access")
}

// mockDemoUploaderDenyUpload Deny all upload
type mockDemoUploaderDenyUpload struct{}

// Upload implements controller.DemoUploader
func (*mockDemoUploaderDenyUpload) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	io.Copy(io.Discard, r)
	return fmt.Errorf("Deny upload")
}

// Verify implements controller.DemoUploader
func (*mockDemoUploaderDenyUpload) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	return nil // always true
}

func TestDemoUploadTD(t *testing.T) {
	asserts := assert.New(t)

	cases := []struct {
		title      string
		uploader   controller.DemoUploader
		auth       controller.Auth
		statusCode int
		// err        error
		headers map[string]string
	}{
		{
			title:      "Grant All Access",
			uploader:   &mockDemoUploaderGrant{},
			statusCode: http.StatusOK,
			auth:       &mockAuth{},
			// err:        nil,
			headers: map[string]string{
				"Get5-DemoName":  "dem_name",
				"Get5-MatchId":   "mID_ABCDEFG",
				"Get5-MapNumber": "1",
				"Get5-ServerId":  "100",
			},
		},
		{
			title:      "Deny All",
			uploader:   &mockDemoUploaderDenyAll{},
			statusCode: http.StatusUnauthorized,
			auth:       &mockAuth{},
			headers: map[string]string{
				"Get5-DemoName":  "dem_name",
				"Get5-MatchId":   "mID_ABCDEFG",
				"Get5-MapNumber": "1",
				"Get5-ServerId":  "100",
			},
			// err:        nil,
		},
		{
			title:      "Internal Error on Upload",
			uploader:   &mockDemoUploaderDenyUpload{},
			statusCode: http.StatusInternalServerError,
			auth:       &mockAuth{},
			headers: map[string]string{
				"Get5-DemoName":  "dem_name",
				"Get5-MatchId":   "mID_ABCDEFG",
				"Get5-MapNumber": "1",
				"Get5-ServerId":  "100",
			},
			// err:        nil, //
		},
		{
			title:      "Invalid request",
			uploader:   &mockDemoUploaderGrant{},
			statusCode: http.StatusBadRequest,
			auth:       &mockAuth{},
			headers: map[string]string{
				"Get5-DemoName":  "dem_name",
				"Get5-MatchId":   "000",
				"Get5-MapNumber": "ABC",
				"Get5-ServerId":  "ABC",
			},
			// err:        nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			// Setup fiber
			app := fiber.New()
			g5test := app.Group("/get5testdemo") // /test
			route.SetupDemoUploadHandler(tt.uploader, tt.auth, g5test)

			r := httptest.NewRequest("POST", "/get5testdemo/demo", bytes.NewBuffer([]byte("demo file data")))
			for k, v := range tt.headers {
				r.Header.Add(k, v)
			}

			resp, _ := app.Test(r, -1)
			// asserts.Equal(tt.err, err)
			asserts.Equal(tt.statusCode, resp.StatusCode)
		})
	}
}
