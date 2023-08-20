package godiscordwebhook

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestCreateFooter$ go-discord-webhook
func TestCreateFooter(t *testing.T) {
	tests := []struct {
		Name, IconUrl string
	}{
		{Name: "Test", IconUrl: "https://github.com/"},
		{Name: "", IconUrl: "https://github.com/"},
		{Name: "Test", IconUrl: ""},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			result := createFooter(test.Name, test.IconUrl)
			if assert.NotNil(t, result, "result should not be nil") {
				assert.Equal(t, test.Name, result.Text)
				assert.Equal(t, test.IconUrl, result.IconURL)
			}
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetFooter$ go-discord-webhook
func TestSetFooter(t *testing.T) {
	tests := []struct {
		Footer Footer
	}{
		{Footer: Footer{Text: "Test", IconURL: "Test"}},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w := newWebhook()
			setFooter(w, test.Footer)
			assert.Equal(t, test.Footer, w.Embeds[0].Footer)
			assert.Equal(t, test.Footer.Text, w.Embeds[0].Footer.Text)
			assert.Equal(t, test.Footer.IconURL, w.Embeds[0].Footer.IconURL)
		})
	}
}
