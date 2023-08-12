package ginroute_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	got5 "github.com/FlowingSPDG/Got5"
	ginroute "github.com/FlowingSPDG/Got5/route/gin"
)

var _ got5.DemoUploader = (*mockDemoUploaderGrant)(nil)
var _ got5.DemoUploader = (*mockDemoUploaderDenyAll)(nil)
var _ got5.DemoUploader = (*mockDemoUploaderDenyUpload)(nil)

// mockDemoUploaderGrant Grand all access
type mockDemoUploaderGrant struct{}

// Upload implements got5.DemoUploader
func (*mockDemoUploaderGrant) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	io.Copy(io.Discard, r)
	return nil
}

// mockDemoUploaderDeny Deny all access
type mockDemoUploaderDenyAll struct{}

// Upload implements got5.DemoUploader
func (*mockDemoUploaderDenyAll) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	io.Copy(io.Discard, r)
	return fmt.Errorf("Deny all access")
}

// mockDemoUploaderDenyUpload Deny all upload
type mockDemoUploaderDenyUpload struct{}

// Upload implements got5.DemoUploader
func (*mockDemoUploaderDenyUpload) Upload(ctx context.Context, mid string, filename string, r io.Reader) error {
	io.Copy(io.Discard, r)
	return fmt.Errorf("Deny upload")
}

// Verify implements got5.DemoUploader
func (*mockDemoUploaderDenyUpload) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error {
	return nil // always true
}

func TestDemoUploadTD(t *testing.T) {
	asserts := assert.New(t)

	cases := []struct {
		title      string
		uploader   got5.DemoUploader
		auth       got5.Auth
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
				"Get5-FileName":  "dem_name",
				"Get5-MatchId":   "mID_ABCDEFG",
				"Get5-MapNumber": "1",
				"Get5-ServerId":  "100",
			},
		},
		{
			title:      "Deny All",
			uploader:   &mockDemoUploaderDenyAll{},
			statusCode: http.StatusInternalServerError,
			auth:       &mockAuth{},
			headers: map[string]string{
				"Get5-FileName":  "dem_name",
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
				"Get5-FileName":  "dem_name",
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
				"Get5-FileName":  "dem_name",
				"Get5-MatchId":   "000",
				"Get5-MapNumber": "ABC",
				"Get5-ServerId":  "ABC",
			},
			// err:        nil,
		},
	}

	for _, tt := range cases {
		t.Run("DEMO UPLOAD:"+tt.title, func(t *testing.T) {
			// Setup Gin
			app := gin.New()
			g5test := app.Group("/get5testdemo") // /test
			ginroute.SetupDemoUploadHandler(tt.uploader, tt.auth, g5test)

			r := httptest.NewRequest("POST", "/get5testdemo/demo", bytes.NewBuffer([]byte("demo file data")))
			for k, v := range tt.headers {
				r.Header.Add(k, v)
			}

			w := httptest.NewRecorder()
			app.ServeHTTP(w, r)

			asserts.Equal(tt.statusCode, w.Code)
		})
	}
}
