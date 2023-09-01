package godiscordwebhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Webhook interface {
	SetUsername(s string)                   // Sets the username of the webhook
	SetAvatarUrl(url string)                // Sets the icon of the webhook
	CreateAuthor(name, url, iconUrl string) // Creates and sets an author to the webhook
	SetAuthor(a Author)                     // Sets a given author to a given webhook
	SetContent(s string)                    // Message that appears outside of the embed. Up to 2k characters

	CreateField(name, value string, inline bool) // Creates and sets a field to the webhook
	AddField(f Field)                            // Adds a given field to the webhook
	AddFields(f []Field)                         // Adds a given list of fields to the webhook

	CreateFooter(text, iconUrl string) // Creates and sets a footer to the webhook
	SetFooter(f Footer)                // Sets a given footer to the webhook
	SetTimestamp(t time.Time)          // Sets a given time to the embed's footer
	SetCurrentTimestamp()              // Adds the current time to the embed's footer

	SetTitle(s string)       // Set the title of the embed
	SetUrl(s string)         // Sets the URL of the embed
	SetDescription(s string) // Sets the description of the embed
	SetColor(i int)          // Sets the color of the embed
	SetImage(url string)     // Sets the image of the embed
	SetThumbnail(url string) // Sets the thumbnail of the embed

	Send(url string) (*http.Response, error) // Sends the webhook to a given url
}

type webhook struct {
	Username  string  `json:"username"`   // The username of the webhook
	AvatarURL string  `json:"avatar_url"` // The avatar fo the webhook
	Content   string  `json:"content"`    // Text message. Up to 2000 characters.
	Embeds    []Embed `json:"embeds"`
}

func NewWebhook() Webhook {
	return newWebhook()
}

func newWebhook() *webhook {
	return &webhook{
		Embeds: []Embed{{Fields: make([]Field, 0)}},
	}
}

func (w *webhook) SetUsername(s string) {
	w.Username = s
}

func (w *webhook) SetContent(s string) {
	w.Content = s
}

func (w *webhook) SetAvatarUrl(s string) {
	w.AvatarURL = s
}

// Sends the webhook to a given url
func (w webhook) Send(url string) (*http.Response, error) {
	embed_str, _ := json.Marshal(w)
	payload := bytes.NewBuffer(embed_str)

	// Send POST request
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, url, payload)
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}
