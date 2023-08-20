package godiscordwebhook

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestCreateAuthor$ go-discord-webhook
func TestCreateAuthor(t *testing.T) {
	tests := []struct {
		Name, Url, IconUrl string
	}{
		{Name: "Test", Url: "https://github.com/", IconUrl: "https://github.com/"},
		{Name: "Test", Url: "", IconUrl: "https://github.com/"},
		{Name: "Test", Url: "", IconUrl: ""},
		{Name: "", Url: "", IconUrl: ""},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			result := createAuthor(test.Name, test.Url, test.IconUrl)
			if assert.NotNil(t, result, "result should not be nil") {
				assert.Equal(t, test.Name, result.Name)
				assert.Equal(t, test.Url, result.URL)
				assert.Equal(t, test.IconUrl, result.IconURL)
			}
		})
	}
}

// To run this test individually, run the command:
// go test -v -timeout 30s -run ^TestSetAuthor$ go-discord-webhook
func TestSetAuthor(t *testing.T) {
	tests := []struct {
		A Author
	}{
		{Author{Name: "Test", URL: "https://github.com/", IconURL: "https://github.com/"}},
	}

	for testNum, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testNum+1), func(t *testing.T) {
			w := newWebhook()
			setAuthor(w, test.A)
			assert.Equal(t, w.Embeds[0].Author, test.A)
			assert.Equal(t, w.Embeds[0].Author.Name, test.A.Name)
			assert.Equal(t, w.Embeds[0].Author.IconURL, test.A.IconURL)
			assert.Equal(t, w.Embeds[0].Author.URL, test.A.URL)
		})
	}
}
