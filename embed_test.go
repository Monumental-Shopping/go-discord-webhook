package godiscordwebhook

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestNewUrlSource$ go-discord-webhook
func TestNewUrlSource(t *testing.T) {
	tests := []struct {
		S string
	}{
		{"Test"},
		{""},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			result := newUrlSource(test.S)
			if assert.NotNil(t, result, "result should not be nil") {
				assert.Equal(t, test.S, result.URL)
			}
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetTimestamp$ go-discord-webhook
func TestSetTimestamp(t *testing.T) {
	tests := []struct {
		Time time.Time
	}{
		{Time: time.Now()},
		{Time: time.Now().Add(-24 * time.Hour)},
		{Time: time.Unix(1405544146, 0)},
	}

	for testNum, test := range tests {
		w := newWebhook()
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			setTimestamp(w, test.Time)
			assert.Equal(t, test.Time.Format(time.RFC3339), w.Embeds[0].Timestamp)
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetTitle$ go-discord-webhook
func TestSetTitle(t *testing.T) {
	tests := []struct {
		Title string
	}{
		{Title: "Test"},
		{Title: ""},
	}

	for testNum, test := range tests {
		w := newWebhook()
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w.SetTitle(test.Title)
			assert.Equal(t, test.Title, w.Embeds[0].Title)
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetUrl$ go-discord-webhook
func TestSetUrl(t *testing.T) {
	tests := []struct {
		S string
	}{
		{S: "Test"},
		{S: ""},
	}

	for testNum, test := range tests {
		w := newWebhook()
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w.SetUrl(test.S)
			assert.Equal(t, test.S, w.Embeds[0].URL)
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetDescription$ go-discord-webhook
func TestSetDescription(t *testing.T) {
	tests := []struct {
		S string
	}{
		{S: "Test"},
		{S: ""},
	}

	for testNum, test := range tests {
		w := newWebhook()
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w.SetDescription(test.S)
			assert.Equal(t, test.S, w.Embeds[0].Description)
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetColor$ go-discord-webhook
func TestSetColor(t *testing.T) {
	tests := []struct {
		I int
	}{
		{I: 1},
		{I: 490492},
	}

	for testNum, test := range tests {
		w := newWebhook()
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w.SetColor(test.I)
			assert.Equal(t, test.I, w.Embeds[0].Color)
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetImage$ go-discord-webhook
func TestSetImage(t *testing.T) {
	tests := []struct {
		S UrlSource
	}{
		{S: UrlSource{"Something"}},
	}

	for testNum, test := range tests {
		w := newWebhook()
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w.SetImage(test.S.URL)
			assert.Equal(t, test.S, w.Embeds[0].Image)
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetThumbnail$ go-discord-webhook
func TestSetThumbnail(t *testing.T) {
	tests := []struct {
		S UrlSource
	}{
		{S: UrlSource{"Something"}},
	}

	for testNum, test := range tests {
		w := newWebhook()
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w.SetThumbnail(test.S.URL)
			assert.Equal(t, test.S, w.Embeds[0].Thumbnail)
		})
	}
}
